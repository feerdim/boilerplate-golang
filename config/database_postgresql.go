package config

import (
	"context"
	"fmt"
	"log"
	"net"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (opt databaseOption) newPostgresql() (dbx *sqlx.DB, db *gorm.DB, err error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", opt.host, opt.port, opt.username, opt.password, opt.schema, opt.sslmode)

	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		err = errors.Wrap(err, "postgres: failed parse connection config")
		return
	}

	if opt.instanceName != "" {
		var (
			opts []cloudsqlconn.Option
			d    *cloudsqlconn.Dialer
		)

		if opt.usePrivate != "" {
			opts = append(opts, cloudsqlconn.WithDefaultDialOptions(cloudsqlconn.WithPrivateIP()))
		}

		d, err = cloudsqlconn.NewDialer(context.Background(), opts...)
		if err != nil {
			err = errors.Wrap(err, "cloudsql: failed to make dialer")
			return
		}

		config.DialFunc = func(ctx context.Context, _, _ string) (net.Conn, error) {
			return d.Dial(ctx, opt.instanceName)
		}
	}

	dbx, err = opt.openSQL("pgx", dsn)
	if err != nil {
		err = errors.Wrap(err, "postgres: failed to open connection")
		return
	}

	db, err = gorm.Open(postgres.New(postgres.Config{Conn: dbx.DB}), &gorm.Config{Logger: opt.cfgLogger})
	if err != nil {
		err = errors.Wrap(err, "gorm: failed to open connection")
		return
	}

	log.Printf("successfully connected to postgresql %s:%d", opt.host, opt.port)

	return
}
