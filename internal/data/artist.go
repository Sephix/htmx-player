package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Artist struct {
	Id       int64
	Name     string
	Img      int64
	DeezerID int64
}

type Album struct {
	Id          int64
	Name        string
	ReleaseDate string
	Img         int64
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

func GetAlbumByArtistId(id int, name string) []Album {
	db := GetDb()
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

type Response struct {
	Data []AlbumDeezer `json:"data"`
}

// A Pokemon Struct to map every pokemon to.
type AlbumDeezer struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Img   string `json:"cover_big"`
	Date  string `json:"release_date"`
}

// A struct to map our Pokemon's Species which includes it's name
type PokemonSpecies struct {
	Name string `json:"name"`
}

func InsertArtistsAlbums() {
	artists := GetAllArtists("")
	for _, artist := range artists {
		fmt.Println("---- ----- ----- -----")
		fmt.Println("---- Artist: , " + artist.Name + " -----")
		GetAblumsFromDeezer(artist.Id, artist.DeezerID)
	}
}

func GetAblumsFromDeezer(id int64, deezer_id int64) {
	url := fmt.Sprintf("https://api.deezer.com/artist/%v/albums", deezer_id)
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer response.Body.Close()

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	db := GetDb()
	defer db.Close()
	for _, album := range responseObject.Data {
		fmt.Println("---- ----- ----- -----")
		fmt.Println("---- Album: , " + album.Title + " -----")
		idAlbum, err := insertAblum(album)
		if err != nil {
			fmt.Println("Error for album: ", album.Title)
			fmt.Println(err.Error())
			break
		}
		_, err = db.Exec("insert into artists_albums (album_id, artist_id) values  (?, ?) ", idAlbum, id)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
}

func insertAblum(album AlbumDeezer) (int64, error) {
	db := GetDb()
	defer db.Close()
	idImg := InsertImg(album.Img)

	result, err := db.Exec("insert into albums (name, release_date, img, deezer_id) values  (?, ?, ?, ?) ", album.Title, album.Date, idImg, album.Id)
	if err != nil {
		return 0, err
	}
	idAlbum, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return idAlbum, nil
}
