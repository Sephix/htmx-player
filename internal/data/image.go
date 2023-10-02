package data

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"net/http"
)

func GetImageById(id int) image.Image {
	db := GetDb()
	row := db.QueryRow("select img from images where id = ?", id)
	var img []byte
	if err := row.Scan(&img); err != nil {
		fmt.Printf("%v\n", err)
		return nil
	}
	result, _ := jpeg.Decode(bytes.NewReader(img))
	return result
}

func InsertImg(imgUrl string) int64 {
	db := GetDb()
	resp, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println(err)
	}
	img, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	result, err := db.Exec("insert into images (img) values  (?) ", img)
	id, _ := result.LastInsertId()
	return id
}
