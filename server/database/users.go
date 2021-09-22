package database

import (
	"context"
	"server/middleware"
)

func (m *MongoDB) CreateNewStudent(student middleware.Student) error {
	ctx := context.TODO()
	collection := m.mongo.Database(USER_DATABASE).Collection(STUDENTS_COLLECTION)
	_, err := collection.InsertOne(ctx, student)
	return err
}
