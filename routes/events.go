package routes

import (
	"gin-quickstart/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)




func getEvents(c *gin.Context){

	events , err :=	models.GetAllEvents()
	  
	if(err !=nil){
		c.JSON(http.StatusInternalServerError , gin.H{
			"message": err.Error(),
		
		})
		return
	}
	  c.JSON(http.StatusOK, gin.H{ 
      "events": events,
    }) 
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