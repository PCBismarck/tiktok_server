package dbconfig

import (
	"log"
	"testing"
)

func TestCreateComment(t *testing.T) {
	InitDB()
	output, err := CreateComment(1, 11, "1x")
	log.Println(err)
	log.Println(output)
}

func TestDeleteComment(t *testing.T) {
	InitDB()
	ok := DeleteComment(2, 2)
	log.Println(ok)
}

func TestGetCommentsByVid(t *testing.T) {
	InitDB()
	output, err := GetCommentsByVid(1)
	log.Println(output)
	log.Println(err)
}
