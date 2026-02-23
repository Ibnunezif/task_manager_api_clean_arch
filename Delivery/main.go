package main

import (
	"context"
	"log"
	"task-manager/Delivery/controllers"
	"task-manager/Delivery/routers"
	"task-manager/Repositories"
	"task-manager/Usecases"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	DB_URL := "mongodb+srv://abdulbasitnezif_db_user:IbnuN3z1f@cluster0.qhmvuqi.mongodb.net/go_test?appName=Cluster0"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DB_URL))
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("task_manager_db")
	taskCol := db.Collection("tasks")
	userCol := db.Collection("users")

	// ===== Task layers =====
	taskRepo := Repositories.NewMongoTaskRepository(taskCol)
	taskUsecase := Usecases.NewTaskUsecase(taskRepo)
	taskCtrl := controllers.NewTaskController(taskUsecase)

	// ===== User layers =====
	userRepo := Repositories.NewMongoUserRepository(userCol)
	userUsecase := Usecases.NewUserUsecase(userRepo)
	userCtrl := controllers.NewUserController(userUsecase)

	// ===== Router =====
	r := routers.SetupRouter(taskCtrl, userCtrl)
	r.Run(":8080")
}