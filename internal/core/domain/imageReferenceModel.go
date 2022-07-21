package domain

import "github.com/google/uuid"

type ImageReference struct {
	Id uuid.UUID
	Name string
	Workdir []string
	Port []string
	Env []EnvVar
}

type EnvVar struct {
	Key string
	Desc string
}
