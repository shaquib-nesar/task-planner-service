package main

// DATABASE
// func main() {
// 	fmt.Println("S T A R T")
// 	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://shaquib:pxLJVm1VqTxogY8o@cluster.zd5et.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Disconnect(ctx)
// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	databases, err := client.ListDatabaseNames(ctx, bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(databases)

// 	database := client.Database("task-planner-db")

// 	collection := database.Collection("task-collection")
// 	// ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
// 	res, err := collection.InsertOne(ctx, bson.D{{"_id", "pi_3.1459"}, {"name", "pi"}, {"value", 3.14159}})
// 	id := res.InsertedID

// 	fmt.Println(id)

// 	fmt.Println("E N D")
// }
