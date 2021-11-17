package database

import (
	"context"
	"errors"
	"log"
	"server/middleware"
	"server/zoom"
	"sort"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var WEEKDAY_MAP = map[string]int{
	"Monday":    0,
	"Tuesday":   1,
	"Wednesday": 2,
	"Thursday":  3,
	"Friday":    4,
}
var loc, _ = time.LoadLocation("America/Los_Angeles")

func (m *MongoDB) AddAppointment(appointment middleware.Appointment) error {
	//TODO: add course check
	ctx := context.TODO()
	app_collection := m.mongo.Database(SJSU_DATABASE).Collection(APPOINTMENTS_COLLECTION)
	tutor_collection := m.mongo.Database(SJSU_DATABASE).Collection(TUTORS_COLLECTION)
	filter := bson.M{"email": bson.M{"$eq": appointment.TutorEmail}}
	result := tutor_collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return errors.New("The tutor does not exist")
	}
	var tutor middleware.Tutor
	result.Decode(&tutor)
	//check course
	courses := tutor.Courses
	for i := 0; i < len(courses); i++ {
		if courses[i] == appointment.CourseCode {
			break
		}
		if i == len(courses)-1 && courses[i] != appointment.CourseCode {
			return errors.New("Requested tutor is not qualified for this course")
		}
	}
	//check for time conflict
	if m.timeConflict(tutor, appointment) {
		return errors.New("Tutor is not available for requested time slot")
	}

	//create join and start zoom link
	if strings.Compare(appointment.MeetingLocation, "Zoom") == 0 {
		joinLink, startLink, _ := zoom.CreateZoomLink(appointment)
		appointment.JoinLink = joinLink
		appointment.StartLink = startLink
	}

	//add appointment to db
	returnedApp, _ := app_collection.InsertOne(ctx, appointment)
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
	//add app to student object
	filter = bson.M{"email": bson.M{"$eq": appointment.StudentEmail}}

	student_collection := m.mongo.Database(SJSU_DATABASE).Collection(STUDENTS_COLLECTION)

	update = bson.M{"$addToSet": bson.M{"appointments": app_id}}
	if _, err := student_collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.New("Could not update student")
	}
	return nil
}

func (m *MongoDB) GetAppointment(id string) (*middleware.Appointment, error) {
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

func (m *MongoDB) DeleteAppointment(id string) error {
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
	log.Println(appointment)

	if err = m.DeleteAppointmentFromTutor(id, appointment.TutorEmail); err != nil {
		return errors.New("Could not delete appointment for tutor")
	}
	if err = m.DeleteAppointmentFromStudent(id, appointment.StudentEmail); err != nil {
		return errors.New("Could not delete appointment for student")
	}
	return nil
}

func (m *MongoDB) DeleteAppointmentFromTutor(app_id string, tutor_email string) error {
	ctx := context.TODO()
	collection := m.mongo.Database(SJSU_DATABASE).Collection(TUTORS_COLLECTION)
	filter := bson.M{"email": bson.M{"$eq": tutor_email}}
	update := bson.M{"$pull": bson.M{"appointments": app_id}}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.New("Could not delete appointment from tutor's schedule")
	}
	return nil
}

func (m *MongoDB) DeleteAppointmentFromStudent(app_id string, student_email string) error {
	ctx := context.TODO()
	collection := m.mongo.Database(SJSU_DATABASE).Collection(STUDENTS_COLLECTION)
	filter := bson.M{"email": bson.M{"$eq": student_email}}
	update := bson.M{"$pull": bson.M{"appointments": app_id}}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		return errors.New("Could not delete appointment from student's schedule")
	}
	return nil
}

func (m *MongoDB) GetAppointmentsForStudent(student_email string) ([]string, error) {
	ctx := context.TODO()
	collection := m.mongo.Database(SJSU_DATABASE).Collection(STUDENTS_COLLECTION)
	filter := bson.M{"email": bson.M{"$eq": student_email}}
	result := collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return nil, errors.New("The student does not exist")
	}
	var student middleware.Student
	result.Decode(&student)
	appointments := student.Appointments
	var results []string
	for i := 0; i < len(appointments); i++ {
		log.Printf("loop")
		results = append(results, appointments[i])

	}
	return results, nil
}

func (m *MongoDB) timeConflict(tutor middleware.Tutor, appointment middleware.Appointment) bool {
	//check if no conflict with tutor's availability
	availability := tutor.Availability
	wkday := WEEKDAY_MAP[appointment.StartTime.Weekday().String()]
	idx := sort.SearchInts(availability[wkday], appointment.StartTime.Hour())
	if len(availability[wkday]) <= idx || availability[wkday][idx] != appointment.StartTime.Hour() {
		return true
	}
	//go through each appointment and check if timeslot booked
	appointments := tutor.Appointments
	for i := 0; i < len(appointments); i++ {
		if app, err := m.GetAppointment(appointments[i]); err == nil {
			if app.StartTime.Equal(appointment.StartTime) {
				return true
			}
		}
	}
	return false
}

/*
func (m *MongoDB) UpdateAppointment(app_id string, newAppointment middleware.Appointment) error {
	ctx := context.TODO()
	app_collection := m.mongo.Database(SJSU_DATABASE).Collection(APPOINTMENTS_COLLECTION)
	objID, err := primitive.ObjectIDFromHex(app_id)
	if err != nil {
		return errors.New("Invalid value for appointment ID")
	}
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	result := app_collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return errors.New("Appointment with such id does not exist")
	}
	var appointment middleware.Appointment
	result.Decode(&appointment)
	tutorObjID, err := primitive.ObjectIDFromHex(appointment.TutorID)
	filter = bson.M{"_id": bson.M{"$eq": tutorObjID}}
	tutor_collection := m.mongo.Database(SJSU_DATABASE).Collection(TUTORS_COLLECTION)
	var tutor middleware.Tutor
	result = tutor_collection.FindOne(ctx, filter)
	result.Decode(&tutor)
	update := bson.M{"$set": bson.M{"meeting_location": newAppointment.MeetingLocation}}
	if newAppointment.MeetingLocation == "" {
		update = nil
	}
	if !newAppointment.StartTime.IsZero() {
		//check for time conflict
		if !m.timeConflict(tutor, newAppointment) {
			update = bson.M{"$set": bson.M{"start_time": newAppointment.StartTime,
				"end_time": newAppointment.StartTime.Add(time.Hour * 1)}}
		} else {
			return errors.New("Tutor is not available for this time slot")
		}
	}
	if update == nil {
		return errors.New("No value to udpate appointment")
	}
	if _, err := app_collection.UpdateByID(ctx, objID, update); err != nil {
		return errors.New("Could not update appointment")
	}
	return nil
}


func (m *MongoDB) FindTutorId(tutor middleware.Tutor) int {

	ctx := context.TODO()
	tutor_collection := m.mongo.Database(SJSU_DATABASE).Collection(TUTORS_COLLECTION)
	filter := bson.M{"_id": bson.M{"$eq": tutor.Email}}
	result := tutor_collection.FindOne(ctx, filter)
	var temp middleware.Tutor
	result.Decode(temp)
	return temp.Uuid
}
*/

/*
func (m *MongoDB) GetAppointmentsForTutor(tutor_id string) ([]middleware.Appointment, error) {
	ctx := context.TODO()
	collection := m.mongo.Database(SJSU_DATABASE).Collection(TUTORS_COLLECTION)
	tutorObjID, err := primitive.ObjectIDFromHex(tutor_id)
	if err != nil {
		return nil, errors.New("Invalid value for tutor ID")
	}
	filter := bson.M{"_id": bson.M{"$eq": tutorObjID}}
	result := collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return nil, errors.New("The tutor does not exist")
	}
	var tutor middleware.Tutor
	result.Decode(&tutor)
	appointments := tutor.Appointments
	var results []middleware.Appointment
	for i := 0; i < len(appointments); i++ {
		if app, err := m.GetAppointment(appointments[i]); err == nil {
			results = append(results, *app)
		}
	}
	return results, nil
}
*/
