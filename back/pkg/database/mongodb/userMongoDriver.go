package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"yes-sharifTube/internal/model/user"
	"yes-sharifTube/pkg/database/status"
)

type UserMongoDriver struct {
	collection *mongo.Collection
}


/*	mongoDB implementation of the UserDBDriver interface
	here we have a UserMongoDriver which can be initialized with a collaborating collection
	to perform the CRUD for user.User model on mongo
*/
func (u UserMongoDriver) GetAll(start,amount int64) ([]*user.User, status.QueryStatus) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	var result []*user.User
	if cur, err := u.collection.Find(ctx, bson.D{}); err != nil {
		return nil, status.FAILED
	} else {
		defer cur.Close(ctx)
		for cur.Next(context.Background()) {
			if start>0{
				start--
				continue
			}
			if amount ==0{
				break
			}
			amount--
			var blogUser user.User
			_ = cur.Decode(&blogUser)
			result = append(result, &blogUser)
		}
		return result,status.SUCCESSFUL
	}
}

func (u UserMongoDriver) Insert(user *user.User) status.QueryStatus {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	user.ID=primitive.NewObjectID()
	if _, err := u.collection.InsertOne(ctx, user); err != nil {
		return status.FAILED
	}
	return status.SUCCESSFUL

}

func (u UserMongoDriver) Get(username *string) (*user.User, status.QueryStatus) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	var result user.User
	if err := u.collection.FindOne(ctx, bson.M{"username":username}).Decode(&result); err != nil {
		return &result, status.FAILED
	}
	return &result, status.SUCCESSFUL
}

func (u UserMongoDriver) Delete(username *string) status.QueryStatus {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	if _, err := u.collection.DeleteOne(ctx, bson.M{"username":username}); err != nil {
		return status.FAILED
	}
	return status.SUCCESSFUL
}

func (u UserMongoDriver) Update(target string, user *user.User) status.QueryStatus {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	query:=bson.M{}
	update:=bson.M{"$set":query}
	if user.Name != ""{
		query["name"]=user.Name
	}
	if user.Username != ""{
		query["username"]=user.Username
	}
	if user.Email != ""{
		query["email"]=user.Email
	}
	if user.Password != ""{
		query["password"]=user.Password
	}
	if updateResult, err := u.collection.UpdateOne(ctx,bson.M{"name":target}, update);
	err != nil || updateResult.MatchedCount==0 {
		return status.FAILED
	}
	return status.SUCCESSFUL

}


/*	here in this new function we take a dbname and collection name and retrieve
	the corresponding collection for our UserMongoDriver instance to work with
*/
func NewUserMongoDriver(db, collection string) *UserMongoDriver {
	return &UserMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}

func (u UserMongoDriver) Replace(target *string, toBe *user.User) status.QueryStatus {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	if _, err := u.collection.ReplaceOne(ctx,bson.M{"username":target},toBe); err != nil {
		return status.FAILED
	}
	return status.SUCCESSFUL

}
