package database

import (
	"context"
	"errors"
	"server/middleware"
)

func (m *MongoDB) CreateNewStudent(student middleware.Student) error {
	ctx := context.TODO()
	collection := m.mongo.Database(USER_DATABASE).Collection(STUDENTS_COLLECTION)
	if m.FindStudent(middleware.LoginRequest{Email: student.Email, Password: student.Password}) {
		return errors.New("Student already exists")
	}
	_, err := collection.InsertOne(ctx, student)
	return err
}

func (m *MongoDB) FindStudent(login middleware.LoginRequest) bool {
	ctx := context.TODO()
	collection := m.mongo.Database(USER_DATABASE).Collection(STUDENTS_COLLECTION)
	result := collection.FindOne(ctx, login)
	if result.Err() != nil {
		return false
	}
	return true
}
