package database

import (
	"context"
	"errors"
	"log"
	"server/middleware"

	"go.mongodb.org/mongo-driver/bson"
)

func (m *MongoDB) CreateNewStudent(student middleware.Student) error {
	ctx := context.TODO()
	check := m.CheckForValidEmail(student)
	if check {
		collection := m.mongo.Database(USER_DATABASE).Collection(STUDENTS_COLLECTION)
		_, err := collection.InsertOne(ctx, student)
		return err
	}
	return errors.New("email already exist")
}

func (m *MongoDB) CheckForValidEmail(student1 middleware.Student) bool {
	ctx := context.TODO()
	collection := m.mongo.Database(USER_DATABASE).Collection(STUDENTS_COLLECTION)
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	var students []bson.M

	if err = cursor.All(ctx, &students); err != nil {
		log.Fatal(err)
	}

	for _, student2 := range students {
		if student2["email"] == student1.Email {
			return false
		}
	}

	return true
}
