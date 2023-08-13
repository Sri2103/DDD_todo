package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("my_secret")

type Claims struct {
	Username string `json:"username"` // this is the username of user who has logged in.
	UserID   string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID, username string)(string,error){
	claims := &Claims{
		UserID: userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour *24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(jwtSecret)
}

func VerifyToken(tokenString string)(*Claims,error) {
	token,err := jwt.ParseWithClaims(tokenString,&Claims{},func (token *jwt.Token)(interface{},error)  {
		return jwtSecret,nil
	})
	if err != nil {
		return nil,err
	}
	if claims,ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrInvalidKey
}