package domain

type ImageReference struct {
	Name string `bson:"Name"`
	Workdir []string `bson:"Workdir"`
	Port []string `bson:"Port"`
	Env []EnvVar `bson:"Env"`
}

type EnvVar struct {
	Key string
	Desc string
}
