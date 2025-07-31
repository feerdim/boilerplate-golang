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
	"github.com/labstack/echo/v4"
)

func main() {
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
		log.Printf("ERROR mail dialer : %s", err.Error())
		return
	}

	t := toolkit.NewToolkit(db, dbx, mail, stg)

	r := config.NewRuntime()

	e := echo.New()
	e.HideBanner = true
	e.Validator = config.NewValidator()
	e.HTTPErrorHandler = domain.ErrorHandler()

	go shutdown(ctx, r, e, dbx)

	domain.Routes(e, t)

	if err := e.Start(fmt.Sprintf("%s:%d", r.Host, r.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("ERROR starting http server : %s", err.Error())
	}
}

func shutdown(ctx context.Context, r *config.Runtime, e *echo.Echo, dbx *sqlx.DB) {
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), r.ShutdownTimeoutDuration)
	defer cancel()

	<-time.After(r.ShutdownWaitDuration)

	if err := e.Shutdown(ctx); err != nil {
		log.Printf("ERROR shutdown server : %s", err.Error())
		return
	}

	if err := dbx.Close(); err != nil {
		log.Printf("ERROR close database connection : %s", err.Error())
		return
	}
}
