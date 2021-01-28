package store

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type MeetingData struct {
	Publisher    string    `json:"publisher" bson:"publisher"`
	Participants string    `json:"participants,omitempty" bson:"participants"`
	Content      string    `json:"content,omitempty" bson:"content"`
	Select       []int     `json:"select" bson:"select"`
	Date         string    `json:"date" bson:"date"`
	Year         int       `json:"year,omitempty" bson:"year"`
	CreateTime   time.Time `json:"create_time,omitempty" bson:"create_time"`
}

func SaveMeetingData(data *MeetingData) error {
	collection := mongoCli.Database(DB).Collection(C_Metting)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	data.Year = time.Now().Year()
	data.CreateTime = time.Now()
	_, err := collection.InsertOne(ctx, data)
	return err
}

func GetMeetingsWithDate(date string) ([]*MeetingData, error) {
	collection := mongoCli.Database(DB).Collection(C_Metting)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{"date": date, "year": time.Now().Year()})
	if err != nil {
		return nil, err
	}

	var results []*MeetingData
	for cur.Next(ctx) {
		var one MeetingData
		if cur.Decode(&one) == nil {
			results = append(results, &one)
		}
	}

	return results, nil
}
