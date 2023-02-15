package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"time"

	"${package}/src/const/conf"
	"${package}/src/types"
)

var seucre_token_padding string
var secure_aes_key, secure_aes_iv []byte

func init() {
	seucre_token_padding = conf.SecureTokenPadding()
	_secure_aes_key := conf.SecureAesKey()
	_secure_aes_iv := conf.SecureAesIv()
	secure_aes_key = []byte(_secure_aes_key)
	secure_aes_iv = []byte(_secure_aes_iv)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesCBCEncrypt(rawData, key []byte, iv []byte) ([]byte, error) {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()

	//fmt.Println(blockSize)
	//fmt.Println(len(secure_aes_iv))

	rawData = PKCS5Padding(rawData, blockSize)
	cipherText := make([]byte, len(rawData))

	mode := cipher.NewCBCEncrypter(block, iv)

	mode.CryptBlocks(cipherText, rawData)

	return cipherText, nil
}

func AesCBCDecrypt(encryptData, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, errors.New("秘钥错误")
	}

	blockSize := block.BlockSize()

	if len(encryptData)%blockSize != 0 {
		return []byte{}, errors.New("密文长度错误")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(encryptData, encryptData)
	//解填充
	encryptData = PKCS5UnPadding(encryptData)
	return encryptData, nil
}

func Encrypt(str string, key []byte, iv []byte) (string, error) {
	data, err := AesCBCEncrypt([]byte(str), key, iv)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func Decrypt(str string, key []byte, iv []byte) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	dnData, err := AesCBCDecrypt(data, key, iv)
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}

// 将被加密后储存在JSONInfo的Token中
type Session struct {
	Hash       string          `json:"hash"`
	Uid        types.PrimaryId `json:"uid"`
	Login_time uint32          `json:"login_time"`
}

type LoginToken struct {
	EncrypedToken string
	Session       Session
}

type JSONInfo struct {
	Uid        types.PrimaryId `json:"uid"`
	Hash       string          `json:"hash"`
	Login_time uint32          `json:"login_time"`
	Token      string          `json:"token"`
}

func md5x(str string) string {
	str = seucre_token_padding + str + seucre_token_padding
	md5ctx := md5.New()
	md5ctx.Write([]byte(str))
	return hex.EncodeToString(md5ctx.Sum(nil))
}

func HashPassword(pass string) string {
	return md5x(md5x(pass) + seucre_token_padding)
}

func getUserKey(uid types.PrimaryId) []byte {
	return []byte(md5x(uid.Hex()))[:16]
}

func NewAuthTokenWithUid(uid types.PrimaryId) *LoginToken {
	var token LoginToken
	token.Session.Uid = uid
	return &token
}

func NewAuthTokenWithToken(token string) *LoginToken {
	var login_token LoginToken
	login_token.EncrypedToken = token
	return &login_token
}

func (c *LoginToken) GenerateToken(timestamp uint32) string {
	random, _ := rand.Int(rand.Reader, big.NewInt(9000))
	hash := md5x(random.Add(random, big.NewInt(1000)).String() + time.Now().String())

	c.Session.Hash = hash

	session_json, _ := json.Marshal(c.Session)
	user_key := getUserKey(c.Session.Uid)
	token, err := Encrypt(string(session_json), user_key, secure_aes_iv)
	if err != nil {
		return ""
	}

	var json_info JSONInfo
	json_info.Hash = hash
	json_info.Login_time = timestamp
	json_info.Uid = c.Session.Uid
	json_info.Token = token

	json_info_str, _ := json.Marshal(json_info)
	cookie := base64.RawStdEncoding.EncodeToString(json_info_str)

	cookie, err = Encrypt(cookie, secure_aes_key, secure_aes_iv)
	if err != nil {
		return ""
	}

	return cookie
}

func (c *LoginToken) AnalysisToken() bool {
	//初步解码得到token的整体信息
	decrypted, err := Decrypt(c.EncrypedToken, secure_aes_key, secure_aes_iv)
	if err != nil {
		return false
	}
	//这里有一个坑，Decrypt出来的字符串由于长度不一定跟4对齐，尤其是Decrypt中也使用了base64，导致了=替代的0消失，我们需要补全，所以这补一下=
	if len(decrypted)%4 != 0 {
		decrypted += string(bytes.Repeat([]byte{61}, 4-len(decrypted)%4))
	}

	//获取base64之前的json，并解析
	json_buf, err := base64.StdEncoding.DecodeString(decrypted)
	if err != nil {
		return false
	}

	var json_info JSONInfo
	err = json.Unmarshal(json_buf, &json_info)
	if err != nil {
		return false
	}

	//校验json合法性
	if json_info.Uid.IsZero() || len(json_info.Hash) != 32 || json_info.Login_time <= 1625731070 || json_info.Token == "" {
		return false
	}

	//解密token，并校验token是否被篡改
	user_key := getUserKey(json_info.Uid)
	token_buf, err := Decrypt(json_info.Token, user_key, secure_aes_iv)
	if err != nil {
		return false
	}

	var session Session
	err = json.Unmarshal([]byte(token_buf), &session)

	if err != nil {
		return false
	}

	if session.Hash != json_info.Hash || session.Uid != json_info.Uid {
		return false
	}

	c.Session.Hash = json_info.Hash
	c.Session.Login_time = json_info.Login_time
	c.Session.Uid = json_info.Uid

	return true
}

func (c *LoginToken) GetUid() types.PrimaryId {
	return c.Session.Uid
}

func GenerateCSRFTokenBySession(session string) string {
	return md5x(seucre_token_padding + session)
}

func CheckCSRFAvaliable(jct string, start_dash string) bool {
	return GenerateCSRFTokenBySession(start_dash) == jct
}
