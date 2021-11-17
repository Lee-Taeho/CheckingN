package database

import (
	"context"
	"server/middleware"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoDB) CreateNewStudent(student middleware.Student) error {
	ctx := context.TODO()
	if student.Appointments == nil {
		student.Appointments = make([]string, 0)
	}
	collection := m.mongo.Database(SJSU_DATABASE).Collection(STUDENTS_COLLECTION)
	result := collection.FindOne(ctx, bson.M{"email": student.Email})
	if result.Err() != nil {
		collection.InsertOne(ctx, student)
		return nil
	}
	return errors.New("User with this email already exists")
}

func (m *MongoDB) FindStudent(login middleware.LoginRequest) *middleware.Student {
	collection := m.mongo.Database(SJSU_DATABASE).Collection(STUDENTS_COLLECTION)
	ctx := context.TODO()
	student := &middleware.Student{}
	err := collection.FindOne(ctx, login).Decode(student)
	if err != nil {
		return nil
	}
	return student
}
