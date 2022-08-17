package main

import (
	"log"
	"os"

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
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	screen.SetStyle(defStyle)

	go Run(screen, defStyle)

	for {
		switch event := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				screen.Fini()
				os.Exit(0)
			}
		}
	}

}

func Run(screen tcell.Screen, defStyle tcell.Style) {
	x := 0
	for {

		screen.SetContent(x, 10, 'H', nil, defStyle)
		screen.SetContent(x+1, 10, 'i', nil, defStyle)
		screen.SetContent(x+2, 10, '!', nil, defStyle)

		screen.Show()

		x++
	}
}
