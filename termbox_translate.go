package main

import termbox "github.com/nsf/termbox-go"

type termboxTranslate struct {
	x, y int
}

func (t termboxTranslate) DrawString(x, y int, str string) {
	for i, ch := range str {
		termbox.SetCell(x+t.x+i, y+t.y, ch, termbox.ColorWhite, termbox.ColorDefault)
	}
}

func (t termboxTranslate) SetCursor(x, y int) {
	termbox.SetCursor(x+t.x, y+t.y)
}
