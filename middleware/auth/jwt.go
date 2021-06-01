package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	SECRETKEY = "243223ffslsfsldfl412fdsfsdf" //私钥
)

// CustomClaims  struct
type CustomClaims struct {
	TokenInfo TokenInfo
	jwt.StandardClaims
}
type TokenInfo struct {
	UserId        string `valid:"-"`
	LoginDeviceId string `valid:"-"`
	RoleGroupId   string `valid:"-"`
	PhoneNumber   int64
}

// GenToken
// TokenInfo
// return (token string, err error)
func GenToken(t *TokenInfo) (string, error) {
	maxAge := 60 * 60 * 24
	//HMAC SHA256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		TokenInfo: TokenInfo{
			UserId:        t.UserId,
			RoleGroupId:   t.RoleGroupId,
			LoginDeviceId: t.LoginDeviceId,
			PhoneNumber:   t.PhoneNumber,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), //Expires Time
			Issuer:    "Gateway",
		},
	})
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//ParseToken
func ParseToken(tokenString string) (*CustomClaims, error) {
	if len(tokenString) < 5 {
		return nil, errors.New("token invalid auth token")
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{
		TokenInfo:      TokenInfo{},
		StandardClaims: jwt.StandardClaims{},
	}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

// Get UserId From Token
func UserClaimFromToken(tokenInfo TokenInfo) string {
	return tokenInfo.UserId
}
