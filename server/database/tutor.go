package database

import (
	"context"
	"fmt"
	"log"
	"server/middleware"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoDB) GetTutorsByCourseAndDate(course_code string, date time.Time) []middleware.Tutor {
	ctx := context.TODO()
	wkday := date.Weekday().String()
	idx := WEEKDAY_MAP[wkday]
	collection := m.mongo.Database(SJSU_DATABASE).Collection(TUTORS_COLLECTION)
	result, err := collection.Find(ctx, bson.M{"courses": course_code,
		fmt.Sprintf("availability.%d.0", idx): bson.M{"$exists": true}})
	if err != nil {
		log.Fatal(err)
	}
	var tutors []middleware.Tutor
	if err = result.All(ctx, &tutors); err != nil {
		log.Fatal(err)
	}
	for _, tutor := range tutors {
		available := tutor.Availability
		for _, appointment := range tutor.Appointments {
			index := -1
			app, _ := m.GetAppointment(appointment)
			if app != nil && app.StartTime.Year() == date.Year() &&
				app.StartTime.Month() == date.Month() &&
				app.StartTime.Day() == date.Day() {
				for i, hour := range available[idx] {
					if hour == app.StartTime.In(loc).Hour() {
						index = i
						break
					}
				}
				if index != -1 {
					available[idx] = append(available[idx][:index], available[idx][index+1:]...)
				}
			}
		}
		tutor.Availability = available
	}
	return tutors
}
