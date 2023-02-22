package redis_config

import (
	"log"
	"testing"
)

func TestCreateFavoriteAction(t *testing.T) {
	InitRedis()
	output, err := CreateFavoriteAction(111)
	log.Println(output)
	log.Println(err)
}

func TestDeleteFavoriteAction(t *testing.T) {

}
