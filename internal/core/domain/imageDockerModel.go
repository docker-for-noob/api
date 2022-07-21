package domain

type DockerHub struct {
	DockerHubData []byte `json:"dockerHubData"`
}

type DockerHubImage struct {
	Results []string `njson:"results"`
}

type DockerHubTags struct {
	Tag string `njson:"name"`
}

type DockerImageResult struct {
	Name string
	Tags []string
}
