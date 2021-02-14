package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"strings"
	"sync"
	"time"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/internal/model/user"
	"yes-sharifTube/pkg/database/status"
)

type CourseMongoDriver struct {
	collection *mongo.Collection
}

var doOnce sync.Once

func (c CourseMongoDriver) Get(courseID primitive.ObjectID) (*course.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id": courseID,
	}
	var res course.Course
	if err := c.collection.FindOne(ctx, target).Decode(&res); err != nil {
		return nil, model.CourseNotFoundException{Message: "course couldn't found."}
	}
	return &res, nil
}

func (c CourseMongoDriver) GetByFilter(keywords []string, start, amount int) ([]*course.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	// indexing for searching by keywords
	var err error
	doOnce.Do(func() {
		_, err = c.collection.Indexes().CreateMany(ctx, []mongo.IndexModel{
			{
				Keys: bsonx.Doc{
					{Key: "title", Value: bsonx.String("text")},
					{Key: "summery", Value: bsonx.String("text")},
					{Key: "contents.title", Value: bsonx.String("text")},
					{Key: "contents.description", Value: bsonx.String("text")}},
			},
		})
	})
	if err != nil {
		return nil, model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}

	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})
	opts.SetLimit(int64(amount))
	opts.SetSkip(int64(start))
	var target bson.M
	if len(keywords) == 0 {
		target = bson.M{}
	} else {
		target = bson.M{
			"$text": bson.M{
				"$search": strings.Join(keywords, " "),
			},
		}
	}
	courr, err := c.collection.Find(ctx, target, opts)
	if err != nil {
		return nil, model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	defer courr.Close(ctx)
	var cours []*course.Course
	i := 0
	for courr.Next(context.Background()) {
		var ctmp course.Course
		_ = courr.Decode(&ctmp)
		cours = append(cours, &ctmp)
		i++
	}
	return cours, nil
}

// todo fix inconsistency
func (c CourseMongoDriver) Insert(course *course.Course) (*course.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	course.ID = primitive.NewObjectID()
	if _, err := c.collection.InsertOne(ctx, course); err != nil {
		return nil, model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	if stat := user.DBD.Enroll(course.ProfUn, course.ID.Hex()); stat == status.FAILED {
		return nil, model.InternalServerException{Message: "database couldn't add user"}
	}
	return course, nil
}

func (c CourseMongoDriver) UpdateInfo(courseID primitive.ObjectID, title, summery, token string) error {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id": courseID,
	}
	change := bson.M{
		"$set": bson.M{
			"title":   title,
			"summery": summery,
			"token":   token,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	return nil
}

// todo fix inconsistency
func (c CourseMongoDriver) Delete(courseID primitive.ObjectID, members []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id": courseID,
	}
	if _, err := c.collection.DeleteOne(ctx, target); err != nil {
		return model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	for i := range members {
		//Ignoring the non existed users
		user.DBD.Leave(members[i], courseID.Hex())
	}
	return nil
}

func (c CourseMongoDriver) UpdateStdList(course *course.Course) error {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id": course.ID,
	}
	change := bson.M{
		"$set": bson.M{
			"students": course.StdUns,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	return nil
}

func (c CourseMongoDriver) UpdateTaList(course *course.Course) error {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id": course.ID,
	}
	change := bson.M{
		"$set": bson.M{
			"tas": course.TaUns,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	return nil
}

//todo fix Inconsistency
func (c CourseMongoDriver) AddStd(username string, courseID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id": courseID,
	}
	change := bson.M{
		"$push": bson.M{
			"students": username,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	if stat := user.DBD.Enroll(username, courseID.Hex()); stat == status.FAILED {
		return model.InternalServerException{Message: "database couldn't add user"}
	}
	return nil
}

//todo fix Inconsistency
func (c CourseMongoDriver) DelStd(username string, courseID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id": courseID,
	}
	change := bson.M{
		"$pull": bson.M{
			"students": username,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	if stat := user.DBD.Leave(username, courseID.Hex()); stat == status.FAILED {
		return model.InternalServerException{Message: "database couldn't add user"}
	}
	return nil
}

//todo fix Inconsistency
func (c CourseMongoDriver) DelTa(username string, courseID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id": courseID,
	}
	change := bson.M{
		"$pull": bson.M{
			"tas": username,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	if stat := user.DBD.Leave(username, courseID.Hex()); stat == status.FAILED {
		return model.InternalServerException{Message: "database couldn't add user"}
	}
	return nil
}

func (c CourseMongoDriver) PromoteDemoteUser(course *course.Course) error {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	target := bson.M{
		"_id": course.ID,
	}
	change := bson.M{
		"$set": bson.M{
			"students": course.StdUns,
			"tas":      course.TaUns,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return model.InternalServerException{Message: "database Internal Error:/n" + err.Error()}
	}
	return nil
}

func NewCourseMongoDriver(db, collection string) *CourseMongoDriver {
	return &CourseMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
