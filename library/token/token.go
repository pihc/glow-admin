package token

import (
	"github.com/dgrijalva/jwt-go"
	"glow-admin/app/model"
	"time"
)

// 指定加密密钥
var jwtSecret = []byte("reed-test")

// Claim是一些实体（通常指的用户）的状态和额外的元数据
type CustomClaims struct {
	Data model.ContextUser
	jwt.StandardClaims
}

// 解析token
func Parse(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}
	// 解密转换类型并返回
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// 生成token
func Generate(data model.ContextUser, expire int) (string, error) {
	claims := CustomClaims{
		Data: data,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "clean code",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Duration(expire) * time.Minute).Unix(),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tkn, err := jwtToken.SignedString(jwtSecret)
	return tkn, err
}
