package gameone

// Question describes a scene question
type Question struct {
	text    string   `yaml:"text"`
	answers []Answer `yaml:"answers"`
}
