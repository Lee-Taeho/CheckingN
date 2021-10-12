package database

import (
	"context"
	"errors"
	"server/middleware"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (m *MongoDB) CreateNewStudent(student middleware.Student) error {
	ctx := context.TODO()
	collection := m.mongo.Database(USER_DATABASE).Collection(STUDENTS_COLLECTION)
	if m.FindUser(middleware.LoginRequest{Email: student.Email, Password: student.Password}, collection) {
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

func (m *MongoDB) FindStudent(login middleware.LoginRequest) bool {
	return m.FindUser(login, m.mongo.Database(USER_DATABASE).Collection(STUDENTS_COLLECTION))
}

func (m *MongoDB) FindUser(login middleware.LoginRequest, collection *mongo.Collection) bool {
	ctx := context.TODO()
	result := collection.FindOne(ctx, login)
	if result.Err() != nil {
		return false
	}
	return true
}
