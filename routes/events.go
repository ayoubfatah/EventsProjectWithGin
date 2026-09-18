package routes

import (
	"fmt"
	"gin-quickstart/models"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)



func getEvents(c *gin.Context) {

	page := 1
	limit := 4



	if pageQuery := c.Query("page"); pageQuery != "" {

		parsedPage, err := strconv.Atoi(pageQuery)

		if err != nil || parsedPage < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "page must be a positive integer",
			})
			return
		}

		page = parsedPage
	}



	if limitQuery := c.Query("limit"); limitQuery != "" {

		parsedLimit, err := strconv.Atoi(limitQuery)

		if err != nil || parsedLimit < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "limit must be a positive integer",
			})
			return
		}

		limit = parsedLimit
	}

	// Prevent clients from requesting huge pages.
	if limit > 50 {
		limit = 50
	}


	result, err := models.GetAllEvents(page, limit)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't fetch events",
		})
		return
	}



	c.JSON(http.StatusOK, result)
}

func getEvent(c *gin.Context){
	id  , err :=  strconv.ParseInt(c.Param("id"),10 , 64)
	if(err !=nil){
		c.JSON(http.StatusBadRequest , gin.H{
			"message": "id must be an int",		
		})
		return
	}
	event , err := models.GetEventById(id)
		if(err !=nil){
		c.JSON(http.StatusInternalServerError , gin.H{
			"message": err.Error(),		
		})
		return
	}
	  c.JSON(http.StatusOK, gin.H{ 
      "event": event,
    }) 

}

func getEventsByCity(c *gin.Context){
	city := c.Param("city")

	
	events , err := models.GetEventsByCity(city)
	if(err!= nil){
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}


func getEventBySlug(c *gin.Context){
	slug   :=  c.Param("slug")

	event , err := models.GetEventBySlug(slug)
		if(err !=nil){
		c.JSON(http.StatusInternalServerError , gin.H{
			"message": err.Error(),		
		})
		return
	}
	  c.JSON(http.StatusOK, gin.H{ 
      "event": event,
    }) 

}

func createEvents( c *gin.Context){
	
  var event models.Event

   if err  :=	c.ShouldBindJSON(&event); err !=nil{

	 c.JSON(http.StatusBadRequest, gin.H{
            "error": "something happend while trying to bind the json",
        })


	return
   }

  userId := c.GetInt64("userId")
  event.UserId= userId
  err := event.Save()

  	if(err !=nil){
		c.JSON(http.StatusInternalServerError , gin.H{
		"message":err.Error(),
		})
		return 
	}
  
  c.JSON(http.StatusCreated, gin.H{
	"message":"Event Created",
	"event": event,
  })

}
func seedEvents(c *gin.Context) {
	code := c.Param("code")

	expectedCode := os.Getenv("SEED_CODE")

	if expectedCode == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Seed code is not configured",
		})
		return
	}

	if code != expectedCode {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Invalid seed code",
		})
		return
	}

	err := models.SeedEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Events created successfully",
	})
}



func GetCurrentUserRegisteredEvents(c *gin.Context){
  	userId := c.GetInt64("userId")
	registeredEvents , err :=  models.GetRegisteredEventsByUser(userId)
	
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			 gin.H {"message":""})
	return
	}
	c.JSON(http.StatusOK , gin.H{
		"events": registeredEvents,
	})

}
 
func updateEvent(c *gin.Context){
	id , err :=  strconv.ParseInt(c.Param("id"),10 , 64)
	if(err !=nil){
		c.JSON(http.StatusBadRequest , gin.H{
			"message": "id must be an int",		
		})
		return
	}
	userId := c.GetInt64("userId")
	event , err := models.GetEventById(id)


	
	
	
	if err !=nil{
		c.JSON(http.StatusInternalServerError , gin.H{"message": err.Error()})
	return}	
		
		
		if event.UserId != userId {
			c.JSON(http.StatusUnauthorized , gin.H{"message": "this event doesnt belong to you ",})
			return
		}	
		
		
	var updatedEvent models.Event


	if err  :=	c.ShouldBindJSON(&updatedEvent); err !=nil{
	 c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
    return
   }
  	updatedEvent.Id = id 
	err = updatedEvent.Update()

  	if(err !=nil){
		c.JSON(http.StatusInternalServerError , gin.H{
		"message":err.Error(),
		})
		return 
	}
  
  	c.JSON(http.StatusOK, gin.H{
	"message":"Event Updated successfully",
	"event":updatedEvent,
  	})	
	
}

func  deleteEvent(c *gin.Context){
	id , err := strconv.ParseInt(c.Param("id"),10 ,64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id must be an int",
		})
		return
	}
    event ,err := models.GetEventById(id)
	userId := c.GetInt64("userId")

		
	if(err !=nil){
	c.JSON(http.StatusNotFound , gin.H{"message": "Couldn't delete the event ",})
	return}	
	



	
	if event.UserId != userId {
		c.JSON(http.StatusUnauthorized , gin.H{"message": "this event doesnt belong to you ",})
		return
	}	


	err = event.Delete()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":"couldn't delete the event ",
		})
		return
	}
		c.JSON(http.StatusOK, gin.H{
			"message":"Event deleted successfully",
		})
}


func registerEvent(c *gin.Context){

	userId := c.GetInt64("userId")
	eventId , err := strconv.ParseInt(c.Param("id"),10 ,64)
	fmt.Println("event iD")
	fmt.Println(eventId)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id must be an int",
		})
		return
	}
	event , err :=  	models.GetEventById(eventId)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't fetch the event ",
		})
		return
	}

   err = event.Register(userId)

   	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't  register event ",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
			"message": "registered event successfully",
	})

}



func cancelRegistration(c *gin.Context){

	userId := c.GetInt64("userId")
	eventId , err := strconv.ParseInt(c.Param("id"),10 ,64)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id must be an int",
		})
		return
	}
	

	var event models.Event
	event.Id = eventId

   err = event.CancelRegistration(userId )

   	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't  cancel registration  ",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
			"message": "Canceled event registration successfully",
	})

}




func getCurrentUserEvents(c *gin.Context){
	userId := c.GetInt64("userId")
	events , err :=	models.GetEventsByUserID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't fetch the user's events",
		})
		return
	}


c.JSON(http.StatusOK , gin.H{
	"events" :events,
})

}	


func getEventReservation(c *gin.Context) {
	fmt.Println("firing")
	userId := c.GetInt64("userId")

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id must be an int",
		})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
		})
		return
	}

	reserved, err := event.IsRegistered(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Couldn't check event reservation",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"reserved": reserved,
	})
}