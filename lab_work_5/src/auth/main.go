package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	jwtKey = []byte("your_secret_key")
)

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var client *mongo.Client

func main() {

	credential := options.Credential{
		Username: "root",
		Password: "admin",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var dbError error

	client, dbError = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017/?connect=direct").SetAuth(credential))

	log.Print(client)

	if dbError != nil {
		log.Fatal("Connect db error")
	}

	var result bson.M

	if err := client.Database("test").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}

	router := gin.Default()
	router.POST("/login", login)
	router.POST("/signup", signup)

	authorized := router.Group("/auth")
	authorized.Use(authMiddleware)
	{
		authorized.GET("", welcome)
	}

	router.Run(":8082")
}

func login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	collection := client.Database("test").Collection("users")
	var existingUser User
	err := collection.FindOne(context.Background(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if user.Password != existingUser.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func signup(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	collection := client.Database("test").Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func authMiddleware(c *gin.Context) {
	tokenHeader := c.GetHeader("Authorization")
	if tokenHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
		return
	}

	// parts := strings.Split(tokenHeader, " ")
	// if len(parts) != 2 || parts[0] != "Bearer" {
	// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
	// 	return
	// }

	tokenString := tokenHeader
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	c.Next()
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the API"})
}
