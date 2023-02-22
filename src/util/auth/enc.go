package auth

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/*
	enc.go provides temprary authrization services

	this is a series of services that require more than one request to complete
	such as login(1. captcha 2. login), register(1. email 2. register), etc.

	cause all of these services are not atomic, we need a token which contains
	required information like uid, correct captcha, etc. to make sure the information
	passed between requests are correct.
*/

var (
	auth_token_key string
)

func init() {
	rand.Seed(time.Now().UnixNano())

	// initialize auth_token_key with a random string
	auth_token_key = md5hash(strconv.FormatInt(rand.Int63(), 16))[0:16]
}

// AuthToken is a token that contains required information to complete a service
// T should be a struct customized for a service
type AuthToken[T any] struct {
	info   T
	expire int64
	token  string
	enc    func(src string, key string) string
	dec    func(src string, key string) string
}

func auth_enc(src string, key string) string {
	// aes encrypt
	enc, _ := Encrypt(src, []byte(key), []byte(key))
	return enc
}

func auth_dec(src string, key string) string {
	// aes decrypt
	dec, _ := Decrypt(src, []byte(key), []byte(key))
	return dec
}

// NewAuthToken creates a new AuthToken
func NewAuthToken[T any](info T) *AuthToken[T] {
	return &AuthToken[T]{
		info: info,
		enc:  auth_enc,
		dec:  auth_dec,
	}
}

// NewAuthTokenFromToken creates a new AuthToken from a token
func NewAuthTokenFromToken[T any](token string) *AuthToken[T] {
	return &AuthToken[T]{
		token: token,
		enc:   auth_enc,
		dec:   auth_dec,
	}
}

// Info returns the information contained in the token
func (t *AuthToken[T]) Info() T {
	return t.info
}

func md5hash(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// expire is the time duration in millseconds which marks how long the token is valid
func (t *AuthToken[T]) GenerateToken(expire int64) string {
	t.expire = time.Now().Unix() + expire

	random_key := md5hash(strconv.FormatInt(rand.Int63(), 16))
	real_key := md5hash(random_key + auth_token_key)[0:16]

	// token = random_key : enc(info, random_key) : expire
	t_json, _ := json.Marshal(t.info)
	token := random_key + ":" + t.enc(string(t_json), real_key) + ":" + strconv.FormatInt(t.expire, 10)
	// token = BASE64(token : md5(token + k))

	token = token + ":" + md5hash(token+auth_token_key)
	token = base64.StdEncoding.EncodeToString([]byte(token))

	t.token = token
	return t.token
}

// current is the current time in seconds
// Check will store the information contained in the token in t.info
func (t *AuthToken[T]) Check(current int64) bool {
	// base64 decode
	token, err := base64.StdEncoding.DecodeString(t.token)
	if err != nil {
		return false
	}

	// split token
	tokens := strings.Split(string(token), ":")
	if len(tokens) != 4 {
		return false
	}

	// check md5
	random_key := tokens[0]
	real_key := md5hash(random_key + auth_token_key)[0:16]
	if md5hash(tokens[0]+":"+tokens[1]+":"+tokens[2]+auth_token_key) != tokens[3] {
		return false
	}

	// check expire
	expire, err := strconv.ParseInt(tokens[2], 10, 64)
	if err != nil {
		return false
	}

	if current > expire {
		return false
	}
	// check info
	info := t.dec(tokens[1], real_key)
	if err := json.Unmarshal([]byte(info), &t.info); err != nil {
		return false
	}

	return true
}
