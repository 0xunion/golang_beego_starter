package main

import (
	"fmt"
	"time"

	"github.com/0xunion/exercise_back/src/util/auth"
)

type LoginEmail struct {
	Email   string
	Captcha string
}

func main() {
	first_token := auth.NewAuthToken(LoginEmail{
		Email:   "example@xx.com",
		Captcha: "123456",
	})

	token_string := first_token.GenerateToken(222)

	second_token := auth.NewAuthTokenFromToken[LoginEmail](token_string)
	fmt.Println(second_token.Check(time.Now().Unix() + 1))
	fmt.Println(second_token.Info())
}
