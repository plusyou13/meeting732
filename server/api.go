package server

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/plusyou13/comm-go/log"
	"github.com/plusyou13/meeting732/store"
)

var weekDayS = []string{"日", "一", "二", "三", "四", "五", "六"}

func GetDate(ctx *gin.Context) {
	now := time.Now()

	result := []map[string]string{}

	for i := 0; i < 10; i++ {
		s := fmt.Sprintf("%d月%d日", now.Month(), now.Day())

		switch i {
		case 0:
			result = append(result, map[string]string{
				"今天": s,
			})
		case 1:
			result = append(result, map[string]string{
				"明天": s,
			})
		case 2:
			result = append(result, map[string]string{
				"后天": s,
			})
		default:
			result = append(result, map[string]string{
				"星期" + weekDayS[now.Weekday()]: s,
			})
		}

		now = now.AddDate(0, 0, 1)
	}

	ctx.JSON(200, result)
}

func PostMeeting(ctx *gin.Context) {
	var data store.MeetingData
	if ctx.ShouldBindJSON(&data) != nil {
		ctx.AbortWithStatus(500)
		return
	}

	err := store.SaveMeetingData(&data)
	if err != nil {
		log.Error(err)
		ctx.AbortWithStatus(501)
		return
	}
	ctx.JSON(200, struct{}{})
}

func DelMeetings(ctx *gin.Context) {
	date := ctx.Query("date")
	store.DeleteMeetingsWithDate(date)
}

func GetMeetings(ctx *gin.Context) {
	date := ctx.Query("date")
	results, err := store.GetMeetingsWithDate(date)
	if err != nil {
		log.Error(err)
		ctx.AbortWithStatus(500)
		return
	}

	if results == nil {
		results = []*store.MeetingData{}
	}
	ctx.JSON(200, results)
}
