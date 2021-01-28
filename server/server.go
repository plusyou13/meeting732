package server

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/plusyou13/comm-go/log"
)

func router() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(log.GinLog)
	g.Static("/static", "./static")
	g.LoadHTMLGlob("./static/views/*.html")

	g.GET("/", func(ctx *gin.Context) {
		date := ctx.Query("date")
		if date == "" {
			now := time.Now()
			date = fmt.Sprintf("%d月%d日", now.Month(), now.Day())
		}
		ctx.HTML(200, "index.html", map[string]string{"date": date})
	})
	g.GET("/create", func(ctx *gin.Context) {
		date := ctx.Query("date")
		ctx.HTML(200, "create.html", map[string]string{"date": date})
	})
	g.GET("/date", GetDate)
	g.POST("/meeting", PostMeeting)
	g.GET("/meeting", GetMeetings)
	return g
}

func RunService(port int) {
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router())
	if err != nil {
		fmt.Println(err, "setup secure api service fail")
		os.Exit(0)
	}
}
