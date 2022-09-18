package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/salimkun/Efishery-Test/Auth/model"
)

func GenerateToken(name, phone, registered_at string, role_id int32) (string, error) {

	// token_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))

	// if err != nil {
	// 	return "", err
	// }

	claims := jwt.MapClaims{}
	claims["Authorized"] = true
	claims["name"] = name
	claims["phone"] = phone
	claims["registered_at"] = registered_at
	claims["role_id"] = role_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(2)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenID(c *gin.Context) (*model.UserClaims, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		role_id, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["role_id"]), 10, 32)
		if err != nil {
			return nil, err
		}

		return &model.UserClaims{
			Name:       fmt.Sprintf("%s", claims["name"]),
			Phone:      fmt.Sprintf("%s", claims["phone"]),
			Registered: fmt.Sprintf("%s", claims["registered_at"]),
			Role:       int32(role_id),
		}, nil

	}
	return nil, nil
}
