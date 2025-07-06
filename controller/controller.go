package controller

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPosts(c *gin.Context) {
	mongoURI := os.Getenv("MONGODB_CONN_URI")
	if mongoURI == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Mongo URI not set"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "MongoDB connection failed"})
		return
	}
	defer client.Disconnect(ctx)

	collection := client.Database("iw-admin-db-v0").Collection("posts-v0")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	defer cursor.Close(ctx)

	var posts []bson.M
	if err := cursor.All(ctx, &posts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cursor decode error"})
		return
	}

	c.JSON(http.StatusOK, posts)
}
