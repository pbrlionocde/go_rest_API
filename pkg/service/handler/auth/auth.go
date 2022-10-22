package auth

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"crypto/sha1"
	"encoding/base64"
	"f_gin/pkg/service/logger"
	"f_gin/pkg/storage/db_conn"
	"f_gin/pkg/storage/models"
	"net/http"
)

var errorLogger *log.Logger
var dbConn *gorm.DB

func init() {
	dbConn = db_conn.GetDbConnection()
	errorLogger = logger.GetErrorLogger()
}

type User struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"  binding:"required"`
	Email     string `json:"email"      binding:"required,email"`
	Password  string `json:"password"   binding:"required"`
}

func PasswordToHash(password string) string {
	bytePassword := []byte(password)
	hasher := sha1.New()
	hasher.Write(bytePassword)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func CreateUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		errorResp := logger.ErrorToJSON(err)
		errorLogger.Println(err)
		c.IndentedJSON(http.StatusBadRequest, errorResp)
		return
	}
	err := dbConn.Create(
		&models.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Password:  PasswordToHash(user.Password),
		},
	)
	if err != nil {
		errorLogger.Println(err)
		return
	}
	var dbUser *models.User
	dbConn.First(&dbUser, "email = ?", user.Email)
	c.IndentedJSON(http.StatusCreated, dbUser)
}
