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

	collectAllScenes(&book)

	return &book, nil
}

func collectAllScenes(book *Book) {
	book.Scenes = map[string]*Scene{}

	scenes := []*Scene{&book.Scene}
	for len(scenes) > 0 {

		// clone scenes and clear the original slice
		children := make([]*Scene, len(scenes))
		copy(children, scenes)
		scenes = nil

		// add all children to AllScenes and fill in new scenes for processing
		for _, scene := range children {
			_, isProcessed := book.Scenes[scene.ID]

			if !isProcessed {
				book.Scenes[scene.ID] = scene

				// prepare next scenes for processing
				if scene.Question != nil {
					for _, answer := range scene.Question.Answers {
						scenes = append(scenes, answer.Scene)
					}
				}
			}
		}
	}
}
