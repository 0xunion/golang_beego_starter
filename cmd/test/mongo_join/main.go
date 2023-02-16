package main

import (
	"fmt"
	"time"

	"github.com/0xunion/exercise_back/src/model"
	"github.com/0xunion/exercise_back/src/types"
)

func main() {
	type test_user struct {
		Id   types.PrimaryId `json:"id" bson:"_id"`
		Name string          `json:"name" bson:"name"`
	}

	type test_post struct {
		Id      types.PrimaryId `json:"id" bson:"_id"`
		Uid     types.PrimaryId `json:"uid" bson:"uid"`
		Content string          `json:"content" bson:"content"`
	}

	type result struct {
		Id      types.PrimaryId `json:"id" bson:"_id"`
		User    test_user       `json:"user" bson:"user" join:"uid=_id"`
		Uid     types.PrimaryId `json:"uid" bson:"uid"`
		Content string          `json:"content" bson:"content"`
	}

	start_time := time.Now().UnixMilli()
	model.ModelGetAllJoin[test_post, result](model.NewMongoFilter())
	end_time := time.Now().UnixMilli()
	fmt.Println("no cache time:", end_time-start_time)
	start_time = time.Now().UnixMilli()
	model.ModelGetAllJoin[test_post, result](model.NewMongoFilter())
	end_time = time.Now().UnixMilli()
	fmt.Println("cache time:", end_time-start_time)
	start_time = time.Now().UnixMilli()
	model.ModelGetAllJoin[test_post, result](model.NewMongoFilter())
	end_time = time.Now().UnixMilli()
	fmt.Println("cache time:", end_time-start_time)

	start_time = time.Now().UnixMilli()
	model.ModelGet[test_post](model.NewMongoFilter())
	end_time = time.Now().UnixMilli()
	fmt.Println("normal time:", end_time-start_time)
}
