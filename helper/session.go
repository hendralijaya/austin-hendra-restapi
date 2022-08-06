package helper

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// Use redis for session id
func SetSession() gin.HandlerFunc {
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	return sessions.Sessions("mysession", store)
}

// Save session to redis
func SaveSession(c *gin.Context, userID int) {
	r := gin.Default()
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
	})
}

// Clear user session
func ClearSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
