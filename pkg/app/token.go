package app

import (
	"github.com/dgrijalva/jwt-go"
	"mikou/global"
	"time"
)

type Claims struct {
	UserId      int
	UserAccount string
	RoleId      int
	jwt.StandardClaims
}

// 生成token
func GenerateToken(UserAccount string, RoleId ,UserId int) (string, error) {
	expireTime := time.Now().Add(time.Duration(global.TokenSetting.Expires) * time.Second)
	claims := &Claims{
		UserAccount: UserAccount,
		RoleId:      RoleId,
		UserId :UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    global.TokenSetting.Issuer,  // 签名颁发者
			Subject:   global.TokenSetting.Subject, //签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.TokenSetting.Key))

}

// 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(global.TokenSetting.Key), nil
	})
	return token, Claims, err
}
