package main

import (
	"github.com/gin-gonic/gin"

	"github.com/house-mates/api/src/controllers"
	"github.com/house-mates/api/src/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.POST("/users", controllers.AddUser)
	r.GET("/users", controllers.FindUsers)
	r.GET("/users/:id", controllers.FindUser)
	r.PATCH("/users/:id", controllers.EditUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.POST("/usermeta", controllers.AddUsermeta)
	r.GET("/usermeta", controllers.FindUsermeta)
	r.GET("/usermeta/:id", controllers.FindUsermetaByID)
	r.GET("/usermeta_by_uid/:user_id", controllers.FindUsermetaByUserID)
	r.PATCH("/usermeta/:id", controllers.EditUsermeta)
	r.DELETE("/usermeta/:id", controllers.DeleteUsermeta)
	r.DELETE("/usermeta_by_uid/:user_id", controllers.DeleteUsermetaByUserID)

	r.Run()
}
