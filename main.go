package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var client *mongo.Client
var coll *mongo.Collection

func InitMongodb() (err error) {
	// 1.连接mongodb
	client, err = mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://mongodb:Gqr18335825724@14.103.150.223:27017").
			SetConnectTimeout(5*time.Second))
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		fmt.Println(err)
		return err
	}

	// 由于都是对同一个集合进行操作，所以在初始化mongodb时就选择了集合，防止后续出现大量重复代码
	coll = client.Database("").Collection("stu")
	return
}

type Student struct {
	ID   int    `bson:"_id"`
	Age  int    `bson:"age"`
	Name string `bson:"name"`
	Addr string `bson:"address"`
}

func InsertDocument(dbName, collectionName string, doc interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	collection := client.Database(dbName).Collection(collectionName)

	_, err := collection.InsertOne(ctx, doc)

	return err

}

func main() {

	stus := []interface{}{
		Student{
			ID:   1,
			Age:  18,
			Name: "alice",
			Addr: "beijing",
		},
		Student{
			ID:   2,
			Age:  19,
			Name: "bob",
			Addr: "shanghai",
		},
		Student{
			ID:   3,
			Age:  20,
			Name: "charlie",
			Addr: "guangzhou",
		},
	}
	err := InitMongodb()
	if err != nil {
		return
	}
	err = InsertDocument("db", "stu", stus)
	if err != nil {
		return
	}
}
