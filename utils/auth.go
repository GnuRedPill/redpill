package utils

import (
	"bytes"
	"crypto/des"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"
)

const AUTH_KEY = "cgi_bbs_"

type AuthTool struct {
}

func (a AuthTool) Gen(pass string) string {
	brr := sha512.Sum384([]byte(pass))
	return hex.EncodeToString(brr[:])
}

func (a AuthTool) Aes(id int, user, pass string) string {
	ids := strconv.Itoa(id)
	token := map[string]string{}
	token["id"] = ids
	token["user"] = user
	token["pass"] = pass
	brr, err := json.Marshal(token)
	if err != nil {
		log.Println(err)
	}
	m, err := Encrypt(brr, []byte(AUTH_KEY))
	return hex.EncodeToString(m)
}

func (a AuthTool) CheckAes(id int, user, pass, cookie string) bool {
	return strings.EqualFold(a.Aes(id, user, pass), cookie)
}

func (a AuthTool) UnAes(cookie string) int {
	brr, _ := hex.DecodeString(cookie)
	s, _ := Decrypt(brr, []byte(AUTH_KEY))
	res := map[string]string{}
	err := json.Unmarshal(s, &res)
	if err != nil {
		log.Println(err)
		return 0
	}
	ids, ok := res["id"]
	if ok {
		id, err := strconv.Atoi(ids)
		if err != nil {
			return 0
		}
		return id
	} else {
		return 0
	}
}

func Encrypt(data, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	data = pkcs5Padding(data, bs)
	if len(data)%bs != 0 {
		return nil, errors.New("need a multiple of the block size")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	return out, nil
}
func Decrypt(data []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = pkcs5UnPadding(out)
	return out, nil
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}
