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
	Results []string
}

type DockerImageVersions struct {
	Details []DockerImageDetails
}

type DockerImageDetails struct {
	Name     string
	Language string
	Version  string
	Tags     []string
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
