package token

import (
	"encoding/base64"
	"encoding/json"
	"github.com/hako/branca"
	"ssn-backend/utils/app_context"
)

func GenerateToken(payload app_context.UserContext) (string, error) {
	b := generateBranca(120)

	bytes, err := json.Marshal(payload)
	sEnc := base64.StdEncoding.EncodeToString(bytes)
	token, err := b.EncodeToString(sEnc)

	return token, err
}

func GenerateRefreshToken(userContext app_context.UserContext) (string, error) {
	b := generateBranca(604800)

	payload := refreshToken{IDUser: userContext.IDUser}

	bytes, err := json.Marshal(payload)
	sEnc := base64.StdEncoding.EncodeToString(bytes)
	refreshToken, err := b.EncodeToString(sEnc)

	return refreshToken, err
}

func ValidateToken(token string) (app_context.UserContext, error) {
	context := app_context.UserContext{}
	b := generateBranca(120)

	payload, err := b.DecodeToString(token)
	bytes, err := base64.StdEncoding.DecodeString(payload)
	err = json.Unmarshal(bytes, &context)

	return context, err
}

func ValidateRefreshToken(rToken string) (uint64, error) {
	context := refreshToken{}
	b := generateBranca(604800)

	payload, err := b.DecodeToString(rToken)
	bytes, err := base64.StdEncoding.DecodeString(payload)
	err = json.Unmarshal(bytes, &context)

	return context.IDUser, err
}

func generateBranca(durationSeconds uint32) *branca.Branca {
	b := branca.NewBranca(getKey())
	b.SetTTL(durationSeconds)
	return b
}

func getKey() string {
	return "supersecretkeyyoushouldnotcommit"
}

type refreshToken struct {
	IDUser uint64
}
