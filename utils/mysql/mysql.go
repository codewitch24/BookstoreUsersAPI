package mysql

import (
	"fmt"
	"github.com/VividCortex/mysqlerr"
	"github.com/codewitch24/BookstoreUsersAPI/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	driverErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(
				fmt.Sprintf("No record matching given id"),
			)
		}
		return errors.NewInternalServerError("Error parsing database response")
	}
	switch driverErr.Number {
	case mysqlerr.ER_DUP_ENTRY:
		return errors.NewBadRequestError("Invalid data")
	}
	return errors.NewInternalServerError("Error processing request")
}
