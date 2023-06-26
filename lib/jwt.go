package lib

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/miajio/www/model"
)

type JwtUserInfoClaims struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Status   int    `json:"status"`
	jwt.RegisteredClaims
}

func (JwtUserInfoClaims) Valid() error {
	return nil
}

type jwtImpl struct{}

type jwtInterface interface {
	GenerateToken([]byte, model.UserInfoModel, time.Duration) (string, error)
	ParseToken([]byte, string) (JwtUserInfoClaims, error)
	IsTokenValid([]byte, string) bool
}

// Jwt操作
var Jwt jwtInterface = (*jwtImpl)(nil)

func (*jwtImpl) GenerateToken(key []byte, userInfo model.UserInfoModel, timeOut time.Duration) (string, error) {
	jwtUserInfoClaims := JwtUserInfoClaims{
		Uid:      userInfo.Uid,
		Username: userInfo.Username,
		Email:    userInfo.Email,
		Status:   userInfo.Status,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(timeOut)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtUserInfoClaims)
	return tk.SignedString(key)
}

func (*jwtImpl) ParseToken(key []byte, token string) (JwtUserInfoClaims, error) {
	res := JwtUserInfoClaims{}

	parse, err := jwt.ParseWithClaims(token, &res, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil && !parse.Valid {
		return res, errors.New("invalid of token")
	}

	return res, err
}

func (j *jwtImpl) IsTokenValid(key []byte, token string) bool {
	_, err := j.ParseToken(key, token)
	if err != nil {
		return false
	}
	return true
}
