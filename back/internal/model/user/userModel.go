package user
import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct{
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string
	Email    string
	Username string
	Password string
	Courses  []string
}