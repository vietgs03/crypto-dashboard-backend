package common

import (
	"encoding/json"
	"time"

	"crypto-dashboard/pkg/response"

	jose "github.com/go-jose/go-jose/v4"
	"github.com/golang-jwt/jwt"

	josejwt "github.com/go-jose/go-jose/v4/jwt"
)

type JWTClaims struct {
	Eu  any   `json:"eu"`
	Exp int64 `json:"exp"`
}

type SecretKey struct {
	TokenSecret   string
	EncryptSecret string
}

func GenerateSignedToken(
	tokenSecret, encryptSecret string,
	payload *JWTClaims,
) (string, *response.AppError) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"eu":  payload,
			"exp": time.Now().Add(time.Minute * 5).Unix(),
		})

	signedToken, err := jwtToken.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", response.UnknownError(err.Error())
	}

	encrypter, err := jose.NewEncrypter(
		jose.A128GCM,
		jose.Recipient{Algorithm: jose.DIRECT, Key: []byte(encryptSecret)},
		(&jose.EncrypterOptions{}).WithType("JWT"),
	)
	if err != nil {
		return "", response.UnknownError(err.Error())
	}

	encryptedToken, err := josejwt.Encrypted(encrypter).Claims(signedToken).Serialize()
	if err != nil {
		return "", response.UnknownError(err.Error())
	}

	return encryptedToken, nil
}

func VerifySignedToken[T any](
	TokenSecret, encryptSecret, encryptedToken string,
) (*T, *response.AppError) {
	tok, err := josejwt.ParseEncrypted(encryptSecret, []jose.KeyAlgorithm{jose.DIRECT}, []jose.ContentEncryption{jose.A128GCM})
	if err != nil {
		return nil, response.Unauthorization("invalid token")
	}

	out := ""
	if err := tok.Claims([]byte(encryptedToken), &out); err != nil {
		return nil, response.Unauthorization("invalid token")
	}

	jwtToken, err := jwt.Parse(out, func(token *jwt.Token) (any, error) {
		return []byte(TokenSecret), nil
	})
	if err != nil {
		return nil, response.Unauthorization("invalid token")
	}

	if !jwtToken.Valid {
		return nil, response.Unauthorization("invalid token")
	}

	var payload T

	jsonData, _ := json.Marshal(jwtToken.Claims.(jwt.MapClaims)["eu"])
	_ = json.Unmarshal(jsonData, &payload)

	return &payload, nil
}
