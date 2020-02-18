package spotify

import "github.com/spotify-shuffler/pkg/entities"

type MusicRepository interface {
	GetPlaylist() entities.Playlist
	GetSongs() []entities.Song
	GetCurrentSongs(id string) []entities.Song
	RemoveSongsFromPlaylist(id string, songs []entities.Song) error
	AddSongsToPlaylist(id string, songs []entities.Song) error
	BlankPlaylist(id string) error
}

type Repository struct {
	
}

func (r Repository) BlankPlaylist(id string) error {
	panic("implement me")
}

func (r Repository) GetPlaylist() entities.Playlist {
	return entities.Playlist{ID: "503S6ES90tdov5agN0DXtj"}
}

func (r Repository) GetSongs() []entities.Song {
	panic("implement me")
}

func (r Repository) GetCurrentSongs(id string) []entities.Song {
	panic("implement me")
}

func (r Repository) RemoveSongsFromPlaylist(id string, songs []entities.Song) error {
	panic("implement me")
}

func (r Repository) AddSongsToPlaylist(id string, songs []entities.Song) error {
	panic("implement me")
}


