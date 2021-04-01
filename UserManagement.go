package user_management

import (
	"github.com/gin-gonic/gin"
	"github.com/golbanstefan/user-management/config"
	"github.com/golbanstefan/user-management/middleware"
	"github.com/golbanstefan/user-management/user"
)

//SetPathToKey in normal case the key will store in root folder keys, in key has another
// path or name you can set them
func SetPathToKey(path string) {
	config.Init().SetPath(path)
}

//ConfigFirebase  will init the communication with firebase,
// will return *auth.Client  and will set them for future work with them
// Also will register the firebase middleware to handle authentications
// Also function will register the routes for API use
func ConfigFirebase(r *gin.Engine) {
	// configure firebase
	firebaseAuth := config.Init().SetupFirebase()

	// set db & firebase auth to gin context with a middleware to all incoming request
	r.Use(func(c *gin.Context) {

		c.Set("firebaseAuth", firebaseAuth)
	})
	r.Use(middleware.AuthMiddleware)

	// Management
	u := Management()
	r.GET("/user", u.GetUser)
	r.GET("/users", u.GetAllUsers)
	r.PUT("/user", u.UpdateUser)
	r.POST("/user", u.CreateUser)
	r.DELETE("/user", u.DeleteUser)
	r.DELETE("/user-bulk", u.BulkDeleteUsers)
	r.POST("/user-uid", u.CreateUserWithUID)
	r.POST("/user-by-email", u.GetUserByEmail)
	r.PUT("/user-custom-claims", u.CustomClaimsSet)
}

//User interface will provide all necessary methods fo mange users
// the methods are ready to use as gin handle methods
type User interface {
	GetUser(c *gin.Context)
	GetUserByEmail(c *gin.Context)
	GetAllUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	CreateUserWithUID(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	BulkDeleteUsers(c *gin.Context)
	CustomClaimsSet(c *gin.Context)
}

//Management  return the user interface
func Management() User {
	var m user.User
	return m
}
