package controllers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/osvaldodmvs/api/initializers"
	"github.com/osvaldodmvs/api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SignUpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.tmpl", gin.H{})
}

func SignUp(c *gin.Context) {
	var userData models.User

	if err := c.Bind(&userData); err != nil {
		//if the json is not valid, return a bad request
		log.Println("Error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error binding JSON"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 10)

	if err != nil {
		//if there is an error, print it and return a bad request
		log.Println("Error hashing password: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error hashing password"})
		return
	}

	user := models.User{Email: userData.Email, Password: string(hashedPassword)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		//if there is an error, print it and return a bad request
		log.Println("Error creating user: ", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

func Login(c *gin.Context) {
	var userData models.User

	if err := c.Bind(&userData); err != nil {
		//if the json is not valid, return a bad request
		log.Println("Error binding JSON: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error binding JSON"})
		return
	}

	var user models.User

	result := initializers.DB.First(&user, "email = ?", userData.Email)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			//id not found
			c.JSON(http.StatusNotFound, gin.H{"message": "Resource not found"})
		} else {
			//400 for other errors
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		}
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))

	if err != nil {
		log.Println("Invalid e-mail or password: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid e-mail or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject":    user.ID,
		"email":      user.Email,
		"expiration": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		log.Println("Error creating token: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error creating token"})
		return
	}
	//no https so no SameSiteNoneMode
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("jwt", tokenString, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}

/*testing only

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(200, gin.H{
		"message": user,
	})
}

*/
