package images

type DockerHubAuthorization struct {
	Token      string
	Expires_in int16
	Issued_at  string
}

type DockerImageManifest struct {
	Name string
	Tag  string
	FsLayers []map[string]string
	Errors []map[string]string
}
