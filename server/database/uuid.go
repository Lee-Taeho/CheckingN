package database

import (
	"context"

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
	collection := m.mongo.Database(USER_DATABASE).Collection(UUID_COLLECTION)

	uuid := &Uuid{}
	err := collection.FindOneAndUpdate(ctx, bson.M{}, bson.D{{"$inc", bson.D{{"uuid", 1}}}}, options.FindOneAndUpdate().SetUpsert(true)).Decode(uuid)
	if err != nil {
		return 0
	}

	return uuid.Uuid
}
