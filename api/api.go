package api

type Image struct {
	Platform     Platform     `json:"platform,omitempty"`
	Dependencies []Dependency `json:"dependencies,omitempty"`
}

type Dependency string

type Platform string

const (
	PlatformJVM Platform = "jvm"
)
