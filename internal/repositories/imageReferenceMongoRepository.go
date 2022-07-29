package repositories

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"github.com/docker-generator/api/internal/core/domain"
	"github.com/docker-generator/api/pkg/goDotEnv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strings"
)

type imageReferenceRepository struct {
	client *mongo.Client
}

type ImageReferenceToAdd struct {
	Name     string   `bson:"Name"`
	Language string   `bson:"Language"`
	Version  string   `bson:"Version"`
	Tags     []string `bson:"Tags"`
	Workdir []string `bson:"Workdir"`
	Port []string `bson:"Port"`
	Env []domain.EnvVar `bson:"Env"`
}

type Formater func(imageName string) (domain.ImageNameDetail, error)

func NewImageReferenceRepository() *imageReferenceRepository {
	mongoUri := goDotEnv.GetEnvVariable("MONGO_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		panic(err)
	}

	return &imageReferenceRepository{
		client: client,
	}
}

func (repository *imageReferenceRepository) Read(imageName string) (domain.ImageReference, error) {
	coll := repository.client.Database("docker-for-noob").Collection("reference")

	var result bson.M

	err := coll.FindOne(context.TODO(), bson.D{{"Name", imageName}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return domain.ImageReference{}, err
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return domain.ImageReference{}, err
	}

	var response domain.ImageReference

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		return domain.ImageReference{}, err
	}

	return response, nil
}

func (repository *imageReferenceRepository) Add(imageReference domain.ImageReference) error {

	coll := repository.client.Database("docker-for-noob").Collection("reference")
	doc := bson.D{{"Name", imageReference.Name}, {"Port", imageReference.Port}, {"workdir", imageReference.Workdir}, {"Env", imageReference.Env}}
	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}

	return nil
}

func (repository *imageReferenceRepository)  AddAllTagReferenceFromApi(fn Formater) error {

	coll := repository.client.Database("docker-for-noob").Collection("reference")

	pathToInputData := goDotEnv.GetEnvVariable("BATCH_REFERENTIEL_BUFFER")

	if _, err := os.Stat(pathToInputData);
		err != nil {
		return err
	}

	f, err := os.Open(pathToInputData)
	if err != nil {
		return err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		return err
	}

	var allReferenceToAdd []interface{}

	for _, element := range csvData {
		allReferenceToAdd = append(allReferenceToAdd, mapCsvResultToTagReferenceStruct(element, fn))
	}

	_, err = coll.InsertMany(context.TODO(), allReferenceToAdd)
	if err != nil {
		return err
	}

	return nil
}
func (repository *imageReferenceRepository)  FindAllPortForLanguageAndVersion(language string, version string) []domain.ImageNameDetail {
	coll := repository.client.Database("docker-for-noob").Collection("reference")

	filter := bson.D{{"Language", language}, {"Version", version}}

	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	var results []domain.ImageNameDetail
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func mapCsvResultToTagReferenceStruct(csvLine []string, fn Formater) ImageReferenceToAdd {

	splitedReferenceName, _ := fn(csvLine[0])

	tagReference := ImageReferenceToAdd{}
	tagReference.Name = splitedReferenceName.Name
	tagReference.Version = splitedReferenceName.Version
	tagReference.Language = splitedReferenceName.Language
	tagReference.Tags = splitedReferenceName.Tags
	tagReference.Port = strings.Fields(csvLine[1])
	tagReference.Workdir = strings.Fields(csvLine[2])
	return tagReference
}