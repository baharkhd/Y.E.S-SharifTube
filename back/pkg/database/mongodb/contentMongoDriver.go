package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"yes-sharifTube/internal/model/content"
)

type ContentMongoDriver struct {
	collection *mongo.Collection
}

func (c *ContentMongoDriver) Get(contentID primitive.ObjectID) (*content.Content, error) {
	panic("not implemented")
}
func (c *ContentMongoDriver) GetAll(courseID primitive.ObjectID, tags []string, start, amount int) ([]*content.Content, error) {
	panic("not implemented")
}
func (c *ContentMongoDriver) Insert(username string, courseID primitive.ObjectID, content *content.Content) (*content.Content, error) {
	panic("not implemented")
}
func (c *ContentMongoDriver) Update(username string, courseID primitive.ObjectID, content *content.Content) (*content.Content, error) {
	panic("not implemented")
}
func (c *ContentMongoDriver) Delete(username string, courseID, contentID primitive.ObjectID) (*content.Content, error) {
	panic("not implemented")
}

func NewContentMongoDriver(db, collection string) *ContentMongoDriver {
	return &ContentMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
