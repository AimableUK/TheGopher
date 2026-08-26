package controllers

import (
	"context"
	"go-mongo-api/configs"
	"go-mongo-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateBook - Create a new book with transaction
func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.ID = primitive.NewObjectID()
	book.PublishedDate = time.Now()

	session, err := configs.DB.StartSession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer session.EndSession(context.Background())

	err = mongo.WithSession(context.Background(), session, func(sessCtx mongo.SessionContext) error {
		collection := configs.GetCollection(configs.DB, "books")
		_, err := collection.InsertOne(sessCtx, book)
		if err != nil {
			return err
		}

		// Example: Update author's book count (if needed)
		authorsColl := configs.GetCollection(configs.DB, "authors")
		_, err = authorsColl.UpdateOne(
			sessCtx,
			bson.M{"_id": book.AuthorID},
			bson.M{"$inc": bson.M{"book_count": 1}},
		)
		return err
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// GetBook - Get a book by ID
func GetBook(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	var book models.Book
	collection := configs.GetCollection(configs.DB, "books")

	err := collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&book)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// ListBooks - List books with cursor-based pagination
func ListBooks(c *gin.Context) {
	var books []models.Book
	collection := configs.GetCollection(configs.DB, "books")

	limit, err := primitive.ParseInt64(c.DefaultQuery("limit", "10"), 10, 64)
	if err != nil {
		limit = 10
	}

	cursorID := c.Query("cursor")
	filter := bson.M{}
	if cursorID != "" {
		objID, _ := primitive.ObjectIDFromHex(cursorID)
		filter = bson.M{"_id": bson.M{"$gt": objID}}
	}

	findOptions := options.Find().SetLimit(limit)
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var book models.Book
		cursor.Decode(&book)
		books = append(books, book)
	}

	c.JSON(http.StatusOK, gin.H{
		"books":       books,
		"next_cursor": books[len(books)-1].ID.Hex(),
	})
}

// ListBooksByCategory - Get book counts grouped by category
func ListBooksByCategory(c *gin.Context) {
	collection := configs.GetCollection(configs.DB, "books")

	// Aggregate books by category and count them
	pipeline := []bson.M{
		{
			"$group": bson.M{
				"_id":   "$category",       // Group by category
				"count": bson.M{"$sum": 1}, // Count books in each category
			},
		},
		{
			"$sort": bson.M{
				"count": -1, // Sort by count in descending order
			},
		},
	}

	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	var categories []bson.M
	if err = cursor.All(context.TODO(), &categories); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// SearchBooks - Perform full-text search on books
func SearchBooks(c *gin.Context) {
	query := c.Query("q")
	collection := configs.GetCollection(configs.DB, "books")

	filter := bson.M{
		"$text": bson.M{
			"$search": query,
		},
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	var books []models.Book
	for cursor.Next(context.TODO()) {
		var book models.Book
		cursor.Decode(&book)
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

// GetAuthorBooks - Example of joining authors and books using aggregation
func GetAuthorBooks(c *gin.Context) {
	authorID := c.Param("authorId")
	objID, _ := primitive.ObjectIDFromHex(authorID)

	pipeline := mongo.Pipeline{
		bson.D{{"$match", bson.D{{"author_id", objID}}}},
		bson.D{{"$lookup", bson.D{
			{"from", "authors"},
			{"localField", "author_id"},
			{"foreignField", "_id"},
			{"as", "author_details"},
		}}},
		bson.D{{"$unwind", "$author_details"}},
	}

	collection := configs.GetCollection(configs.DB, "books")
	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	var books []bson.M
	for cursor.Next(context.TODO()) {
		var book bson.M
		cursor.Decode(&book)
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

// UpdateBook - Update a book's info
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := configs.GetCollection(configs.DB, "books")
	update := bson.M{
		"$set": bson.M{
			"title":          book.Title,
			"author_id":      book.AuthorID,
			"published_date": book.PublishedDate,
			"details":        book.Details,
		},
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

// UpdateBookDetails - Update a field inside the JSONB 'details'
func UpdateBookDetails(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := configs.GetCollection(configs.DB, "books")
	update := bson.M{
		"$set": bson.M{"details": updateData},
	}

	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book details updated"})
}

// DeleteBook - Delete a book by ID
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(id)

	collection := configs.GetCollection(configs.DB, "books")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
