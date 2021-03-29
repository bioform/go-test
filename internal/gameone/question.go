package gameone

// Question describes a scene question
type Question struct {
	Text    string    `yaml:"Текст"`
	Answers []*Answer `yaml:"Ответы"`
}
