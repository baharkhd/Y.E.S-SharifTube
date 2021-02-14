package controller

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/model/comment"
	"yes-sharifTube/internal/model/content"
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/internal/model/user"
)

func CreateComment(authorUsername, contentID, body string, repliedID *string) (*comment.Comment, *comment.Reply, error) {
	// check if user exists in database
	if _, err := user.Get(authorUsername); err != nil {
		return nil, nil, err
	}
	// get the content from database
	con, err := content.Get(nil, contentID)
	if err != nil {
		return nil, nil, err
	}
	// get the course from database
	cr, err := course.Get(con.CourseID)
	if err != nil {
		return nil, nil, err
	}
	// can user insert any comments
	err = cr.IsUserAllowedToInsertComment(authorUsername)
	if err != nil {
		return nil, nil, err
	}
	// create comment or reply
	var cmd *comment.Comment = nil
	var rep *comment.Reply = nil
	if repliedID == nil {
		cmd, err = comment.New(body, authorUsername, contentID)
		if err != nil {
			return nil, nil, err
		}
		// insert comment in database
		cmd, err = comment.InsertComment(con.CourseID, contentID, cmd)
		if err != nil {
			return nil, nil, err
		}
	} else {
		// check if the comment exists in the
		repID, err := primitive.ObjectIDFromHex(*repliedID)
		if err != nil {
			return nil, nil, model.InternalServerException{Message: err.Error()}
		}
		if ccmt, _ := con.GetComment(repID); ccmt == nil {
			return nil, nil, model.CommentNotFoundException{Message: "comment not found"}
		}
		// create new rely
		rep, err = comment.NewReply(body, authorUsername, *repliedID)
		if err != nil {
			return nil, nil, err
		}
		// insert reply in database
		rep, err = comment.InsertReply(con.CourseID, contentID, *repliedID, rep)
		if err != nil {
			return nil, nil, err
		}
	}
	return cmd, rep, err
}

func UpdateComment(authorUsername, contentID, commentID string, newBody *string) (*comment.Comment, *comment.Reply, error) {
	// check if user exists in database
	if _, err := user.Get(authorUsername); err != nil {
		return nil, nil, err
	}
	// get the content from database
	con, err := content.Get(nil, contentID)
	if err != nil {
		return nil, nil, err
	}
	// get the course from database
	cr, err := course.Get(con.CourseID)
	if err != nil {
		return nil, nil, err
	}
	cmdID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return nil, nil, model.InternalServerException{Message: err.Error()}
	}
	ccmt, rcmt := con.GetComment(cmdID)
	if ccmt == nil {
		return nil, nil, model.CommentNotFoundException{Message: "the target comment not found"}
	}
	if rcmt == nil {
		// get the comments
		err = cr.IsUserAllowedToUpdateComment(authorUsername, ccmt.AuthorUn)
		if err != nil {
			return nil, nil, err
		}
		err = ccmt.Update(newBody)
		if err != nil {
			return nil, nil, err
		}
	} else {
		err = cr.IsUserAllowedToUpdateComment(authorUsername, rcmt.AuthorUn)
		if err != nil {
			return nil, nil, err
		}
		err = rcmt.Update(newBody)
		if err != nil {
			return nil, nil, err
		}
	}
	// update the comments in database
	err = comment.Update(con.CourseID, contentID, con.Comments)
	if err != nil {
		return nil, nil, err
	}
	if rcmt != nil {
		return nil, rcmt, nil
	}
	return ccmt, nil, nil
}

func DeleteComment(authorUsername, contentID, commentID string) (*comment.Comment, *comment.Reply, error) {
	// check if user exists in database
	if _, err := user.Get(authorUsername); err != nil {
		return nil, nil, err
	}
	// get the content from database
	con, err := content.Get(nil, contentID)
	if err != nil {
		return nil, nil, err
	}
	// get the course from database
	cr, err := course.Get(con.CourseID)
	if err != nil {
		return nil, nil, err
	}
	// get comments
	cmdID, err := primitive.ObjectIDFromHex(commentID)
	if err != nil {
		return nil, nil, model.InternalServerException{Message: err.Error()}
	}
	ccmt, rcmt, err := cr.RemoveComment(authorUsername, cmdID, con)
	if err != nil {
		return nil, nil, err
	}
	// update the comments in database
	err = comment.Update(con.CourseID, contentID, con.Comments)
	if err != nil {
		return nil, nil, err
	}
	if rcmt != nil {
		return nil, rcmt, nil
	}
	return ccmt, nil, nil
}
