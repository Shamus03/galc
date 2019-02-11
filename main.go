package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/shamus03/galc/stack"

	termbox "github.com/nsf/termbox-go"
)

var (
	stk stack.Float64Stack
	buf string
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch {
			case ev.Key == termbox.KeyCtrlC:
				break loop
			case ev.Key == termbox.KeyEnter:
				val, _ := strconv.ParseFloat(buf, 64)
				stk.Push(val)
				buf = ""
			case ev.Key == termbox.KeyBackspace:
				if len(buf) > 0 {
					buf = buf[:len(buf)-1]
				}
			case '0' <= ev.Ch && ev.Ch <= '9':
				buf += string(ev.Ch)
			case ev.Ch == '.':
				if !strings.ContainsRune(buf, '.') {
					buf += "."
				}
			case ev.Ch == '+':
				b := bufferOrPop()
				a := stk.Pop()
				stk.Push(a + b)
			case ev.Ch == '-':
				b := bufferOrPop()
				a := stk.Pop()
				stk.Push(a - b)
			case ev.Ch == '*':
				b := bufferOrPop()
				a := stk.Pop()
				stk.Push(a * b)
			case ev.Ch == '/':
				b := bufferOrPop()
				a := stk.Pop()
				stk.Push(a / b)
			case ev.Ch == '%':
				stk.Push(math.Sqrt(bufferOrPop()))
			case ev.Ch == '^':
				b := bufferOrPop()
				a := stk.Pop()
				stk.Push(math.Pow(a, b))
			}
		}
		draw()
	}
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	var i int
	stk.Walk(func(v float64) {
		drawString(0, i, fmt.Sprintf("%f", v))
		i++
	})
	drawString(0, stk.Len(), buf)
	termbox.SetCursor(len(buf), stk.Len())
	termbox.Flush()
}

func drawString(x, y int, str string) {
	for i, ch := range str {
		termbox.SetCell(x+i, y, ch, termbox.ColorWhite, termbox.ColorDefault)
	}
}

func bufferOrPop() float64 {
	if len(buf) > 0 {
		val, _ := strconv.ParseFloat(buf, 64)
		buf = ""
		return val
	}
	return stk.Pop()
}