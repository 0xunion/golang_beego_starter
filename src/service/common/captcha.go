package common

import (
	"strconv"
	"sync"
	"time"

	"${package}/src/routine/cache"
	"${package}/src/types"
	"${package}/src/util/auth"
	"${package}/src/util/num"
	"${package}/src/util/strings"
)

/*
	ErrCode:
		-1001: generate captcha failed
*/

type Captcha struct {
	Type   int    `json:"type"`    // captcha type like image, sms, email etc.
	Method string `json:"method"`  // captcha method like mathmatical, image etc.
	Ans    string `json:"ans"`     // answer
	Expire int64  `json:"expire"`  // when expire(UNIX timestamp)
	MaxTry int    `json:"max_try"` // max try times
	Id     string `json:"id"`      // captcha id(random string)
}

const (
	CAPTCHA_TYPE_MATHMATICAL_IMAGE = 1
	CAPTCHA_TYPE_MATHMATICAL_SMS   = 2

	CAPTCHA_METHOD_LOGIN = "login"

	CAPTCHA_EXPIRE_TIME = 60 * 5
	CAPTCHA_MAX_TRY     = 3
)

func (c *Captcha) IsExpired() bool {
	return c.Expire < time.Now().Unix()
}

func (c *Captcha) IsMaxTry(try int) bool {
	return try > c.MaxTry
}

func (c *Captcha) IsCorrect(ans string) bool {
	return c.Ans == ans
}

func (c *Captcha) GetId() string {
	return c.Id
}

// compare to IsCorrect, this function will check the captcha is expired or not
// and store the try times in cache at the same time.
var try_locker = sync.Mutex{}

func (c *Captcha) Try(ans string) bool {
	// store try times
	try_locker.Lock()
	defer try_locker.Unlock()
	times, err := cache.GetCache[int](c.Id, "captcha", "try_times")
	if err != nil {
		*times = 0
	}
	*times++
	err = cache.SetCache(c.Id, *times, "captcha", "try_times")
	if err != nil {
		return false
	}

	// check captcha is expired or not
	if c.IsExpired() {
		return false
	}

	// check try times
	if c.IsMaxTry(*times) {
		return false
	}

	// check answer
	success := c.IsCorrect(ans)
	if success {
		// set try times to max try times, so that the captcha will be expired
		*times = c.MaxTry
		err = cache.SetCache(c.Id, *times, "captcha", "try_times")
		if err != nil {
			return false
		}
	}

	return success
}

// NewMathmaticalImageCaptcha create a new mathmatical image captcha and
// return the captcha id and the base64 encoded image.
func newMathmaticalImageCaptcha(method string) (*Captcha, string) {
	//生成验证码
	var exp string = ""
	//生成两个数用于计算
	x, y, ans := num.Random(1, 30), num.Random(1, 30), ""
	switch num.Random(0, 3) {
	//乘法
	case 0:
		ary := []string{strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(x * y)}
		index := num.Random(0, 2)
		ans = ary[index]
		ary[index] = "?"
		exp = ary[0] + "*" + ary[1] + "=" + ary[2]
	//除法
	case 1:
		ary := []string{strconv.Itoa(x * y), strconv.Itoa(y), strconv.Itoa(x)}
		index := num.Random(0, 2)
		ans = ary[index]
		ary[index] = "?"
		exp = ary[0] + "/" + ary[1] + "=" + ary[2]
	//加法
	case 2:
		ary := []string{strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(x + y)}
		index := num.Random(0, 2)
		ans = ary[index]
		ary[index] = "?"
		exp = ary[0] + "+" + ary[1] + "=" + ary[2]
	//减法
	case 3:
		ary := []string{strconv.Itoa(x + y), strconv.Itoa(y), strconv.Itoa(x)}
		index := num.Random(0, 2)
		ans = ary[index]
		ary[index] = "?"
		exp = ary[0] + "-" + ary[1] + "=" + ary[2]
	}

	captcha, ok := auth.GenerateImageFromText(exp, 30)
	if !ok {
		return nil, ""
	}

	return &Captcha{
		Type:   CAPTCHA_TYPE_MATHMATICAL_IMAGE,
		Method: method,
		Ans:    ans,
		Expire: time.Now().Unix() + CAPTCHA_EXPIRE_TIME,
		MaxTry: CAPTCHA_MAX_TRY,
		Id:     strings.RandomAlphaString(16),
	}, captcha
}

func LoginCaptchaService() *types.MasterResponse {
	captcha, captchaImg := newMathmaticalImageCaptcha(CAPTCHA_METHOD_LOGIN)
	if captcha == nil {
		return types.ErrorResponse(-1001, "generate captcha failed")
	}

	auth_token := auth.NewAuthToken(captcha)
	var data struct {
		Token   string `json:"token"`
		Captcha string `json:"captcha"`
	}
	data.Token = auth_token.GenerateToken(300)
	data.Captcha = captchaImg

	return types.SuccessResponse(data)
}
