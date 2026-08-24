package yaml

type Manifest struct {
	Database Database `yaml:"database"`
}

type includes struct {
	Files []map[string]string `yaml:"manifest"`
}
