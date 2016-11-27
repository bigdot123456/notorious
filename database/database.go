package db

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	// We use a blank import here because I'm afraid of breaking anything
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// SQLStore is the base implementation for a database which will be used to
// store stats and retrieve whitelisted torrents.
type SQLStore interface {
	OpenConnection() (*gorm.DB, error)
	GetTorrent(string) (*Torrent, error)
	GetWhitelistedTorrent(string) (*White_Torrent, error)
	UpdateStats(uint64, uint64)
	UpdateTorrentStats(int64, int64)
	ScrapeTorrent(string) *Torrent
	GetWhitelistedTorrents() (*sql.Rows, error)
	UpdatePeerStats(uint64, uint64, string)
}
