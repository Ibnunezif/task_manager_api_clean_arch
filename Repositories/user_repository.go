package Repositories

import (
	"context"
	"errors"
	"task-manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Create(user *Domain.User) (*Domain.User, error)
	GetByUsername(username string) (*Domain.User, error)
	UpdateRole(username, role string) error
}

type mongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(col *mongo.Collection) UserRepository {
	return &mongoUserRepository{collection: col}
}

func (r *mongoUserRepository) Create(u *Domain.User) (*Domain.User, error) {
	result, err := r.collection.InsertOne(context.Background(), bson.M{
		"username": u.Username(),
		"password": u.Password(),
		"role":     u.Role(),
	})
	if err != nil {
		return nil, err
	}
	u.SetID(result.InsertedID.(primitive.ObjectID).Hex())
	return u, nil
}

func (r *mongoUserRepository) GetByUsername(username string) (*Domain.User, error) {
	var u struct {
		ID       primitive.ObjectID `bson:"_id"`
		Username string             `bson:"username"`
		Password string             `bson:"password"`
		Role     string             `bson:"role"`
	}
	err := r.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&u)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	user := Domain.NewUser(u.ID.Hex(), u.Username, u.Password, u.Role)
	return user, nil
}

func (r *mongoUserRepository) UpdateRole(username, role string) error {
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"username": username}, bson.M{"$set": bson.M{"role": role}})
	return err
}