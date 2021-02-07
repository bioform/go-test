package gameone

import (
	"fmt"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/manifoldco/promptui"
)

// Scene describes a game scene
type Scene struct {
	desc     string   `yaml:"desc"`
	question Question `yaml:"question"`
}

// Run will run the game
func (initialScene Scene) Run() {
	scene := &initialScene

	for scene != nil {
		ui.Clear()

		fmt.Println(scene.desc)
		fmt.Println(strings.Repeat("-", 40)) // ----------------------------------
		fmt.Println(scene.question.text)

		index := -1
		var err error

		for index < 0 {
			prompt := promptui.Select{
				Label: "Что выбрать,",
				Items: scene.question.answers,
			}

			index, _, err = prompt.Run()

			if err != nil {
				fmt.Println("Что-то пошло не так...")
				return
			}
		}

		answer := scene.question.answers[index]
		scene = &answer.nextScene

		/*
			reader := bufio.NewReader(os.Stdin)

			for i, answer := range scene.question.answers {


				fmt.Printf("|%-6d|%s\n", i+1, answer)
				fmt.Printf("|%-6d|%s\n", 0, "Exit")
				fmt.Print("-> ")

				for selection, _ := reader.ReadString('\n'); len(selection) != 0; {
					if selection == "0" {
						os.Exit(0)
					}

					number, err := strconv.ParseInt(selection, 10, 0)
					if err != nil {
						fmt.Println("Попробуйте ещё раз")
						fmt.Print("-> ")
					} else {
						answer := scene.question.answers[number]
						scene = &answer.nextScene
					}
				}
			}
		*/
	}
}
