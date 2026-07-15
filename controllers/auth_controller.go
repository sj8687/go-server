package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sj8687/backend/config"
	"github.com/sj8687/backend/models"
	"github.com/sj8687/backend/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	var body LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if body.Username == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username and password are required",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User

	err := config.UserCollection.FindOne(
		ctx,
		bson.M{"username": body.Username},
	).Decode(&user)

	// -------------------------
	// USER DOESN'T EXIST
	// -------------------------
	if err != nil {

		hashedPassword, err := utils.HashPassword(body.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to hash password",
			})
			return
		}

		newUser := models.User{
			Username: body.Username,
			Password: hashedPassword,
		}

		
		_, err = config.UserCollection.InsertOne(ctx, newUser)


		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create user",
			})
			return
		}

		token, err := utils.GenerateJWT(body.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate token",
			})
			return
		}

		c.SetCookie(
			"token",
			token,
			86400,
			"/",
			"",
			false,
			true,
		)

		c.JSON(http.StatusCreated, gin.H{
			"message": "User created successfully",
		})

		return
	}

	// -------------------------
	// USER EXISTS
	// -------------------------
	if !utils.CheckPassword(user.Password, body.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid password",
		})
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	c.SetCookie(
		"token",
		token,
		86400,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}