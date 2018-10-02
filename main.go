package main

import (
	cw "TCellConsoleWrapper/tcell_wrapper"
	"fmt"
)

//func makebox(s tcell.Screen) {
//	w, h := s.Size()
//
//	if w == 0 || h == 0 {
//		return
//	}
//
//	glyphs := []rune{'@', '#', '&', '*', '=', '%', 'Z', 'A'}
//
//	lx := rand.Int() % w
//	ly := rand.Int() % h
//	lw := rand.Int() % (w - lx)
//	lh := rand.Int() % (h - ly)
//	st := tcell.StyleDefault
//	gl := ' '
//	if s.Colors() > 256 {
//		rgb := tcell.NewHexColor(int32(rand.Int() & 0xffffff))
//		st = st.Background(rgb)
//	} else if s.Colors() > 1 {
//		st = st.Background(tcell.Color(rand.Int() % s.Colors()))
//	} else {
//		st = st.Reverse(rand.Int()%2 == 0)
//		gl = glyphs[rand.Int()%len(glyphs)]
//	}
//}

func main() {

	cw.Init_console()
	cw.Clear_console()
	defer cw.Close_console()

	test_wrapper()
	//for {
	//	str := cw.ReadKey()
	//	cw.PutString(str, 0, 0)
	//	cw.Flush_console()
	//	if str == "ESCAPE" {
	//		break
	//	}
	//}

//	quit := make(chan struct{})
//	go func() {
//		for {
//			ev := s.PollEvent()
//			switch ev := ev.(type) {
//			case *tcell.EventKey:
//				switch ev.Key() {
//				case tcell.KeyEscape, tcell.KeyEnter:
//					close(quit)
//					return
//				case tcell.KeyCtrlL:
//					s.Sync()
//				}
//			case *tcell.EventResize:
//				s.Sync()
//			}
//		}
//	}()
//
//	cnt := 0
//	dur := time.Duration(0)
//loop:
//	for {
//		select {
//		case <-quit:
//			break loop
//		case <-time.After(time.Millisecond * 50):
//		}
//		start := time.Now()
//		makebox(s)
//		cnt++
//		dur += time.Now().Sub(start)
//	}
}


func test_wrapper() {
	const linesPerColumn = 32
	x := 0
	for i := 0; i < 256; i++ {
		if i % linesPerColumn == 0 && i != 0{
			x+= 9
		}
		cw.SetColor(i, 0)
		cw.PutString(fmt.Sprintf("CODE %d", i), x, i%linesPerColumn)
	}
	cw.Flush_console()
	for key_pressed := cw.ReadKey(); key_pressed != "ENTER"; {
		if key_pressed == "ESCAPE" {
			break
		} else {
			key_pressed = cw.ReadKey()
		}
	}
}
