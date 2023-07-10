package authentication

import (
	"encoding/base64"
)

func Decode(text_to_decode string) string {
	base64_decode_result, err := base64.RawStdEncoding.DecodeString(text_to_decode)

	if err == nil {
		return string(base64_decode_result)
	} else {
		return ""
	}
}

func Encode(text_to_encode string) string {
	base_64_encode_result := base64.RawStdEncoding.EncodeToString([]byte(text_to_encode))

	return base_64_encode_result
}
