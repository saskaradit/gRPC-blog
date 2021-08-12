package blog

import (
	"github.com/saskaradit/grpc-blog/blogpb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Blog data
type Blog struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id`
	Content  string             `bson:"content"`
	Title    string             `bson:"title`
}

// DataToBlogPb function converts a Blog struct to BlogPb Pointer
func DataToBlogPb(data *Blog) *blogpb.Blog {
	return &blogpb.Blog{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorID,
		Title:    data.Content,
		Content:  data.Content,
	}
}
