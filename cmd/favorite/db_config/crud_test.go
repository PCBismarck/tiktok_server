package dbconfig

import (
	"log"
	"testing"
)

func TestGetFavoriteList(t *testing.T) {
	InitDB()
	output, err := GetFavoriteList(1)
	log.Println(output)
	log.Println(err)
}

func TestCreateFavorite(t *testing.T) {
	InitDB()
	output := CreateFavorite(1, 13, "3x", "2x", "tx")
	log.Println(output)
}

func TestDeleteFavorite(t *testing.T) {
	InitDB()
	output := DeleteFavorite(1, 13)
	log.Println(output)
}

func TestGetVideoByVid(t *testing.T) {
	InitDB()

}
