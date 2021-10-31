package handlers

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"server/middleware"
	"server/utils"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Uuid string `json:"uuid"`
	jwt.StandardClaims
}

func (h *Handlers) authorized(r *http.Request) *middleware.Student {
	bearer := r.Header.Get("Authorization")
	if len(bearer) == 0 {
		return nil
	}

	split := strings.Fields(bearer)
	uuidStr := decrypt(aes_key, split[1])
	uuid, _ := strconv.Atoi(uuidStr)
	if student := h.db.FindStudentUUID(uuid); student != nil {
		log.Printf(LOGGER_INFO_HELPERS+" Student info by uuid\n%+v", student)
		return student
	} else {
		return nil
	}
}

func encrypt(key []byte, text string) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(LOGGER_ERROR_HELPERS, err.Error())
		return ""
	}
	plaintext := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, aes_key)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return encodeBase64(ciphertext)
}

func decrypt(key []byte, text string) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println(LOGGER_ERROR_HELPERS, err.Error())
		return ""
	}
	ciphertext := decodeBase64(text)
	cfb := cipher.NewCFBEncrypter(block, aes_key)
	plaintext := make([]byte, len(ciphertext))
	cfb.XORKeyStream(plaintext, ciphertext)
	return string(plaintext)
}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Println(LOGGER_ERROR_HELPERS, err.Error())
		return nil
	}
	return data
}

func (h *Handlers) googleRespDecoder(resp http.Response) middleware.GoogleUser {
	contents := []byte(utils.JsonifyHttpResponse(resp))

	var user middleware.GoogleUser
	json.Unmarshal(contents, &user)

	return user
}
