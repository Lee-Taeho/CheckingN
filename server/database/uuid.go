package database

import (
	"context"
	"fmt"
	"server/middleware"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Uuid struct {
	Uuid int `bson:"uuid"`
}

func (m *MongoDB) GetUUID() int {
	m.uuidLock.Lock()
	defer m.uuidLock.Unlock()

	ctx := context.TODO()
	collection := m.mongo.Database(SJSU_DATABASE).Collection(UUID_COLLECTION)

	uuid := &Uuid{}
	err := collection.FindOneAndUpdate(ctx, bson.M{}, bson.D{{"$inc", bson.D{{"uuid", 1}}}}, options.FindOneAndUpdate().SetUpsert(true)).Decode(uuid)
	if err != nil {
		return 0
	}

	return uuid.Uuid
}

func (m *MongoDB) FindStudentUUID(uuid int) *middleware.Student {
	ctx := context.TODO()
	collection := m.mongo.Database(SJSU_DATABASE).Collection(STUDENTS_COLLECTION)
	filter := bson.M{"uuid": uuid}

	student := &middleware.Student{}
	err := collection.FindOne(ctx, filter).Decode(student)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	student.Uuid = 0
	student.Password = ""

	return student
}
