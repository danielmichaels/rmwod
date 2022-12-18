// Package database
//
// seeder fetches the initial playlist and loads it into the database.
package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/danielmichaels/rmwod/assets"
	"github.com/danielmichaels/rmwod/internal/config"
	"log"
)

var playlist = "migrations/playlist.json"

func Seed() {
	p, err := assets.EmbeddedFiles.ReadFile("migrations/playlist.json")
	if err != nil {
		log.Fatalf("failed to find %q", playlist)
	}

	var videos []Videos
	err = json.Unmarshal(p, &videos)

	dbconn, err := sql.Open("sqlite3", config.AppConfig().Db.DbName)
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	db := New(dbconn)

	for _, v := range videos {
		var ip InsertVideoParams
		ip.Duration = v.Duration
		ip.Description = v.Description
		ip.Title = v.Title
		ip.Url = v.Url
		err := db.InsertVideo(context.Background(), ip)
		if err != nil {
			log.Fatalf("failed to insert: %q with error: %s", ip.Title, err)
		}

	}
}
