package gameone

import (
	"os"

	"gopkg.in/yaml.v2"
)

// LoadBook loads scenario from YAML file
func LoadBook(configPath string) (*Book, error) {
	// Create config structure
	book := Book{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)
	d.SetStrict(true)

	// Start YAML decoding from file
	if err := d.Decode(&book); err != nil {
		return nil, err
	}

	return &book, nil
}
