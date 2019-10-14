package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"

	termbox "github.com/nsf/termbox-go"
)

//go:generate stacker -type float64

var (
	stk float64Stack
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
		draw()
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch {
			case ev.Key == termbox.KeyCtrlC:
				break loop
			case ev.Key == termbox.KeyEnter:
				pushBuffer()
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
				pushBuffer()
				b, _ := stk.Pop()
				a, _ := stk.Pop()
				stk.Push(a + b)
			case ev.Ch == '_':
				pushBuffer()
				a, _ := stk.Pop()
				stk.Push(-a)
			case ev.Ch == '-':
				pushBuffer()
				b, _ := stk.Pop()
				a, _ := stk.Pop()
				stk.Push(a - b)
			case ev.Ch == '*':
				pushBuffer()
				b, _ := stk.Pop()
				a, _ := stk.Pop()
				stk.Push(a * b)
			case ev.Ch == '/':
				pushBuffer()
				b, _ := stk.Pop()
				a, _ := stk.Pop()
				stk.Push(a / b)
			case ev.Ch == '%':
				pushBuffer()
				v, _ := stk.Pop()
				stk.Push(math.Sqrt(v))
			case ev.Ch == '^':
				pushBuffer()
				b, _ := stk.Pop()
				a, _ := stk.Pop()
				stk.Push(math.Pow(a, b))
			case ev.Ch == '[':
				pushBuffer()
				stk.Unrotate()
			case ev.Ch == ']':
				pushBuffer()
				stk.Rotate()
			case ev.Ch == '\\':
				pushBuffer()
				a, _ := stk.Pop()
				b, _ := stk.Pop()
				stk.Push(a)
				stk.Push(b)
			case ev.Ch == 'c':
				pushBuffer()
				v, _ := stk.Peek()
				clipboard.WriteAll(fmt.Sprintf("%f", v))
			case ev.Ch == 'v':
				clip, _ := clipboard.ReadAll()
				lines := strings.Split(strings.NewReplacer(
					"\r", "\n",
					", ", "\n",
					"\t", "\n",
				).Replace(clip), "\n")
				for _, line := range lines {
					line := strings.TrimSpace(line)
					if _, err := strconv.ParseFloat(line, 64); err == nil {
						pushBuffer()
						buf = line
					}
				}
			default:
				if *debugKeys {
					buf = fmt.Sprintf("%#v", ev)
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func draw() {
	t := termboxTranslate{}
	_, height := termbox.Size()
	// need a little bit of padding for the cursor
	height--
	if stk.Len() > height {
		t.y = height - stk.Len()
	}

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	prefix := "> "
	var i int
	stk.Walk(func(v float64) {
		t.DrawString(len(prefix), i, fmt.Sprintf("%f", v))
		i++
	})
	t.DrawString(len(prefix), stk.Len(), buf)
	arrowHeight := stk.Len() - 1
	if len(buf) > 0 {
		arrowHeight++
	}
	t.DrawString(0, arrowHeight, prefix)
	t.SetCursor(len(buf)+len(prefix), stk.Len())
	termbox.Flush()
}

func pushBuffer() {
	if len(buf) > 0 {
		val, _ := strconv.ParseFloat(buf, 64)
		buf = ""
		stk.Push(val)
	}
}
