package bolt

import (
	"github.com/asdine/storm/v3"

	"github.com/shmelll/filebrowser-master/auth"
	"github.com/shmelll/filebrowser-master/settings"
	"github.com/shmelll/filebrowser-master/share"
	"github.com/shmelll/filebrowser-master/storage"
	"github.com/shmelll/filebrowser-master/users"
)

// NewStorage creates a storage.Storage based on Bolt DB.
func NewStorage(db *storm.DB) (*storage.Storage, error) {
	userStore := users.NewStorage(usersBackend{db: db})
	shareStore := share.NewStorage(shareBackend{db: db})
	settingsStore := settings.NewStorage(settingsBackend{db: db})
	authStore := auth.NewStorage(authBackend{db: db}, userStore)

	err := save(db, "version", 2) //nolint:gomnd
	if err != nil {
		return nil, err
	}

	return &storage.Storage{
		Auth:     authStore,
		Users:    userStore,
		Share:    shareStore,
		Settings: settingsStore,
	}, nil
}
