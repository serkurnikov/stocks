package goose

import (
	"database/sql"
	"regexp"

	"github.com/pkg/errors"
)

// Run a migration specified in raw SQL.
//
// Sections of the script can be annotated with a special comment,
// starting with "-- +goose" to specify whether the section should
// be applied during an Up or Down migration
//
// All statements following an Up or Down directive are grouped together
// until another direction directive is found.
func (in *Instance) runSQLMigration(db *sql.DB, statements []string, useTx bool, v int64, direction bool) error {
	if useTx {
		// TRANSACTION.

		in.verboseInfo("Begin transaction")

		tx, err := db.Begin()
		if err != nil {
			return errors.Wrap(err, "failed to begin transaction")
		}

		for _, query := range statements {
			in.verboseInfo("Executing statement: %s\n", clearStatement(query))
			if _, err = tx.Exec(query); err != nil {
				in.verboseInfo("Rollback transaction")
				tx.Rollback()
				return errors.Wrapf(err, "failed to execute SQL query %q", clearStatement(query))
			}
		}

		if direction {
			if _, err := tx.Exec(in.GetDialect().insertVersionSQL(), v, direction); err != nil {
				in.verboseInfo("Rollback transaction")
				tx.Rollback()
				return errors.Wrap(err, "failed to insert new goose version")
			}
		} else {
			if _, err := tx.Exec(in.GetDialect().deleteVersionSQL(), v); err != nil {
				in.verboseInfo("Rollback transaction")
				tx.Rollback()
				return errors.Wrap(err, "failed to delete goose version")
			}
		}

		in.verboseInfo("Commit transaction")
		if err := tx.Commit(); err != nil {
			return errors.Wrap(err, "failed to commit transaction")
		}

		return nil
	}

	// NO TRANSACTION.
	for _, query := range statements {
		in.verboseInfo("Executing statement: %s", clearStatement(query))
		if _, err := db.Exec(query); err != nil {
			return errors.Wrapf(err, "failed to execute SQL query %q", clearStatement(query))
		}
	}
	if _, err := db.Exec(in.GetDialect().insertVersionSQL(), v, direction); err != nil {
		return errors.Wrap(err, "failed to insert new goose version")
	}

	return nil
}

const (
	grayColor  = "\033[90m"
	resetColor = "\033[00m"
)

func (in *Instance) verboseInfo(s string, args ...interface{}) {
	if in.verbose {
		in.log.Printf(grayColor+s+resetColor, args...)
	}
}

var (
	matchSQLComments = regexp.MustCompile(`(?m)^--.*$[\r\n]*`)
	matchEmptyEOL    = regexp.MustCompile(`(?m)^$[\r\n]*`) // TODO: Duplicate
)

func clearStatement(s string) string {
	s = matchSQLComments.ReplaceAllString(s, ``)
	return matchEmptyEOL.ReplaceAllString(s, ``)
}
