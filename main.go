package main

import (
	"flag"
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
	debugKeys := flag.Bool("debug", false, "debug keypresses")
	flag.Parse()

	stk.Push(0)

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
			case ev.Key == termbox.KeyBackspace2:
				// Ctrl+Backspace
				if buf == "" {
					stk.Pop()
				} else {
					buf = ""
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
			case ev.Ch == '[':
				stk.Push(bufferOrPop())
				stk.Rotate()
			case ev.Ch == ']':
				stk.Push(bufferOrPop())
				stk.Unrotate()
			case ev.Ch == '\\':
				a := bufferOrPop()
				b := stk.Pop()
				stk.Push(a)
				stk.Push(b)
			default:
				if *debugKeys {
					buf = fmt.Sprintf("%#v", ev)
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		draw()
	}
}

func draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	prefix := "> "
	var i int
	stk.Walk(func(v float64) {
		drawString(len(prefix), i, fmt.Sprintf("%f", v))
		i++
	})
	drawString(len(prefix), stk.Len(), buf)
	arrowHeight := stk.Len() - 1
	if len(buf) > 0 {
		arrowHeight++
	}
	drawString(0, arrowHeight, prefix)
	termbox.SetCursor(len(buf)+len(prefix), stk.Len())
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
