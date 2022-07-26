package domain

import "github.com/google/uuid"

type ImageReference struct {
	Id uuid.UUID `bson:"Id"`
	Name string `bson:"Name"`
	Workdir []string `bson:"Workdir"`
	Port []string `bson:"Port"`
	Env []EnvVar `bson:"Env"`
}

func NewImageReference() ImageReference {
	return ImageReference{
		Id: uuid.New(),
	}
}

type EnvVar struct {
	Key string
	Desc string
}
