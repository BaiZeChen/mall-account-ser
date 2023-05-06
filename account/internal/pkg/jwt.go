package pkg

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	secret     = "EuKizqESuZONTiokl0zHer9tJLMa8tDb"
	expireTime = 3600
)

type JWTClaims struct {
	jwt.StandardClaims
	AccountId   uint32
	AccountName string
}

func (j *JWTClaims) setExpiredAt() {
	j.ExpiresAt = time.Now().Add(time.Second * time.Duration(expireTime)).Unix()
}

func (j *JWTClaims) Generate() (string, error) {
	j.Issuer = "account-ser"
	j.IssuedAt = time.Now().Unix()
	j.Subject = "auth"
	j.setExpiredAt()

	res := jwt.NewWithClaims(jwt.SigningMethodHS256, j)
	token, err := res.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (j *JWTClaims) Check(token string) (bool, error) {
	res, err := jwt.ParseWithClaims(token, j, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}

	if claims, ok := res.Claims.(*JWTClaims); ok && res.Valid {
		now := time.Now().Unix()
		if now >= claims.ExpiresAt {
			return false, nil
		}
	} else {
		return false, nil
	}
	return true, nil
}

func (j *JWTClaims) Refresh(token string) error {
	res, err := jwt.ParseWithClaims(token, j, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return err
	}

	if claims, ok := res.Claims.(*JWTClaims); ok && res.Valid {
		claims.setExpiredAt()
		return nil
	} else {
		return errors.New("token校验失败，无法刷新")
	}
}
