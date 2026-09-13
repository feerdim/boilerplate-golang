package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/feerdim/boilerplate-golang/config"
	logger "github.com/feerdim/boilerplate-golang/log"
	"github.com/feerdim/boilerplate-golang/src/domain"
	"github.com/feerdim/boilerplate-golang/src/toolkit"
	"github.com/feerdim/boilerplate-golang/src/toolkit/mail"
	"github.com/feerdim/boilerplate-golang/src/toolkit/storage"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
)

func main() {
	const readHeaderTimeout = 10 * time.Second

	var err error

	if os.Getenv("APP_ENV") == "" {
		err = godotenv.Load(".env")
		if err != nil {
			log.Fatalf("ERROR load env file : %s", err.Error())
		}
	}

	ctx, cancel := config.NewRuntimeContext()

	defer func() {
		cancel()

		if err != nil {
			log.Printf("found error : %s", err.Error())
		}
	}()

	config.SetDefaultTimezone()

	logger.SetDefaultLogger()

	dbx, db, err := config.NewDatabase()
	if err != nil {
		log.Printf("ERROR database : %s", err.Error())
		return
	}

	mail, err := mail.NewMail()
	if err != nil {
		log.Printf("ERROR mail dialer : %s", err.Error())
		return
	}

	stg, err := storage.NewStorage(ctx)
	if err != nil {
		log.Printf("ERROR storage : %s", err.Error())
		return
	}

	t := toolkit.NewToolkit(db, dbx, mail, stg)

	r := config.NewRuntime()

	e := echo.New()
	e.Validator = config.NewValidator()
	e.HTTPErrorHandler = domain.ErrorHandler()

	s := &http.Server{
		Addr:              fmt.Sprintf("%s:%d", r.Host, r.Port),
		Handler:           e,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	domain.Routes(e, t)

	go func() {
		log.Printf("Starting http server %s:%d\n--------------------------------", r.Host, r.Port)

		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("ERROR starting http server : %s", err.Error())
		}
	}()

	gracefulShutdown(ctx, r, s, dbx)
}

func gracefulShutdown(ctx context.Context, r *config.Runtime, s *http.Server, dbx *sqlx.DB) {
	<-ctx.Done()

	log.Printf("Graceful shutdown starting ...")

	ctx, cancel := context.WithTimeout(context.Background(), r.ShutdownTimeoutDuration)
	defer cancel()

	<-time.After(r.ShutdownWaitDuration)

	if err := s.Shutdown(ctx); err != nil {
		log.Printf("ERROR shutdown server : %s", err.Error())
		return
	}

	if err := dbx.Close(); err != nil {
		log.Printf("ERROR close database connection : %s", err.Error())
		return
	}

	log.Printf("Graceful shutdown success")
}
