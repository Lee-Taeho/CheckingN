package database

import (
	"context"
	"errors"
	"server/middleware"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoDB) CreateNewStudent(student middleware.Student) error {
	ctx := context.TODO()
	collection := m.mongo.Database(USER_DATABASE).Collection(STUDENTS_COLLECTION)
	if m.FindStudent(middleware.LoginRequest{Email: student.Email, Password: student.Password}) != nil {
		return errors.New("Student already exists")
	}
	_, err := collection.InsertOne(ctx, student)
	return err
}

func (m *MongoDB) CreateNewGoogleStudent(student middleware.GoogleUser) {
	ctx := context.TODO()
	collection := m.mongo.Database(USER_DATABASE).Collection(GOOGLE_STUDENTS_COLLECTION)
	result := collection.FindOne(ctx, bson.M{"email": student.Email})
	if result.Err() != nil {
		collection.InsertOne(ctx, student)
	}
}

func (m *MongoDB) FindStudent(login middleware.LoginRequest) *middleware.Student {
	collection := m.mongo.Database(USER_DATABASE).Collection(STUDENTS_COLLECTION)
	ctx := context.TODO()
	student := &middleware.Student{}
	err := collection.FindOne(ctx, login).Decode(student)
	if err != nil {
		return nil
	}
	return student
}
