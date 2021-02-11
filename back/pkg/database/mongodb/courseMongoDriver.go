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
	"yes-sharifTube/internal/model"
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/pkg/database"
)

type CourseMongoDriver struct {
	collection *mongo.Collection
}

var doOnce sync.Once

func (c CourseMongoDriver) GetAll(courseIDs []primitive.ObjectID) ([]*course.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	res := make([]*course.Course, len(courseIDs))
	for i, cID := range courseIDs {
		target := bson.M{"_id": cID}
		res[i] = &course.Course{}
		if err := c.collection.FindOne(ctx, target).Decode(res[i]); err != nil {
			return nil, database.ThrowCourseNotFoundException(courseIDs[i].Hex())
		}
	}
	return res, nil
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
		return nil, database.ThrowInternalDBException(err.Error())
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
		return nil, database.ThrowInternalDBException(err.Error())
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

func (c CourseMongoDriver) Insert(username string, title, token string, summery *string) (*course.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	cou, err := course.New(primitive.NewObjectID(), title, username, token, summery)
	if err != nil {
		return nil, err
	}
	if _, err = c.collection.InsertOne(ctx, cou); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return cou, nil
}

func (c CourseMongoDriver) Update(username string, courseID primitive.ObjectID, title, token, summery *string) (*course.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := c.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	if !fc.IsUserNotStudent(username) {
		return nil, database.ThrowUserNotAllowedException(username)
	}
	err := fc.Update(title, summery, token)
	if err != nil {
		return nil, err
	}
	change := bson.M{
		"$set": bson.M{
			"title":   fc.Title,
			"summery": fc.Summery,
			"token":   fc.Token,
		},
	}
	if _, err = c.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return &fc, nil
}

func (c CourseMongoDriver) Delete(username string, courseID primitive.ObjectID) (*course.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := c.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	if !fc.IsUserProfessor(username) {
		return nil, database.ThrowUserNotAllowedException(username)
	}
	if _, err := c.collection.DeleteOne(ctx, target); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return &fc, nil
}

func (c CourseMongoDriver) AddUser(username, token string, courseID primitive.ObjectID) (*course.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := c.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	if fc.IsUserParticipateInCourse(username) {
		return nil, database.ThrowDuplicateUsernameException()
	}
	if !fc.CheckCourseToken(token) {
		return nil, database.ThrowIncorrectTokenException()
	}
	fc.StdUns = append(fc.StdUns, username)
	change := bson.M{
		"$set": bson.M{
			"students": fc.StdUns,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return &fc, nil
}

func (c CourseMongoDriver) DeleteUser(username, targetUsername string, courseID primitive.ObjectID) (*course.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection
	//todo check if the targetUsername exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := c.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	if !fc.IsUserParticipateInCourse(targetUsername) {
		return nil, database.ThrowUserNotFoundException(targetUsername)
	}
	if !fc.IsUserAllowedToDeleteUser(username, targetUsername) {
		return nil, database.ThrowUserNotAllowedException(username)
	}
	var change bson.M
	if model.ContainsInStringArray(fc.TaUns, targetUsername) {
		fc.TaUns = model.RemoveFromStringArray(fc.TaUns, targetUsername)
		change = bson.M{
			"$set": bson.M{
				"tas": fc.TaUns,
			},
		}
	} else {
		fc.StdUns = model.RemoveFromStringArray(fc.StdUns, targetUsername)
		change = bson.M{
			"$set": bson.M{
				"students": fc.StdUns,
			},
		}
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return &fc, nil
}

func (c CourseMongoDriver) PromoteToTA(username, targetUsername string, courseID primitive.ObjectID) (*course.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection
	//todo check if the targetUsername exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := c.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	if !fc.IsUserNotStudent(username) {
		return nil, database.ThrowUserNotAllowedException(username)
	}
	if !fc.IsUserStudent(targetUsername) {
		return nil, database.ThrowUserIsNotSTDException(targetUsername)
	}
	fc.StdUns = model.RemoveFromStringArray(fc.StdUns, targetUsername)
	fc.TaUns = append(fc.TaUns, targetUsername)
	change := bson.M{
		"$set": bson.M{
			"tas":      fc.TaUns,
			"students": fc.StdUns,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return &fc, nil
}

func (c CourseMongoDriver) DemoteToSTD(username, targetUsername string, courseID primitive.ObjectID) (*course.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection
	//todo check if the targetUsername exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := c.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	if !fc.IsUserNotStudent(username) {
		return nil, database.ThrowUserNotAllowedException(username)
	}
	if !fc.IsUserTA(targetUsername) {
		return nil, database.ThrowUserIsNotTAException(targetUsername)
	}
	fc.TaUns = model.RemoveFromStringArray(fc.TaUns, targetUsername)
	fc.StdUns = append(fc.StdUns, targetUsername)
	change := bson.M{
		"$set": bson.M{
			"tas":      fc.TaUns,
			"students": fc.StdUns,
		},
	}
	if _, err := c.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return &fc, nil
}

func NewCourseMongoDriver(db, collection string) *CourseMongoDriver {
	return &CourseMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
