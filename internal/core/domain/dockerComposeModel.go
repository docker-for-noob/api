package domain

import "github.com/google/uuid"

type DockerCompose struct{
	Id                 string `json:"id"`
	DockerComposeDatas []byte `json:"dockerData"`
}

func NewDockerCompose(data []byte) DockerCompose{
	return DockerCompose{
		Id: uuid.New().String(),
		DockerComposeDatas: data,
		}
}
