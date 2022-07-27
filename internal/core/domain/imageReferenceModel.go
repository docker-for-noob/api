package domain

import "github.com/docker-generator/api/pkg/uidgen"

type ImageReference struct {
	Id uidgen.UIDGen `bson:"Id"`
	Name string `bson:"Name"`
	Workdir []string `bson:"Workdir"`
	Port []string `bson:"Port"`
	Env []EnvVar `bson:"Env"`
}

func NewImageReference() ImageReference {
	return ImageReference{
		Id: uidgen.New(),
	}
}

type EnvVar struct {
	Key string
	Desc string
}
