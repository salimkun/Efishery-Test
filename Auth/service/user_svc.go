package service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/salimkun/Efishery-Test/Auth/common/util"
	"github.com/salimkun/Efishery-Test/Auth/model"
	"github.com/salimkun/Efishery-Test/Auth/repository"
)

func RegisterUser(c *gin.Context) {
	var request model.User

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validation phone number
	users, err := repository.ReadFile()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	for _, i := range users {
		if i.Phone == request.Phone {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "phone number already register"})
			return
		}
	}

	request.Password = util.GeneratePassword()
	request.Registered = time.Now().String()

	e := repository.CreateUser(&request)
	if e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": e})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": request})
}

func LoginUser(c *gin.Context) {
	var request model.User

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validation phone number & password
	users, err := repository.ReadFile()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	for _, i := range users {
		// generate jwt and return
		if i.Phone == request.Phone && i.Password == request.Password {
			token, err := util.GenerateToken(i.Name, i.Phone, i.Registered, i.Role)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error CUA": err})
				return
			}
			c.JSON(http.StatusOK, gin.H{"data": model.JwtToken{Token: token}})
			return
		}
	}

	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "validation failed please check phone and password"})
}

func GetUserByToken(c *gin.Context) {

	data, err := util.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": data})
}
