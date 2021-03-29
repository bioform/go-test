package gameone

import (
	"fmt"

	"github.com/bioform/go-test/internal/gameone/os"
	"github.com/manifoldco/promptui"
)

// Answer describes one of question answers
type Answer struct {
	Text  string `yaml:"Текст"`
	Scene *Scene `yaml:"Сцена"`
}

// AnswerConsoleTemplate defines "promptui" console template for an answer
var AnswerConsoleTemplate = &promptui.SelectTemplates{
	Active:   os.Pointer + ` {{ .Text | cyan | bold }}`,
	Inactive: `   {{ .Text | cyan }}`,
	Selected: `{{ "` + os.Checkmark + `" | green | bold }} {{ .Text | cyan }}`,
	Help: fmt.Sprintf(`{{ "Используй стрелочки что бы выбрать ответ:" | faint }} {{ .NextKey | faint }} ` +
		`{{ .PrevKey | faint }} {{ .PageDownKey | faint }} {{ .PageUpKey | faint }} ` +
		`{{ if .Search }} {{ "and" | faint }} {{ .SearchKey | faint }} {{ "toggles search" | faint }}{{ end }}`),
}
