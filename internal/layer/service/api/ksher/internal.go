package ksher

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/rs/zerolog/log"
)

func toJson(data interface{}) map[string]interface{} {
	var pregen map[string]interface{}
	b, _ := json.Marshal(data)
	json.Unmarshal(b, &pregen)
	return pregen
}

func generateSignature(url string, data map[string]interface{}, token string) string {
	concat := ""
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if k == "signature" {
			continue
		}
		concat = fmt.Sprintf("%s%s%v", concat, k, data[k])
	}

	params := url + concat
	log.Debug().Str("path", url).Str("params", params).Msg("generate signature")

	// Create a new HMAC by defining the hash type and the key (as byte array)
	hash := hmac.New(sha256.New, []byte(token))

	// Write Data to it
	hash.Write([]byte(params))

	// to hexits
	signature := strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))

	return signature
}
