// Package bni This library is given by BNI, with addition comment and fix linter
package bni

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

const timeDiffLimit = 480

// Encrypt encrypting json string data
func Encrypt(json string, clientID string, secretKey string) string {
	return doubleEncrypt(reverse(fmt.Sprintf("%v", time.Now().Unix()))+"."+json, clientID, secretKey)
}

// Decrypt decrypting encrypted string to be readable data
func Decrypt(encrypted string, clientID string, secretKey string) (string, error) {
	parsedString := doubleDecrypt(encrypted, clientID, secretKey)
	var lst = strings.SplitN(parsedString, ".", 2)
	if len(lst) < 2 {
		return "", errors.New("bniEnc: parsing error, wrong cid or sck or invalid data")
	}
	if tsDiff(reverse(lst[0])) == false {
		return "", errors.New("bniEnc: data has been expired")
	}
	return lst[1], nil
}

func tsDiff(ts string) bool {
	_ts, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return false
	}
	return math.Abs(float64(_ts-time.Now().Unix())) <= timeDiffLimit
}

func doubleEncrypt(str string, cid string, sck string) string {
	arr := []byte(str)
	result := encrypt(arr, cid)
	result = encrypt(result, sck)
	return strings.Replace(strings.Replace(strings.TrimRight(base64.StdEncoding.EncodeToString(result), "="), "+", "-", -1), "/", "_", -1)
}

func encrypt(str []byte, k string) []byte {
	var result []byte
	strls := len(str)
	strlk := len(k)
	for i := 0; i < strls; i++ {
		char := str[i]
		keychar := k[(i+strlk-1)%strlk]
		char = byte((int(char) + int(keychar)) % 128)
		result = append(result, char)
	}
	return result
}

func doubleDecrypt(str string, cid string, sck string) string {
	if i := len(str) % 4; i != 0 {
		str += strings.Repeat("=", 4-i)
	}
	result, err := base64.StdEncoding.DecodeString(strings.Replace(strings.Replace(str, "-", "+", -1), "_", "/", -1))
	if err != nil {
		return ""
	}
	result = decrypt(result, cid)
	result = decrypt(result, sck)
	return string(result[:])
}

func decrypt(str []byte, k string) []byte {
	var result []byte
	strls := len(str)
	strlk := len(k)
	for i := 0; i < strls; i++ {
		char := str[i]
		keychar := k[(i+strlk-1)%strlk]
		char = byte(((int(char) - int(keychar)) + 256) % 128)
		result = append(result, char)
	}
	return result
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
