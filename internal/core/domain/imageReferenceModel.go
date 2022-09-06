package domain

type ImageReference struct {
	Name string
	Workdir []string
	Port []string
	Env []EnvVar
}

type EnvVar struct {
	Key string
	Desc string
}
