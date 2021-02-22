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
// reurun (token string, err error)
func GenToken(t *TokenInfo) (string, error) {
	maxAge := 60 * 60 * 24
	//采用HMAC SHA256加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &CustomClaims{
		TokenInfo: TokenInfo{
			UserId:        t.UserId,
			RoleGroupId:   t.RoleGroupId,
			LoginDeviceId: t.LoginDeviceId,
			PhoneNumber:   t.PhoneNumber,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
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
	if  len(tokenString) < 5 {
		return nil, errors.New("token invalid auth token")
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

//从token中获取用户唯一标识
func UserClaimFromToken(tokenInfo TokenInfo) string {
	return tokenInfo.UserId
}
