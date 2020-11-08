package goose

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/pkg/errors"
)

type tmplVars struct {
	Version   string
	CamelName string
}

// Create writes a new blank migration file.
func CreateWithTemplate(db *sql.DB, dir string, tmpl *template.Template, name, migrationType string) error {
	return def.CreateWithTemplate(db, dir, tmpl, name, migrationType)
}

// Create writes a new blank migration file.
func (in *Instance) CreateWithTemplate(db *sql.DB, dir string, tmpl *template.Template, name, migrationType string) error {
	version := time.Now().Format(timestampFormat)
	filename := fmt.Sprintf("%v_%v.%v", version, snakeCase(name), migrationType)

	if tmpl == nil {
		if migrationType == "go" {
			tmpl = goSQLMigrationTemplate
		} else {
			tmpl = sqlMigrationTemplate
		}
	}

	path := filepath.Join(dir, filename)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return errors.Wrap(err, "failed to create migration file")
	}

	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "failed to create migration file")
	}
	defer f.Close()

	vars := tmplVars{
		Version:   version,
		CamelName: camelCase(name),
	}
	if err := tmpl.Execute(f, vars); err != nil {
		return errors.Wrap(err, "failed to execute tmpl")
	}

	in.log.Printf("Created new file: %s\n", f.Name())
	return nil
}

// Create writes a new blank migration file.
func Create(db *sql.DB, dir, name, migrationType string) error {
	return def.Create(db, dir, name, migrationType)
}

// Create writes a new blank migration file.
func (in *Instance) Create(db *sql.DB, dir, name, migrationType string) error {
	return in.CreateWithTemplate(db, dir, nil, name, migrationType)
}

var sqlMigrationTemplate = template.Must(template.New("goose.sql-migration").Parse(`-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
`))

var goSQLMigrationTemplate = template.Must(template.New("goose.go-migration").Parse(`package migrations

import (
	"database/sql"
	"github.com/powerman/goose/v2"
)

func init() {
	goose.AddMigration(up{{.CamelName}}, down{{.CamelName}})
}

func up{{.CamelName}}(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func down{{.CamelName}}(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
`))
