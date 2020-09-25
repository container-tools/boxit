package api

type ImageRequest struct {
	Platform     Platform     `json:"platform,omitempty"`
	Dependencies []Dependency `json:"dependencies,omitempty"`
}

type ImageResult struct {
	ID        string     `json:"id,omitempty"`
	Artifacts []Artifact `json:"artifact,omitempty"`
}

type Dependency string

type Platform string

const (
	PlatformJVM Platform = "jvm"
)

type Artifact struct {
	//ID       string `json:"id,omitempty"`
	//Checksum string `json:"checksum,omitempty"`
	//Target   string `json:"target,omitempty"`
	Location string `json:"location,omitempty"`
}
