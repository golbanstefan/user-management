package utils

import (
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/golbanstefan/user-management/response"
)

func GetAuthClient(c *gin.Context) *auth.Client {
	authI, exist := c.Get("firebaseAuth")
	if !exist {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Get("Firebase auth is loosed"))
		c.Abort()
		return nil
	}
	return authI.(*auth.Client)

}
