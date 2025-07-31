package util

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/feerdim/boilerplate-golang/log"
	"gorm.io/gorm"
)

func Transaction(
	ctx context.Context,
	db *gorm.DB,
	txFunc func(db *gorm.DB) (err error),
) (err error) {
	db = db.WithContext(ctx)

	tx := db.Begin(&sql.TxOptions{})

	err = txFunc(db)
	if err != nil {
		if errRollback := tx.Rollback().Error; errRollback != nil {
			log.WithContext(ctx).Error(errRollback, "error rollback")
			return
		}

		return
	}

	if err = tx.Commit().Error; err != nil {
		log.WithContext(ctx).Error(err, "error commit")
		return
	}

	return
}

func GenerateStatement(stmt *string, op, key, col, dtype string, data []string) string {
	var rword string

	if len(data) == 0 {
		rword = "TRUE"
	} else {
		switch dtype {
		case "char":
			rword = fmt.Sprintf("%s %s ('%v')", col, op, strings.Join(data, "', '"))
		case "int":
			rword = fmt.Sprintf("%s %s (%v)", col, op, strings.Join(data, ", "))
		}
	}

	return strings.ReplaceAll(*stmt, key, rword)
}

func ValidateUnique(err, errNew error) error {
	if strings.Contains(err.Error(), "SQLSTATE 23505") {
		err = errNew
	}

	return err
}
