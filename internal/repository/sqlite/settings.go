package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"github.com/RyanTrue/GophKeeper/internal/repository"
)

type SettingsRepository struct {
	db *SQLite
}

type settings struct {
	key   string
	value string
}

var _ repository.Settings = (*SettingsRepository)(nil)

func NewSettingsRepository(db *SQLite) repository.Settings {
	return &SettingsRepository{
		db: db,
	}
}

func (r *SettingsRepository) Get(ctx context.Context, key string) (string, bool, error) {
	query := `SELECT key, value FROM settings WHERE key = $1;`

	setting := new(settings)

	if err := r.db.QueryRowContext(ctx, query, key).Scan(&setting.key, &setting.value); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", false, nil
		}
		return "", false, err
	}

	return setting.value, true, nil
}

func (r *SettingsRepository) Set(ctx context.Context, key, value string) (bool, error) {
	query := `SELECT key, value FROM settings WHERE key = $1;`

	setting := new(settings)
	existed := false

	err := r.db.QueryRowContext(ctx, query, key).Scan(&setting.key, &setting.value)
	if errors.Is(err, sql.ErrNoRows) {
		existed = false
		query = `INSERT INTO settings (key, value) VALUES ($1, $2);`
	} else {
		existed = true
		query = `UPDATE settings SET value = $2 WHERE key = $1;`
	}

	if _, err = r.db.ExecContext(ctx, query, key, value); err != nil {
		return existed, err
	}

	return existed, nil
}

func (r *SettingsRepository) Delete(ctx context.Context, key string) (bool, error) {
	query := `DELETE FROM settings WHERE key = $1;`

	result, err := r.db.ExecContext(ctx, query, key)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (r *SettingsRepository) Truncate(ctx context.Context) error {
	query := `DELETE FROM settings;`

	if _, err := r.db.ExecContext(ctx, query); err != nil {
		return err
	}

	return nil
}
