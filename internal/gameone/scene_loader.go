package gameone

import (
	"os"

	"gopkg.in/yaml.v2"
)

// LoadScene loads initial scene from YAML file
func LoadScene(configPath string) (*Scene, error) {
	// Create config structure
	scene := &Scene{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&scene); err != nil {
		return nil, err
	}

	return scene, nil
}
