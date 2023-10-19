package data

import "fmt"

type PlaylistElement struct {
	Track   Track
	Order   int64
	Current bool
	Artist  Artist
	AlbumId int64
}

func AddAblumToPlaylist(albumId int64, trackId int64) {
	db := GetDb()
	defer db.Close()
	tracks := GetTrackByAlbumId(int(albumId))
	_, err := db.Exec("DELETE from playlists")
	if err != nil {
		fmt.Println("Could not DELETE playlists")
		fmt.Printf("%v\n", err)
	} else {
		for i, track := range tracks {
			_, err := db.Exec("INSERT INTO playlists (track_id, 'order', current) VALUES (?, ?, ?)", track.Id, i, trackId == track.Id)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		}
	}
}

func AddTrackToPlaylist(trackId int64) {
	db := GetDb()
	defer db.Close()
	orderQuery := db.QueryRow("SELECT max(order) from playlists")
	var order int64
	orderQuery.Scan(&order)
	_, err := db.Exec("INSERT INTO playlists (song_id, order, current) VALUES (?, ?, ?)", trackId, order+1, false)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

func GetPlaylist() []PlaylistElement {
	result := make([]PlaylistElement, 0, 20)
	db := GetDb()
	defer db.Close()
	rows, err := db.Query(
		"select t.id, t.title, t.duration, t.song, t.deezer_id, a.img, " +
			"a.id, " +
			"p.'order', p.current " +
			"from playlists p " +
			"inner join tracks t on t.id = p.track_id " +
			"inner join tracks_albums ta on ta.track_id = t.id " +
			"inner join albums a on a.id = ta.album_id")
	if err != nil {
		fmt.Println("Could not query playlists")
	} else {
		for rows.Next() {
			var playlistElement PlaylistElement
			if err := rows.Scan(
				&playlistElement.Track.Id, &playlistElement.Track.Title, &playlistElement.Track.Duration, &playlistElement.Track.Song, &playlistElement.Track.DeezerId, &playlistElement.Track.Img,
				&playlistElement.AlbumId, &playlistElement.Order, &playlistElement.Current,
			); err != nil {
				fmt.Println("Could not scan row")
				fmt.Printf("%v\n", err)
			} else {
				playlistElement.Track.IsLiked = IsTrackLiked(playlistElement.Track.Id)
				playlistElement.Artist = GetArtistByTrackId(int(playlistElement.Track.Id))
				result = append(result, playlistElement)
			}
		}
		rows.Close()
	}
	return result
}

func SetCurrentPlaylistTrack(trackId int64) {
	db := GetDb()
	defer db.Close()
	_, err := db.Exec("update playlists set current = 0")
	if err != nil {
		fmt.Println("Could not set current track")
		fmt.Printf("%v\n", err)
	} else {
		_, err := db.Exec("update playlists set current = 1 where track_id = ?", trackId)
		if err != nil {
			fmt.Println("Could not set current track")
			fmt.Printf("%v\n", err)
		}
	}
}
func GetCurrentPlaylistTrack() int64 {
	db := GetDb()
	defer db.Close()
	row := db.QueryRow("select track_id from playlists where current = 1")
	var result int64
	row.Scan(&result)
	return result
}

func SetNextTrackPlaylist() bool {
	db := GetDb()
	defer db.Close()

	row := db.QueryRow("select \"order\" from playlists where current = 1")
	var currentSong int64
	row.Scan(&currentSong)

	row = db.QueryRow("select max(\"order\") from playlists")
	var maxOrder int64
	row.Scan(&maxOrder)

	_, err := db.Exec("update playlists set current = 0 where current = 1")
	if err != nil {
		fmt.Println("Could not update playlist current")
		fmt.Printf("%v\n", err)
		return false
	}

	_, err = db.Exec("update playlists set current = 1 where \"order\" = ?", (currentSong+1)%(maxOrder+1))
	if err != nil {
		fmt.Println("Could not update playlist current playing song")
		fmt.Printf("%v\n", err)
		return false
	}
	return true
}
