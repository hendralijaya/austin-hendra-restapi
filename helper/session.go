package helper

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

const userKey = "session_id"

// Use cookie for session id
func SetSession() {
	r := gin.Default()
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(userKey))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	r.Run(":8000")
}


// Save user session
func SaveSession(c *gin.Context, userID int) {
	session := sessions.Default(c)
	session.Set(userKey, userID)
	session.Save()
}

// Clear user session
func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
