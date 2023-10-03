package data

import "fmt"

type Track struct {
	Id       int64
	Title    string
	Duration int64
	Song     int64
	DeezerId int64
	IsLiked  bool
	Img      int64
}

func GetTrackByAlbumId(id int) []Track {
	db := GetDb()
	defer db.Close()

	rows, _ := db.Query("select tracks.id, tracks.title, tracks.duration, tracks.song, tracks.deezer_id, a.img "+
		"from tracks "+
		"inner join tracks_albums ta on tracks.id = ta.track_id "+
		"inner join albums a on a.id = ta.album_id "+
		"where ta.album_id = ?", id)

	defer rows.Close()

	result := make([]Track, 0, 20)
	for rows.Next() {
		var track Track
		if err := rows.Scan(&track.Id, &track.Title, &track.Duration, &track.Song, &track.DeezerId, &track.Img); err != nil {
			fmt.Printf("%v\n", err)
		}
		track.IsLiked = IsTrackLiked(track.Id)
		result = append(result, track)
	}

	return result
}

func GetTrackId(id int) Track {
	db := GetDb()
	row := db.QueryRow("select tracks.id, tracks.title, tracks.duration, tracks.song, tracks.deezer_id from tracks where tracks.id = ?", id)

	var track Track
	row.Scan(&track.Id, &track.Title, &track.Duration, &track.Song, &track.DeezerId)
	track.IsLiked = IsTrackLiked(track.Id)

	return track
}
