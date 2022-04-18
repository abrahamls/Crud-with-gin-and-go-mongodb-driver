package controllers

import (
	"context"
	"errors"
	"example/web-service-gin/db"
	"example/web-service-gin/helpers"
	"example/web-service-gin/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsers(c *gin.Context) {

	var req models.UserCreateRequest
	err := c.Bind(&req)
	if err != nil {
		helpers.SendBadRequest(c, err)
		return
	}

	//check for email duplicate
	var user models.User
	filter := bson.M{"email": req.Email}
	err = db.UserColl.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			helpers.SendBadRequest(c, err)
		}
	} else {
		helpers.SendBadRequest(c, errors.New("email is already registered"))
		return
	}

	_, err = db.UserColl.InsertOne(context.TODO(), req)
	if err != nil {
		helpers.SendInternalServerError(c, err)
		return
	}
	helpers.SendSuccess(c)
}
