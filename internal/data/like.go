package data

import "fmt"

func GetLikeBySongId(id int64) bool {
	db := GetDb()
	row := db.QueryRow("select true from likes where song_id = ?", id)
	var liked bool
	if err := row.Scan(&liked); err != nil {
		return false
	}
	return liked
}

func InsertLike(trackID int64) bool {
	db := GetDb()
	_, err := db.Exec("INSERT INTO likes (song_id) VALUES (?)", trackID)
	if err != nil {
		fmt.Printf("Could not like song with id %v\n", err)
		fmt.Printf("%v\n", err)
		return false
	}
	return true
}
func DeleteLike(trackID int64) bool {
	db := GetDb()
	_, err := db.Exec("delete from likes where song_id = ?", trackID)
	if err != nil {
		fmt.Printf("Could not not delete liked song with id %v\n", err)
		fmt.Printf("%v\n", err)
		return true
	}
	return false
}
func IsTrackLiked(trackID int64) bool {
	return GetLikeBySongId(trackID)
}

func UpdateLike(trackID int64) bool {
	if IsTrackLiked(trackID) {
		return DeleteLike(trackID)
	}
	return InsertLike(trackID)
}
