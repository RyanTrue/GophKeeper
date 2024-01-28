package memory

import "github.com/RyanTrue/GophKeeper/internal/repository"

type Factory struct{}

var _ repository.Factory = (*Factory)(nil)

func NewFactory() repository.Factory {
	return &Factory{}
}

func (f *Factory) CreateUsersRepository() repository.Users {
	return NewUsersRepository()
}

func (f *Factory) CreateSettingsRepository() repository.Settings {
	return NewSettingsRepository()
}

func (f *Factory) CreateCredsSecretsRepository() repository.CredsSecrets {
	return NewCredsSecretsRepository()
}
