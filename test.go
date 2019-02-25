package tcell_wrapper

//
////func makebox(s tcell.Screen) {
////	w, h := s.Size()
////
////	if w == 0 || h == 0 {
////		return
////	}
////
////	glyphs := []rune{'@', '#', '&', '*', '=', '%', 'Z', 'A'}
////
////	lx := rand.Int() % w
////	ly := rand.Int() % h
////	lw := rand.Int() % (w - lx)
////	lh := rand.Int() % (h - ly)
////	st := tcell.StyleDefault
////	gl := ' '
////	if s.Colors() > 256 {
////		rgb := tcell.NewHexColor(int32(rand.Int() & 0xffffff))
////		st = st.Background(rgb)
////	} else if s.Colors() > 1 {
////		st = st.Background(tcell.Color(rand.Int() % s.Colors()))
////	} else {
////		st = st.Reverse(rand.Int()%2 == 0)
////		gl = glyphs[rand.Int()%len(glyphs)]
////	}
////}

//func main() {
//	Init_console()
//	Clear_console()
//	defer Close_console()
//
//	// test mouse events
//	// var origx, origy, tox, toy int
//	moves := 0
//	for ReadKeyAsync() != "ESCAPE" {
//		// draw shit by mouse
//		// Clear_console()
//		//SetBgColor(WHITE)
//		//for x := origx; x < tox; x++ {
//		//	for y := origy; y < toy; y++ {
//		//		PutChar(' ', x, y)
//		//	}
//		//}
//
//		if WasMouseMovedSinceLastEvent() {
//			moves += 1
//		}
//
//		PutString(strconv.Itoa(moves), 0, 0)
//
//		// moves--
//
//		//if mouseButton != "NONE" {
//		//	if !mouseHeld {
//		//		origx, origy = GetMouseCoords()
//		//	}
//		//	tox, toy = GetMouseCoords()
//		//}
//
//		Flush_console()
//
//	}
//}

//func main() {
//
//	Init_console()
//	Clear_console()
//	defer Close_console()
//
//	test_wrapper()
//	//for {
//	//	str := ReadKey()
//	//	PutString(str, 0, 0)
//	//	Flush_console()
//	//	if str == "ESCAPE" {
//	//		break
//	//	}
//	//}
//
////	quit := make(chan struct{})
////	go func() {
////		for {
////			ev := s.PollEvent()
////			switch ev := ev.(type) {
////			case *tcell.EventKey:
////				switch ev.Key() {
////				case tcell.KeyEscape, tcell.KeyEnter:
////					close(quit)
////					return
////				case tcell.KeyCtrlL:
////					s.Sync()
////				}
////			case *tcell.EventResize:
////				s.Sync()
////			}
////		}
////	}()
////
////	cnt := 0
////	dur := time.Duration(0)
////loop:
////	for {
////		select {
////		case <-quit:
////			break loop
////		case <-time.After(time.Millisecond * 50):
////		}
////		start := time.Now()
////		makebox(s)
////		cnt++
////		dur += time.Now().Sub(start)
////	}
//}
//
//
//func test_wrapper() {
//	const linesPerColumn = 32
//	x := 0
//	for i := 0; i < 256; i++ {
//		if i % linesPerColumn == 0 && i != 0{
//			x+= 9
//		}
//		SetColor(i, 0)
//		PutString(fmt.Sprintf("CODE %d", i), x, i%linesPerColumn)
//	}
//	Flush_console()
//	for key_pressed := ReadKey(); key_pressed != "ENTER"; {
//		if key_pressed == "ESCAPE" {
//			break
//		} else {
//			key_pressed = ReadKey()
//		}
//	}
//}
