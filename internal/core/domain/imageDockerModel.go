package domain

type DockerHub struct {
	DockerHubData []byte `json:"dockerHubData"`
}

type DockerImages struct {
	Images []string `njson:"results"`
}

type DockerImageVersions struct {
	Name     string
	Versions []string
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
