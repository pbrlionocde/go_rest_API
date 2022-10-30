package auth

import (
	"f_gin/pkg/service/config_loader"
	"f_gin/pkg/service/logger"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func getExpirationTime() time.Time {
	var expirationTime time.Time
	expirationTime = time.Now()
	expirationTime = expirationTime.AddDate(0, 0, 30)
	location, err := time.LoadLocation(config_loader.GetJwtYamlConfig().TimeZone)
	if err != nil {
		panic(err)
	}
	expirationTime = expirationTime.In(location)
	fmt.Println(expirationTime)
	return expirationTime
}

func CreateToken(email string, role string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["role"] = role
	claims["expirationTime"] = getExpirationTime()
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_STRING")))
	if err != nil {
		panic(err)
	}
	return tokenString
}

type LoginEmail struct {
	Email string `json:"email" binding:"required"`
}

func GetToken(c *gin.Context) {
	var loginEmail LoginEmail
	err := c.ShouldBindJSON(&loginEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": logger.ErrorToJSON((err))})
		return
	}
	token := CreateToken(loginEmail.Email, "User")
	c.IndentedJSON(http.StatusOK, token)
}

type TokenUser struct {
	Email          string `json:"email"`
	Role           string `json:"role"`
	ExpirationTime string `json:"expirationTime"`
	jwt.StandardClaims
}

func ParseToken(tokenString string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenUser{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_STRING")), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenUser); ok && token.Valid {
		fmt.Printf("%v %v", claims.Email, claims.StandardClaims.ExpiresAt)
		return claims, nil
	} else {
		return nil, err
	}
}
