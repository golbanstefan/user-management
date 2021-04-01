# Firebase User Management
## Golang package for user management in Google  Firebase 

This  is ready to use package for user management via standard json API
 - package use Gin Web Framework [https://github.com/gin-gonic/gin] for handle API request and authentication
 - has built in middleware
 - allow tu use firebase auth in your application
 
## Installation
 ```sh
  go get github.com/golbanstefan/user-management
 ```
## Methods
  	GetUser(c *gin.Context)
	GetUserByEmail(c *gin.Context)
	GetAllUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	CreateUserWithUID(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	BulkDeleteUsers(c *gin.Context)
	CustomClaimsSet(c *gin.Context)

## Used routes
    r.GET("/user", u.GetUser)
	r.GET("/users", u.GetAllUsers)
	r.PUT("/user", u.UpdateUser)
	r.POST("/user", u.CreateUser)
	r.DELETE("/user", u.DeleteUser)
	r.DELETE("/user-bulk", u.BulkDeleteUsers)
	r.POST("/user-uid", u.CreateUserWithUID)
	r.POST("/user-by-email", u.GetUserByEmail)
	r.PUT("/user-custom-claims", u.CustomClaimsSet)

## Examples
### Example to use
```go 
import (
	"github.com/gin-gonic/gin"
	user_management "github.com/golbanstefan/user-management"
	"sync"
)

type IRouter interface {
	Start() *gin.Engine
}
type router struct{}

func (router router) Start() *gin.Engine {
	// initialize new gin engine (for server)
	r := gin.Default()
	// using the gin router register api requests and middleware
	user_management.ConfigFirebase(r)
	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func Router() IRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}
func main() {
	err := http.ListenAndServe(":"+"8080", Router().Start())
	errors.CheckError(err)

}

```

### Access to firebase auth from *gin.Context 
```go
authI, exist := c.Get("firebaseAuth")
	if !exist {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Get("Firebase auth is loosed"))
		c.Abort()
		return nil
	}
```
