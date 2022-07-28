package domain

type ImageReference struct {
	Name string `bson:"Name"`
	Workdir []string `bson:"Workdir"`
	Port []string `bson:"Port"`
	Env []EnvVar `bson:"Env"`
}

type ImageNameDetail struct {
	Name     string   `bson:"Name"`
	Language string   `bson:"Language"`
	Version  string   `bson:"Version"`
	Tags     []string `bson:"Tags"`
}

type EnvVar struct {
	Key string
	Desc string
}
