package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	dao "github.com/shaquib-nesar/task-planner-service.git/dao"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	congifSetup()
	dbSetup()
}

func main() {

	router := gin.Default()

	router.GET("/api/getTaskById", getTaskById)
	router.GET("/api/getAllTaskOfUser", getAllTaskOfUserByUid)
	router.GET("/api/getSubTaskListByParentId", getSubTaskListById)

	router.POST("/api/createTask", createTask)

	router.DELETE("/api/deleteTaskById")
	router.DELETE("/api/deleteAllTaskOfUser")

	router.PUT("/api/updateTaskById")

	router.Run(":" + viper.GetString("server.port"))
}

func getAllTaskOfUserByUid(c *gin.Context) {

	query := c.Request.URL.Query()
	userId := query["userId"][0]

	result := dao.GetAllTaskByUserId(collection, ctx, userId)
	if len(result) != 0 {
		c.JSON(200, result)
	} else {
		c.JSON(404, "No Record found with userId "+userId)
	}
}

func getTaskById(c *gin.Context) {

	query := c.Request.URL.Query()
	id := query["id"][0]

	doc, hasData := dao.GetTaskById(collection, ctx, id)

	if hasData {
		c.JSON(404, "No Record found with id "+id)
	} else {
		c.JSON(200, doc)
	}
}

func getSubTaskListById(c *gin.Context) {

	query := c.Request.URL.Query()
	id := query["id"][0]

	doc, hasData := dao.GetSubTaskListById(collection, ctx, id)

	if hasData {
		c.JSON(404, "No Sub Task found with parentId "+id)
	} else {
		c.JSON(200, doc)
	}
}

func createTask(c *gin.Context) {

	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
	}

	insertedId, err := dao.ProcessDocument(collection, ctx, bodyBytes)
	if err != nil {
		c.String(http.StatusBadRequest, "Hello %s")
	}

	fmt.Println(insertedId)
	c.JSON(200, insertedId)
}

func congifSetup() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}

func dbSetup() {
	clientOptions := options.Client().ApplyURI(viper.GetString("database.uri"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database(viper.GetString("database.name")).Collection(viper.GetString("database.collection"))
}
