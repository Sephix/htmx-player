package data

import (
	"fmt"
)

func GetSongById(id int) []byte {
	db := GetDb()
	defer db.Close()
	row := db.QueryRow("select song from songs where id = ?", id)
	var song []byte
	if err := row.Scan(&song); err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	return song
}
