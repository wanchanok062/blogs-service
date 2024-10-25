package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/wanchanok6698/web-blogs/api/v1/models"
	"github.com/wanchanok6698/web-blogs/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BlogService struct {
	Collection *mongo.Collection
}

func NewBlogService() (*BlogService, error) {
	collection, err := config.GetBlogCollection()
	if err != nil {
		return nil, errors.New("failed to get blog collection: " + err.Error())
	}
	return &BlogService{Collection: collection}, nil
}

func (bs *BlogService) GetAllBlogs(ctx context.Context, filterOptions models.FilterBlogsOptions) ([]models.BlogPost, error) {
	var blogs []models.BlogPost

	filter := bson.M{}
	if filterOptions.AuthorID != "" {
		filter["authorId"] = filterOptions.AuthorID
	}

	if filterOptions.Search != "" {
		filter["$or"] = bson.A{
			bson.M{"title": bson.M{"$regex": filterOptions.Search, "$options": "i"}},
			bson.M{"content": bson.M{"$regex": filterOptions.Search, "$options": "i"}},
		}
	}

	sortOptions := bson.M{}

	if filterOptions.Sort != "" {
		sortFields := strings.Split(filterOptions.Sort, ":")
		log.Println(sortFields)
		field := sortFields[0]
		order := 1
		if len(sortFields) > 1 && sortFields[1] == "desc" {
			order = -1
		}
		sortOptions[field] = order
	}

	cursor, err := bs.Collection.Find(ctx, filter, options.Find().SetSort(sortOptions))
	if err != nil {
		return nil, err
	}

	cursor.All(ctx, &blogs)
	return blogs, nil
}

func (bs *BlogService) GetBlogByID(ctx context.Context, id string) (*models.BlogPost, error) {
	var blog models.BlogPost

	filter := bson.M{"id": id}
	err := bs.Collection.FindOne(ctx, filter).Decode(&blog)
	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func (bs *BlogService) CreateBlog(ctx context.Context, blog *models.BlogPost) (*models.BlogPost, error) {
	blog.PostID = primitive.NewObjectID().Hex()
	blog.CreatedAt = time.Now()
	blog.UpdatedAt = time.Now()

	_, err := bs.Collection.InsertOne(ctx, blog)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (bs *BlogService) UpdateBlog(ctx context.Context, id string, updatedBlog *models.BlogUpdate) (*models.BlogUpdate, error) {
	filter := bson.M{"id": id}

	update := bson.M{"$set": bson.M{}}
	setMap := update["$set"].(bson.M)

	if updatedBlog.Title != "" {
		setMap["title"] = updatedBlog.Title
	}
	if updatedBlog.Content != "" {
		setMap["content"] = updatedBlog.Content
	}
	if updatedBlog.AuthorID != "" {
		setMap["authorId"] = updatedBlog.AuthorID
	}
	setMap["updatedAt"] = time.Now()

	_, err := bs.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	err = bs.Collection.FindOne(ctx, filter).Decode(updatedBlog)
	if err != nil {
		return nil, err
	}
	return updatedBlog, nil
}

func (bs *BlogService) DeleteBlog(ctx context.Context, id string) error {
	filter := bson.M{"id": id}
	var blog models.BlogPost
	err := bs.Collection.FindOne(ctx, filter).Decode(&blog)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fmt.Errorf("ไม่พบบล็อกที่มี ID: %s", id)
		}
		return err
	}
	_, err = bs.Collection.DeleteOne(ctx, filter)
	return err
}
