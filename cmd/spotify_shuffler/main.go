package spotify_shuffler

import "github.com/spotify-shuffler/pkg/spotify"

func main() {
	manager := spotify.Manager{}
	_ = manager.ShufflePlaylist()
}

