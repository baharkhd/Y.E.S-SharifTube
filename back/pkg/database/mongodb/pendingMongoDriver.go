package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"yes-sharifTube/internal/model/content"
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/internal/model/pending"
	"yes-sharifTube/pkg/database"
)

type PendingMongoDriver struct {
	collection *mongo.Collection
}

func (p PendingMongoDriver) GetByFilter(courseID *primitive.ObjectID, status *pending.Status, uploaderUsername *string, start, amount int) ([]*pending.Pending, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	filter := bson.A{}
	if courseID != nil {
		filter = append(filter,
			bson.M{
				"$eq": bson.A{
					"$$item.course",
					courseID.Hex(),
				},
			},
		)
	}
	if status != nil {
		filter = append(filter,
			bson.M{
				"$eq": bson.A{
					"$$item.status",
					status,
				},
			},
		)
	}
	if uploaderUsername != nil {
		filter = append(filter,
			bson.M{
				"$eq": bson.A{
					"$$item.uploaded_by",
					uploaderUsername,
				},
			},
		)
	}
	pipeline := []bson.M{
		{
			"$project": bson.M{
				"created_at": 1,
				"inventory":  1,
				"prof":       1,
				"students":   1,
				"summery":    1,
				"tas":        1,
				"title":      1,
				"token":      1,
				"contents":   1,
				"pends": bson.M{
					"$filter": bson.M{
						"input": "$pends",
						"as":    "item",
						"cond": bson.M{
							"$and": filter,
						},
					},
				},
			},
		},
	}

	courr, err := p.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	defer courr.Close(ctx)
	var pendings []*pending.Pending
	i := 0
	for courr.Next(context.Background()) {
		var ctmp course.Course
		_ = courr.Decode(&ctmp)
		for j, _ := range ctmp.Pends {
			pendings = append(pendings, ctmp.Pends[j])
		}
		i++
	}
	return pending.GetAll(pendings, start, amount), nil

}

func (p PendingMongoDriver) Insert(username string, courseID primitive.ObjectID, title, furl string, description *string) (*pending.Pending, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := p.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	if !fc.IsUserStudent(username) {
		return nil, database.ThrowUserNotAllowedException(username)
	}
	pnt, err := pending.New(primitive.NewObjectID(), title, username, furl, courseID.Hex(), description)
	if err != nil {
		return nil, err
	}
	change := bson.M{
		"$push": bson.M{
			"pends": pnt,
		},
	}
	if _, err = p.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return pnt, nil
}

func (p PendingMongoDriver) Update(username string, courseID, pendingID primitive.ObjectID, title, description *string) (*pending.Pending, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := p.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	pcnt := fc.GetPending(pendingID)
	if pcnt == nil {
		return nil, database.ThrowPendingNotFoundException(pendingID.Hex())
	}
	if !fc.IsUserStudent(username) || username != pcnt.UploadedByUn {
		return nil, database.ThrowUserNotAllowedException(username)
	}
	if pcnt.Status != pending.PENDING {
		return nil, database.ThrowOfferedContentNotPendingException(pendingID.Hex())
	}
	err := pcnt.Update(title, description)
	if err != nil {
		return nil, err
	}
	target = bson.M{
		"_id":       courseID,
		"pends._id": pcnt.ID,
	}
	change := bson.M{
		"$set": bson.M{
			"pends.$.title":       pcnt.Title,
			"pends.$.description": pcnt.Description,
			"pends.$.timestamp":   pcnt.Timestamp,
		},
	}
	if _, err = p.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return pcnt, nil
}

func (p PendingMongoDriver) Delete(username string, courseID, pendingID primitive.ObjectID) (*pending.Pending, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := p.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	pcnt := fc.GetPending(pendingID)
	if pcnt == nil {
		return nil, database.ThrowPendingNotFoundException(pendingID.Hex())
	}
	if !fc.IsUserStudent(username) || username != pcnt.UploadedByUn {
		return nil, database.ThrowUserNotAllowedException(username)
	}
	change := bson.M{
		"$pull": bson.M{
			"pends": bson.M{
				"_id": pendingID,
			},
		},
	}
	if _, err := p.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return pcnt, nil
}

func (p PendingMongoDriver) Accept(username string, courseID, pendingID primitive.ObjectID, title, description *string) (*pending.Pending, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := p.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	if !fc.IsUserNotStudent(username) {
		return nil, database.ThrowUserNotAllowedException(username)
	}
	pcnt := fc.GetPending(pendingID)
	if pcnt == nil {
		return nil, database.ThrowPendingNotFoundException(pendingID.Hex())
	}
	if pcnt.Status != pending.PENDING {
		return nil, database.ThrowOfferedContentNotPendingException(pendingID.Hex())
	}
	err := pcnt.Update(title, description)
	if err != nil {
		return nil, err
	}
	ncnt, err := content.New(primitive.NewObjectID(), pcnt.Title, pcnt.UploadedByUn, pcnt.Furl, pcnt.CourseID, &pcnt.Description, &username, []string{})
	if err != nil {
		return nil, err
	}
	change := bson.M{
		"$push": bson.M{
			"contents": ncnt,
		},
	}
	if _, err = p.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}

	pcnt.Accept()
	target = bson.M{
		"_id":       courseID,
		"pends._id": pcnt.ID,
	}
	change = bson.M{
		"$set": bson.M{
			"pends.$.title":       pcnt.Title,
			"pends.$.description": pcnt.Description,
			"pends.$.timestamp":   pcnt.Timestamp,
			"pends.$.status":      pending.ACCEPTED,
		},
	}
	if _, err = p.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return pcnt, nil
}

func (p PendingMongoDriver) Reject(username string, courseID, pendingID primitive.ObjectID) (*pending.Pending, error) {
	ctx, cancel := context.WithTimeout(context.Background(), LongTimeOut*time.Millisecond)
	defer cancel()

	//todo check if the username exists in user collection

	var fc course.Course
	target := bson.M{
		"_id": courseID,
	}
	if err := p.collection.FindOne(ctx, target).Decode(&fc); err != nil {
		return nil, database.ThrowCourseNotFoundException(courseID.Hex())
	}
	if !fc.IsUserNotStudent(username) {
		return nil, database.ThrowUserNotAllowedException(username)
	}
	pcnt := fc.GetPending(pendingID)
	if pcnt == nil {
		return nil, database.ThrowPendingNotFoundException(pendingID.Hex())
	}
	if pcnt.Status != pending.PENDING {
		return nil, database.ThrowOfferedContentNotPendingException(pendingID.Hex())
	}

	pcnt.Reject()
	target = bson.M{
		"_id":       courseID,
		"pends._id": pendingID,
	}
	change := bson.M{
		"$set": bson.M{
			"pends.$.status": pending.REJECTED,
		},
	}
	if _, err := p.collection.UpdateOne(ctx, target, change); err != nil {
		return nil, database.ThrowInternalDBException(err.Error())
	}
	return pcnt, nil
}

func NewPendingMongoDriver(db, collection string) *PendingMongoDriver {
	return &PendingMongoDriver{
		collection: client.Database(db).Collection(collection),
	}
}
