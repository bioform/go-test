package gameone

// Answer describes one of question answers
type Answer struct {
	text      string `yaml:"text"`
	nextScene Scene  `yaml:"next scene"`
}
