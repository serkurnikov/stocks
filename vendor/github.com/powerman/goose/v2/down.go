package goose

import (
	"database/sql"
	"fmt"
)

// Down rolls back a single migration from the current version.
func Down(db *sql.DB, dir string) error { return def.Down(db, dir) }

// Down rolls back a single migration from the current version.
func (in *Instance) Down(db *sql.DB, dir string) error {
	currentVersion, err := in.GetDBVersion(db)
	if err != nil {
		return err
	}

	migrations, err := in.CollectMigrations(dir, minVersion, maxVersion)
	if err != nil {
		return err
	}

	current, err := migrations.Current(currentVersion)
	if err != nil {
		return fmt.Errorf("no migration %v", currentVersion)
	}

	return current.Down(db)
}

// DownTo rolls back migrations to a specific version.
func DownTo(db *sql.DB, dir string, version int64) error { return def.DownTo(db, dir, version) }

// DownTo rolls back migrations to a specific version.
func (in *Instance) DownTo(db *sql.DB, dir string, version int64) error {
	migrations, err := in.CollectMigrations(dir, minVersion, maxVersion)
	if err != nil {
		return err
	}

	for {
		currentVersion, err := in.GetDBVersion(db)
		if err != nil {
			return err
		}

		current, err := migrations.Current(currentVersion)
		if err != nil {
			in.log.Printf("goose: no migrations to run. current version: %d\n", currentVersion)
			return nil
		}

		if current.Version <= version {
			in.log.Printf("goose: no migrations to run. current version: %d\n", currentVersion)
			return nil
		}

		if err = current.Down(db); err != nil {
			return err
		}
	}
}
