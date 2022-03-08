package domain

type DockerCompose struct{
	Id                  string `json:"id"`
	Name				string `json:"name"`
	DockerComposeDatas  string `json:"dockerData"`
}