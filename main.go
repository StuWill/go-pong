package main

import (
	"log"

	tcell "github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()

	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorBlack)
	screen.SetStyle(defStyle)

	for {

		screen.SetContent(0, 0, 'H', nil, defStyle)
		screen.SetContent(1, 0, 'i', nil, defStyle)
		screen.SetContent(2, 0, '!', nil, defStyle)

		screen.Show()
	}
}
