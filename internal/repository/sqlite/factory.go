package sqlite

import "github.com/RyanTrue/GophKeeper/internal/repository"

type Factory struct {
	db *SQLite
}

var _ repository.Factory = (*Factory)(nil)

func NewFactory(db *SQLite) repository.Factory {
	return &Factory{
		db: db,
	}
}

func (f *Factory) CreateUsersRepository() repository.Users {
	return NewUsersRepository(f.db)
}

func (f *Factory) CreateSettingsRepository() repository.Settings {
	return NewSettingsRepository(f.db)
}

func (f *Factory) CreateCredsSecretsRepository() repository.CredsSecrets {
	return NewCredsSecretsRepository(f.db)
}
