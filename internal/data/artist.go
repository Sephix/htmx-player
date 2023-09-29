package data

import (
	"fmt"
)

type Artist struct {
	Id       int64
	Name     string
	Img      int64
	DeezerID int64
}

func GetAllArtists(name string) []Artist {
	fmt.Println("Find artists with name: ", name)
	db := GetDb()
	rows, _ := db.Query("select artists.id, artists.name, artists.img, artists.deezer_id from artists where artists.name like ?", "%"+name+"%")

	defer rows.Close()

	result := make([]Artist, 0)
	for rows.Next() {
		var artist Artist
		if err := rows.Scan(&artist.Id, &artist.Name, &artist.Img, &artist.DeezerID); err != nil {
			fmt.Printf("%v\n", err)
		}
		result = append(result, artist)
	}

	return result
}

func GetArtistById(id int) Artist {
	db := GetDb()
	row := db.QueryRow("select artists.id, artists.name, artists.img from artists where artists.id = ?", id)

	var result Artist
	if err := row.Scan(&result.Id, &result.Name, &result.Img); err != nil {
		fmt.Printf("%v\n", err)
	}

	return result
}
