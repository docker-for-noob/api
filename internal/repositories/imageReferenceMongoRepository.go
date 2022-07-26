package repositories

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/pkg/goDotEnv"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"strings"
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

	coll := client.Database("docker-for-noob").Collection("reference")

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

func (repository *imageReferenceRepository) Add(imageReference domain.ImageReference) error {
	mongoUri := goDotEnv.GetEnvVariable("MONGO_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("docker-for-noob").Collection("reference")
	doc := bson.D{{"Id", imageReference.Id}, {"Name", imageReference.Name}, {"Port", imageReference.Port}, {"workdir", imageReference.Workdir}, {"Env", imageReference.Env}}
	_, err = coll.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (repository *imageReferenceRepository)  AddAllTagReferenceFromApi() error {

	mongoUri := goDotEnv.GetEnvVariable("MONGO_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("docker-for-noob").Collection("reference")

	pathToInputData := goDotEnv.GetEnvVariable("BATCH_REFERENTIEL_BUFFER")

	if _, err := os.Stat(pathToInputData);
		err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(pathToInputData)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var allReferenceToAdd []interface{}

	for _, element := range csvData {
		allReferenceToAdd = append(allReferenceToAdd, mapCsvResultToTagReferenceStruct(element))
	}

	result, err := coll.InsertMany(context.TODO(), allReferenceToAdd)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	return nil
}

func mapCsvResultToTagReferenceStruct(csvLine []string) domain.ImageReference {

	tagReference := domain.ImageReference{}
	tagReference.Name = csvLine[0]
	tagReference.Id, _ = uuid.Parse(csvLine[1])
	tagReference.Port = strings.Fields(csvLine[2])
	tagReference.Workdir = strings.Fields(csvLine[3])
	return tagReference
}