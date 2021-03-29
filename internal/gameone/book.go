package gameone

import (
	"fmt"
	"strings"

	"github.com/bioform/go-test/internal/gameone/storage"
	"github.com/inancgumus/screen"
	"github.com/manifoldco/promptui"
)

// Book is an entry point for a game scenario
type Book struct {
	Desc       string   `yaml:"Описание"`
	Scene      Scene    `yaml:"Сцена первая"`
	Dictionary *[]Scene `yaml:"Содержание"`
	Scenes     map[string]*Scene
}

// Run will run the game
func (book *Book) Run(scene *Scene) error {
	if scene == nil {
		scene = &book.Scene
	}

	fmt.Println(book.Desc)
	fmt.Scanln()

	var err error
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

		// save last position
		err = storage.GetInstance().WriteLastPosition(scene.ID)
		if err != nil {
			return err
		}

		index := -1
		for index < 0 {
			prompt := promptui.Select{
				Label:     "Что выбрать,",
				Items:     question.Answers,
				Templates: AnswerConsoleTemplate,
			}

			index, _, err = prompt.Run()

			if err != nil {
				fmt.Println("Что-то пошло не так...", err)
				return nil
			}
		}

		answer := question.Answers[index]
		scene = answer.Scene
	}

	return nil
}
