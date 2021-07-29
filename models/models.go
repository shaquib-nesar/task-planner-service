package models

import "time"

type Task struct {
	Title        string   `bson:"title" json:"title"`
	Parent_id    string   `bson:"parentId" json:"parentId"`
	Id           string   `bson:"_id" json:"_id"`
	Description  string   `bson:"description" json:"description"`
	CreatedTime  int64    `bson:"createdTime" json:"createdTime"`
	StartTime    int64    `bson:"startTime" json:"startTime"`
	ModifiedTime int64    `bson:"modifiedTime" json:"modifiedTime"`
	EndTime      int64    `bson:"endTime" json:"endTime"`
	Priority     string   `bson:"priority" json:"priority"`
	SubTask      []string `bson:"subTask" json:"subTask"`
	UserId       string   `bson:"userId" json:"userId"`
}

func GetTestTask() Task {

	return Task{
		Title: "Test Title", Parent_id: "",
		Id:          "HASTALAVISTA",
		Description: "Beware, This is a test task.",
		CreatedTime: time.Now().Unix(),
		StartTime:   time.Now().Unix(),
		EndTime:     time.Now().Unix(),
		Priority:    "medium",
		SubTask:     []string{"090LK", "AB09PO"},
		UserId:      "USR938"}
}
