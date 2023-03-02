package main

import (
	"fmt"

	"github.com/0xunion/exercise_back/src/util/auth"
	"github.com/0xunion/exercise_back/src/util/hash"
)

func main() {
	fmt.Println(auth.HashPassword(hash.Md5("155709232943652")))
}
