package storage

import (
	"github.com/shmelll/filebrowser-master-master/auth"
	"github.com/shmelll/filebrowser-master-master/settings"
	"github.com/shmelll/filebrowser-master-master/share"
	"github.com/shmelll/filebrowser-master-master/users"
)

// Storage is a storage powered by a Backend which makes the necessary
// verifications when fetching and saving data to ensure consistency.
type Storage struct {
	Users    users.Store
	Share    *share.Storage
	Auth     *auth.Storage
	Settings *settings.Storage
}
