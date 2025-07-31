package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() (dbx *sqlx.DB, db *gorm.DB, err error) {
	opt, err := newDatabaseOption()
	if err != nil {
		return
	}

	switch opt.driver {
	case "postgresql":
		dbx, db, err = opt.newPostgresql()
	case "mysql":
		dbx, db, err = opt.newMysql()
	case "":
	default:
		err = errors.Wrapf(errors.New("invalid datasources driver"), "db: driver=%s", opt.driver)
		return
	}

	go opt.keepAlive(dbx)

	return
}

type databaseOption struct {
	driver       string
	host         string
	port         int
	username     string
	password     string
	schema       string
	sslmode      string
	usePrivate   string
	instanceName string
	*databaseConnection
}

type databaseConnection struct {
	maxIdle           int
	maxOpen           int
	maxLifetime       time.Duration
	keepAliveInterval time.Duration
	gormOption
}

type gormOption struct {
	isLog         bool
	slowThreshold time.Duration
	level         logger.LogLevel
	ignoreErr     bool
	colorful      bool
	cfgLogger     logger.Interface
}

func newDatabaseOption() (*databaseOption, error) {
	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	schema := os.Getenv("DB_SCHEMA")
	sslmode := os.Getenv("DB_SSLMODE")
	usePrivate := os.Getenv("PRIVATE_IP")
	instanceName := os.Getenv("INSTANCE_CONNECTION_NAME")

	if portStr == "" {
		portStr = "0"
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, errors.Wrapf(err, "error parse int on port db env : %s", portStr)
	}

	if host == "" {
		return nil, errors.Wrapf(errors.New("invalid data source host or port"), "db: host=%s port=%d", host, port)
	}

	conn := &databaseConnection{
		maxIdle:           defaultDBMaxIdle,
		maxOpen:           defaultDBMaxOpen,
		maxLifetime:       defaultDBMaxLifetime,
		keepAliveInterval: defaultDBKeepAliveInterval,
		gormOption: gormOption{
			isLog:         defaultGormLog,
			slowThreshold: defaultGormLogSlowThreshold,
			level:         defaultGormLogLevel,
			ignoreErr:     defaultGormLogIgnoreErr,
			colorful:      defaultGormLogColorful,
		},
	}

	conn.maxIdle = parseInt(conn.maxIdle, os.Getenv("DB_MAX_IDLE_CONN"))
	conn.maxOpen = parseInt(conn.maxOpen, os.Getenv("DB_MAX_OPEN_CONN"))
	conn.maxLifetime = parseDuration(conn.maxLifetime, os.Getenv("DB_MAX_LIFETIME_CONN"))
	conn.keepAliveInterval = parseDuration(conn.keepAliveInterval, os.Getenv("DB_KEEP_ALIVE_INTERVAL_CONN"))

	if conn.isLog {
		conn.cfgLogger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             conn.slowThreshold,
				LogLevel:                  conn.level,
				IgnoreRecordNotFoundError: conn.ignoreErr,
				Colorful:                  conn.colorful,
			},
		)
	}

	return &databaseOption{
		driver:             driver,
		host:               host,
		port:               port,
		username:           username,
		password:           password,
		schema:             schema,
		sslmode:            sslmode,
		usePrivate:         usePrivate,
		instanceName:       instanceName,
		databaseConnection: conn,
	}, nil
}

func (opt databaseOption) openSQL(driver, source string) (dbx *sqlx.DB, err error) {
	dbx, err = sqlx.Open(driver, source)
	if err != nil {
		return
	}

	dbx.SetMaxIdleConns(opt.maxIdle)
	dbx.SetMaxOpenConns(opt.maxOpen)
	dbx.SetConnMaxLifetime(opt.maxLifetime)
	dbx.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)

	return
}

func (opt databaseOption) keepAlive(db *sqlx.DB) {
	for {
		err := db.Ping()
		if err != nil {
			log.Printf("ERROR db.keepAlive driver=%s schema=%s \n%s \n\n", opt.driver, opt.schema, err)
		}

		time.Sleep(opt.keepAliveInterval)
	}
}
