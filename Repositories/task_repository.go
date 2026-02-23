package Repositories

import (
	"context"
	"errors"
	"task-manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	Create(task *Domain.Task) (*Domain.Task, error)
	GetAll() ([]*Domain.Task, error)
	GetByID(id string) (*Domain.Task, error)
	Update(task *Domain.Task) (*Domain.Task, error)
	Delete(id string) error
}

type mongoTaskRepository struct {
	collection *mongo.Collection
}

func NewMongoTaskRepository(col *mongo.Collection) TaskRepository {
	return &mongoTaskRepository{collection: col}
}

func (r *mongoTaskRepository) Create(t *Domain.Task) (*Domain.Task, error) {
	result, err := r.collection.InsertOne(context.Background(), bson.M{
		"title":       t.Title,
		"description": t.Description,
		"status":      t.Status,
	})
	if err != nil {
		return nil, err
	}
	t.ID=result.InsertedID.(primitive.ObjectID).Hex()
	return t, nil
}

func (r *mongoTaskRepository) GetAll() ([]*Domain.Task, error) {
	ctx := context.Background()
	cur, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var tasks []*Domain.Task
	for cur.Next(ctx) {
		var t struct {
			ID          primitive.ObjectID `bson:"_id"`
			Title       string             `bson:"title"`
			Description string             `bson:"description"`
			Status      string             `bson:"status"`
		}
		if err := cur.Decode(&t); err != nil {
			return nil, err
		}
		task := Domain.NewTask(t.ID.Hex(), t.Title, t.Description)
		task.Update("", "", t.Status)
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *mongoTaskRepository) GetByID(id string) (*Domain.Task, error) {
	ctx := context.Background()
	objID, _ := primitive.ObjectIDFromHex(id)
	var t struct {
		ID          primitive.ObjectID `bson:"_id"`
		Title       string             `bson:"title"`
		Description string             `bson:"description"`
		Status      string             `bson:"status"`
	}
	err := r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&t)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	task := Domain.NewTask(t.ID.Hex(), t.Title, t.Description)
	task.Update("", "", t.Status)
	return task, nil
}

func (r *mongoTaskRepository) Update(t *Domain.Task) (*Domain.Task, error) {
	objID, _ := primitive.ObjectIDFromHex(t.ID)
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{
		"$set": bson.M{
			"title":       t.Title,
			"description": t.Description,
			"status":      t.Status,
		},
	})
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *mongoTaskRepository) Delete(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	return err
}