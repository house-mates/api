package main

import (
	"github.com/gin-gonic/gin"

	"github.com/house-mates/api/controllers"
	"github.com/house-mates/api/models"
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

	r.POST("/actions", controllers.AddAction)
	r.GET("/actions", controllers.FindActions)
	r.GET("/actions/:id", controllers.FindAction)
	r.GET("/actions_by_uid/:user_id", controllers.FindActionByUserID)
	r.PATCH("/actions/:id", controllers.EditAction)
	r.DELETE("/actions/:id", controllers.DeleteAction)
	r.DELETE("/actions_by_uid/:user_id", controllers.DeleteActionsByUserID)

	r.POST("/events", controllers.AddEvent)
	r.GET("/events", controllers.FindEvents)
	r.GET("/events/:id", controllers.FindEvent)
	r.GET("/events_by_aid/:action_id", controllers.FindEventsByActionID)
	r.GET("/events_by_range", controllers.FindEventsByTimeRange)
	r.GET("/events_by_range_and_aid/:action_id", controllers.FindEventsByTimeRangeAndActionID)
	r.PATCH("/events/:id", controllers.EditEvent)
	r.DELETE("/events/:id", controllers.DeleteEvent)
	r.DELETE("/events_by_aid/:action_id", controllers.DeleteEventsByActionID)

	r.POST("/responses", controllers.AddResponse)
	r.GET("/responses", controllers.FindResponses)
	r.GET("/responses/:id", controllers.FindResponse)
	r.GET("/responses_by_eid/:event_id", controllers.FindResponsesByEventID)
	r.GET("/responses_by_uid/:user_id", controllers.FindResponsesByUserID)
	r.DELETE("/responses/:id", controllers.DeleteResponse)
	r.DELETE("/responses_by_eid/:event_id", controllers.DeleteResponsesByEventID)
	r.DELETE("/responses_by_uid/:user_id", controllers.DeleteResponsesByUserID)

	r.Run()
}
