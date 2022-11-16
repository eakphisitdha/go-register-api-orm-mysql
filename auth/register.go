package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- Binding from JSON
type RegisterBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Avatar   string `json:"avartar"`
}

func Register(c *gin.Context) {

	//---Required alert
	var json RegisterBody
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//-- Check user exist
	var userExist User
	Db.Where("username = ?", json.Username).First(&userExist)
	if userExist.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"Status": "error", "message": "User Existed"})
		return
	}

	//---Create Record
	user := User{Username: json.Username, Password: json.Password, Fullname: json.Fullname, Avatar: json.Avatar}
	Db.Create(&user)

	if user.ID > 0 {
		c.JSON(http.StatusOK, gin.H{"Status": "ok", "message": "Create User Successful", "userId": user.ID})
	} else {
		c.JSON(http.StatusOK, gin.H{"Status": "error", "message": "Create User Failed"})
		return
	}
}
