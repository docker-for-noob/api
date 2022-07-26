package referenceBatchHandler

import (
	"encoding/csv"
	"fmt"
	"github.com/docker-generator/api/internal/core/ports"
	"github.com/docker-generator/api/pkg/goDotEnv"
	"log"
	"os"
	"time"
)

type referenceBatchHandler struct {
	imageReferenceService ports.ImageReferenceService
}

func NewReferenceBatchHandler(imageReferenceService ports.ImageReferenceService) *referenceBatchHandler {
	return &referenceBatchHandler{
		imageReferenceService: imageReferenceService,
	}
}

func (handler *referenceBatchHandler) Run() error {

	fmt.Println("Début du batch : ", time.Now().Format("2006-01-02 15:04:05"))

	pathToInputData := goDotEnv.GetEnvVariable("BATCH_REFERENTIEL_INPUT")

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

	err = handler.imageReferenceService.AddAllTagReference(csvData[0])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[OK] - Fin du batch : ", time.Now().Format("2006-01-02 15:04:05"))
	return nil
}