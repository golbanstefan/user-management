package user

import (
	"context"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golbanstefan/user-management/errors"
	"github.com/golbanstefan/user-management/response"
	"github.com/golbanstefan/user-management/utils"
	"google.golang.org/api/iterator"
	"log"
	"net/http"
)

type User struct {
	Email         string `json:"email"     binding:"required,email"`
	EmailVerified *bool  `json:"email_verified" binding:"required"`
	PhoneNumber   string `json:"phone_number,omitempty"`
	Password      string `json:"password" binding:"required"`
	DisplayName   string `json:"display_name,omitempty" `
	PhotoURL      string `json:"photo_url,omitempty"`
	Disabled      *bool  `json:"disabled,omitempty" `
	UID           string `json:"uid,omitempty"`
}
type ModelUpdate struct {
	Email         string `json:"email,omitempty"`
	EmailVerified *bool  `json:"email_verified,omitempty"`
	PhoneNumber   string `json:"phone_number,omitempty"`
	Password      string `json:"password,omitempty"`
	DisplayName   string `json:"display_name,omitempty" `
	PhotoURL      string `json:"photo_url,omitempty"`
	Disabled      *bool  `json:"disabled,omitempty" `
	UID           string `json:"uid,omitempty" binding:"required"`
}

//GetUser
func (m User) GetUser(c *gin.Context) {

	// [END create_user_golang]
	uuid, exist := c.Get("UUID")
	if !exist {
		c.JSON(http.StatusBadRequest, response.Get("UUID is absent"))
		return
	}
	ac := utils.GetAuthClient(c)
	u, err := ac.GetUser(context.Background(), uuid.(string))
	errors.CheckError(err)
	c.JSON(http.StatusOK, response.Get(u))
}

//GetUserByEmail
func (m User) GetUserByEmail(c *gin.Context) {
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, response.Get(errors.ErrToJson(err)))
		return
	}
	ac := utils.GetAuthClient(c)
	u, err := ac.GetUserByEmail(context.Background(), m.Email)
	errors.CheckError(err)
	c.JSON(http.StatusOK, response.Get(u))
}

//GetAllUsers
func (m User) GetAllUsers(c *gin.Context) {

	ac := utils.GetAuthClient(c)

	// Iterating by pages 100 users at a time.
	// Note that using both the Next() function on an iterator and the NextPage()
	// on a Pager wrapping that same iterator will result in an error.
	pager := iterator.NewPager(ac.Users(context.Background(), ""), 100, "")
	var users []*auth.ExportedUserRecord
	for {
		nextPageToken, err := pager.NextPage(&users)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Get(fmt.Sprintf("paging error %v\n", err)))
		}
		if nextPageToken == "" {
			break
		}

	}
	c.JSON(http.StatusOK, response.Get(users))
}

//CreateUser
func (m User) CreateUser(c *gin.Context) {
	ac := utils.GetAuthClient(c)
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, response.Get(errors.ErrToJson(err)))
		return
	}
	params := m.setParams()
	u, err := ac.CreateUser(context.Background(), params)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Get(errors.ErrToJson(err)))
		//log.Fatalf("error creating user: %v\n", err)
		return
	}
	log.Printf("Successfully created user: %v\n", u)
	// [END create_user_golang]
	errors.CheckError(err)
	c.JSON(http.StatusCreated, response.Get(u))
}

//CreateUserWithUID
func (m User) CreateUserWithUID(c *gin.Context) {
	ac := utils.GetAuthClient(c)
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, response.Get(errors.ErrToJson(err)))
		return
	}
	params := m.setParams()
	u, err := ac.CreateUser(context.Background(), params)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Get(errors.ErrToJson(err)))
		//log.Fatalf("error creating user: %v\n", err)
		return
	}
	log.Printf("Successfully created user: %v\n", u)
	// [END create_user_golang]
	errors.CheckError(err)
	c.JSON(http.StatusOK, response.Get(u))
}

//UpdateUser
func (m User) UpdateUser(c *gin.Context) {
	ac := utils.GetAuthClient(c)
	var mu ModelUpdate
	if err := c.ShouldBindJSON(&mu); err != nil {
		c.JSON(http.StatusBadRequest, response.Get(errors.ErrToJson(err)))
		return
	}

	params := mu.setParams()
	u, err := ac.UpdateUser(context.Background(), mu.UID, params)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Get(errors.ErrToJson(err)))
		return
	}
	log.Printf("Successfully created user: %v\n", u)
	// [END create_user_golang]
	errors.CheckError(err)
	c.JSON(http.StatusOK, response.Get(u))
}

//DeleteUser
func (m User) DeleteUser(c *gin.Context) {
	ac := utils.GetAuthClient(c)
	uid := struct {
		Uid string `json:"uid" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&uid); err != nil {
		c.JSON(http.StatusBadRequest, response.Get(errors.ErrToJson(err)))
		return
	}
	err := ac.DeleteUser(context.Background(), uid.Uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Get(errors.ErrToJson(err)))
		return
	}
	c.JSON(http.StatusOK, nil)
}

//BulkDeleteUsers
func (m User) BulkDeleteUsers(c *gin.Context) {
	ac := utils.GetAuthClient(c)
	var uidS []struct {
		Uid string `json:"uid" binding:"required"`
	}

	if err := c.ShouldBindJSON(&uidS); err != nil {
		c.JSON(http.StatusBadRequest, response.Get(errors.ErrToJson(err)))
		return
	}
	var uids []string
	for _, v := range uidS {
		uids = append(uids, v.Uid)
	}
	deleteUsersResult, err := ac.DeleteUsers(context.Background(), uids)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Get(errors.ErrToJson(err)))
		return
	}
	c.JSON(http.StatusOK, response.Get(deleteUsersResult))
}

//CustomClaimsSet
func (m User) CustomClaimsSet(c *gin.Context) {
	ac := utils.GetAuthClient(c)
	claims := struct {
		Uid    string `json:"uid" binding:"required"`
		Claims []struct {
			Key   string      `json:"key" binding:"required"`
			Value interface{} `json:"value" binding:"required"`
		}
	}{}
	if err := c.ShouldBindJSON(&claims); err != nil {
		c.JSON(http.StatusBadRequest, response.Get(errors.ErrToJson(err)))
		return
	}
	cl := make(map[string]interface{})
	for _, v := range claims.Claims {
		cl[v.Key] = v.Value
	}
	err := ac.SetCustomUserClaims(context.Background(), claims.Uid, cl)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Get(errors.ErrToJson(err)))
		return
	}
	c.JSON(http.StatusOK, nil)
}

//setParams
func (m User) setParams() *auth.UserToCreate {
	u := (&auth.UserToCreate{}).
		Email(m.Email).
		EmailVerified(*m.EmailVerified).
		Password(m.Password)
	if m.UID != "" {
		u.UID(m.UID)
	}
	if m.PhoneNumber != "" {
		u.PhoneNumber(m.PhoneNumber)
	}
	if m.DisplayName != "" {
		u.DisplayName(m.DisplayName)
	}
	if m.PhotoURL != "" {
		u.PhotoURL(m.PhotoURL)
	}
	if m.Disabled != nil {
		u.Disabled(*m.Disabled)
	}
	return u
}

//setParams
func (m ModelUpdate) setParams() *auth.UserToUpdate {
	u := (&auth.UserToUpdate{}).
		Email(m.Email).
		EmailVerified(*m.EmailVerified).
		Password(m.Password)

	if m.PhoneNumber != "" {
		u.PhoneNumber(m.PhoneNumber)
	}
	if m.DisplayName != "" {
		u.DisplayName(m.DisplayName)
	}
	if m.PhotoURL != "" {
		u.PhotoURL(m.PhotoURL)
	}
	if m.Disabled != nil {
		u.Disabled(*m.Disabled)
	}
	return u
}
