package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/pkg/goDotEnv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type imageReferenceRepository struct {}

func NewImageReferenceRepository() *imageReferenceRepository {
	return &imageReferenceRepository{}
}

func (repository *imageReferenceRepository) Read(imageName string) (domain.ImageReference, error) {

	mongoUri := goDotEnv.GetEnvVariable("MONGO_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("reference").Collection("reference")



	var result bson.M

	err = coll.FindOne(context.TODO(), bson.D{{"Name", imageName}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", imageName)
		return domain.ImageReference{}, err
	}

	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}

	var response domain.ImageReference

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		panic(err)
	}

	return response, nil
}