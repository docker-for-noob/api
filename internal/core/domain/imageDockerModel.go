package domain

type DockerHub struct {
	DockerHubData []byte `json:"dockerHubData"`
}

type DockerImages struct {
	Next    string `njson:"next"`
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
}

type DockerImagesParse struct {
	Images []string
}

type DockerImageVersions struct {
	Name     string
	Versions []string
}

type ImageNameDetail struct {
	Name     string   `bson:"Name"`
	Language string   `bson:"Language"`
	Version  string   `bson:"Version"`
	Tags     []string `bson:"Tags"`
}

type DockerHubImage struct {
	Next    string   `njson:"next"`
	Results []string `njson:"results"`
}

type DockerHubTags struct {
	Tag string `njson:"name"`
}

type DockerImageResult struct {
	Name string
	Tags []string
}
