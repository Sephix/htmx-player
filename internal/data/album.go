package data

import "fmt"

type Album struct {
	Id          int64
	Name        string
	ReleaseDate string
	Img         int64
	DeezerID    int64
}

func GetAlbumByArtistId(id int, name string) []Album {
	db := GetDb()
	defer db.Close()
	rows, _ := db.Query("select albums.id, albums.name, albums.release_date, albums.img from albums inner join artists_albums aa on albums.id = aa.album_id where artist_id = ? and name like ?", id, "%"+name+"%")

	defer rows.Close()

	result := make([]Album, 0)
	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.Id, &album.Name, &album.ReleaseDate, &album.Img); err != nil {
			fmt.Printf("%v\n", err)
		}
		result = append(result, album)
	}

	return result
}

func GetAlbumById(id int) Album {
	db := GetDb()
	defer db.Close()
	row := db.QueryRow("select albums.id, albums.name, albums.release_date, albums.img from albums where albums.id = ? ", id)

	var album Album
	if err := row.Scan(&album.Id, &album.Name, &album.ReleaseDate, &album.Img); err != nil {
		fmt.Printf("%v\n", err)
	}

	return album
}
