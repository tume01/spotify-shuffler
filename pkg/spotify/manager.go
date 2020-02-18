package spotify

import "github.com/spotify-shuffler/pkg/entities"

type Manager struct {
	config     ManagerConfig
	repository MusicRepository
}

func (m *Manager) RefreshPlaylist(id string) error {
	return m.repository.BlankPlaylist(id)
}

func (m *Manager) FindSongs(max int) ([]entities.Song, error) {
	return []entities.Song{}, nil
}

func (m *Manager) ShufflePlaylist() error {
	playlist := m.repository.GetPlaylist()
	err := m.RefreshPlaylist(playlist.ID)
	if err != nil {
		return err
	}

	songs, err := m.FindSongs(m.config.MaxSongs)

	if err != nil {
		return err
	}
	err = m.repository.AddSongsToPlaylist(playlist.ID, songs)
	return nil
}