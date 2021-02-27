package gameone

import (
	"fmt"
	"strings"

	"github.com/inancgumus/screen"
	"github.com/manifoldco/promptui"
)

// Book is an entry point for a game scenario
type Book struct {
	Desc       string   `yaml:"Описание"`
	Scene      Scene    `yaml:"Сцена первая"`
	Dictionary *[]Scene `yaml:"Содержание"`
}

// Run will run the game
func (book *Book) Run() {
	fmt.Println(book.Desc)

	fmt.Scanln()

	scene := &book.Scene

	for scene != nil {
		screen.Clear()
		screen.MoveTopLeft()

		fmt.Println(scene.Desc)

		question := scene.Question
		if question == nil {
			break
		}

		fmt.Println(strings.Repeat("-", 40)) // ----------------------------------
		fmt.Println(question.Text)

		index := -1
		var err error

		for index < 0 {
			prompt := promptui.Select{
				Label:     "Что выбрать,",
				Items:     question.Answers,
				Templates: AnswerConsoleTemplate,
			}

			index, _, err = prompt.Run()

			if err != nil {
				fmt.Println("Что-то пошло не так...", err)
				return
			}
		}

		answer := question.Answers[index]
		scene = &answer.Scene
	}
}
