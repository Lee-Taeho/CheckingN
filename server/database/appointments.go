package database

import (
	"context"
	"server/middleware"
	"log"
	"errors"
	"sort"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


var WEEKDAY_MAP = map[string]int{
	"Monday": 0,
	"Tuesday": 1,
	"Wednesday": 2,
	"Thursday": 3,
	"Friday": 4,
}
var loc, _ = time.LoadLocation("America/Los_Angeles")

func (m *MongoDB) AddAppointment (appointment middleware.Appointment) error {
	ctx := context.TODO()
	app_collection := m.mongo.Database(SJSU_DATABASE).Collection(APPOINTMENTS_COLLECTION)
	tutor_collection := m.mongo.Database(SJSU_DATABASE).Collection(TUTORS_COLLECTION)
	tutorObjID, err := primitive.ObjectIDFromHex(appointment.TutorID)
	filter := bson.M{"_id": bson.M{"$eq": tutorObjID}}
	result := tutor_collection.FindOne(ctx, filter) 
	if result.Err() != nil {
		return errors.New("The tutor does not exist")
	}
	var tutor middleware.Tutor
	result.Decode(&tutor)
	//check if no conflict with tutor's availability
	availability := tutor.Availability
	wkday := WEEKDAY_MAP[appointment.StartTime.Weekday().String()]
	idx := sort.SearchInts(availability[wkday], appointment.StartTime.Hour())
	if len(availability[wkday]) <= idx || availability[wkday][idx] != appointment.StartTime.Hour() {
		return errors.New("The tutor is not available for this time slot")
	}
	//go through each appointment and check if timeslot booked
	appointments := tutor.Appointments
	for i := 0; i < len(appointments); i++ {
		if app, err := m.GetAppointment(appointments[i]); err == nil {
			if app.StartTime.Equal(appointment.StartTime) {
				return errors.New("The requested timeslot is already booked")
			}
		}
	}
	//add appointment to db
	returnedApp, err := app_collection.InsertOne(ctx, appointment)
	app_id := ""
	if oid, ok := returnedApp.InsertedID.(primitive.ObjectID); ok {
		app_id = oid.Hex()
	} else {		
		return errors.New("Could not save appointment")
	}
	//add app_id to tutor object
	update := bson.M{"$addToSet": bson.M{"appointments": app_id}}
	if _, err := tutor_collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.New("Could not update tutor")
	}
	log.Println(err)
	//need to also add app to student object
	studentObjID, err := primitive.ObjectIDFromHex(appointment.StudentID)
	filter = bson.M{"_id": bson.M{"$eq": studentObjID}}
	student_collection := m.mongo.Database(SJSU_DATABASE).Collection(STUDENTS_COLLECTION)
	update = bson.M{"$addToSet": bson.M{"appointments": app_id}}
	if _, err := student_collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.New("Could not update student")
	}
	return nil
}

func (m *MongoDB) GetAppointment (id string) (*middleware.Appointment, error) {
	ctx := context.TODO()
	app_collection := m.mongo.Database(SJSU_DATABASE).Collection(APPOINTMENTS_COLLECTION)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("Invalid data")
	}
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	result := app_collection.FindOne(ctx, filter) 
	if result.Err() != nil {
		return nil, errors.New("No appointment with this id")
	}
	var appointment middleware.Appointment
	result.Decode(&appointment)
	return &appointment, nil
}

func (m *MongoDB) DeleteAppointment (id string) error {
	ctx := context.TODO()
	app_collection := m.mongo.Database(SJSU_DATABASE).Collection(APPOINTMENTS_COLLECTION)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("Invalid value")
	}
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	result := app_collection.FindOneAndDelete(ctx, filter)
	if result.Err() != nil {
		return errors.New("This appointment does not exist")
	}
	var appointment middleware.Appointment
	result.Decode(&appointment)
	tutorObjID, err := primitive.ObjectIDFromHex(appointment.TutorID)
	studentObjID, err := primitive.ObjectIDFromHex(appointment.StudentID)
	if err = m.DeleteAppointmentFromTutor(objID, tutorObjID); err != nil {
		if err = m.DeleteAppointmentFromStudent(objID, studentObjID); err != nil {
			return nil
		}
	}
	return errors.New("Could not update student and/or tutor")
}

func (m *MongoDB) DeleteAppointmentFromTutor (app_id primitive.ObjectID, tutor_id primitive.ObjectID) error {
	ctx := context.TODO()
	collection := m.mongo.Database(SJSU_DATABASE).Collection(TUTORS_COLLECTION)
	filter := bson.M{"_id": bson.M{"$eq": tutor_id}}
	update := bson.M{"$pull": bson.M{"appointments": app_id}}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.New("Could not delete appointment from tutor's schedule")
	}
	return nil
}

func (m *MongoDB) DeleteAppointmentFromStudent (app_id primitive.ObjectID, student_id primitive.ObjectID) error {
	ctx := context.TODO()
	collection := m.mongo.Database(SJSU_DATABASE).Collection(STUDENTS_COLLECTION)
	filter := bson.M{"_id": bson.M{"$eq": student_id}}
	update := bson.M{"$pull": bson.M{"appointments": app_id}}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.New("Could not delete appointment from student's schedule")
	}
	return nil
}

func (m *MongoDB) GetAppointmentsForTutor (tutor_id string) ([]*middleware.Appointment, error) {
	ctx := context.TODO()
	collection := m.mongo.Database(SJSU_DATABASE).Collection(TUTORS_COLLECTION)
	tutorObjID, err := primitive.ObjectIDFromHex(tutor_id)
	if err != nil {
		return nil, errors.New("Invalid data")
	}
	filter := bson.M{"_id": bson.M{"$eq": tutorObjID}}
	result := collection.FindOne(ctx, filter) 
	if result.Err() != nil {
		return nil, errors.New("The tutor does not exist")
	}
	var tutor middleware.Tutor
	result.Decode(&tutor)
	appointments := tutor.Appointments
	var results []*middleware.Appointment
	for i := 0; i < len(appointments); i++ {
		if app, err := m.GetAppointment(appointments[i]); err == nil {
			results = append(results, app)
		}
	}
	return results, nil
}
