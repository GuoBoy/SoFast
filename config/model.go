package config

type CModel struct {
	Port       uint   `yaml:"port"`
	UploadPath string `yaml:"upload_path"`
	Dev        bool   `yaml:"dev"`
}
