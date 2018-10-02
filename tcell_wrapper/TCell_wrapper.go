package tcell_wrapper

import (
	"github.com/gdamore/tcell"
)

var (
	fg_color = tcell.ColorWhite
	bg_color = tcell.ColorBlack
	style    tcell.Style
	screen   tcell.Screen
)

/* PUBLIC SHIT BELOW */

const (
	BLACK        = 0
	DARK_RED     = 1
	DARK_GREEN   = 2
	DARK_YELLOW  = 3
	DARK_BLUE    = 4
	DARK_MAGENTA = 5
	DARK_CYAN    = 6
	BEIGE        = 7
	DARK_GRAY    = 8
	RED          = 9
	GREEN        = 10
	YELLOW       = 11
	BLUE         = 12
	MAGENTA      = 13
	CYAN         = 14
	WHITE        = 15
)

func Init_console() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	var e error
	screen, e = tcell.NewScreen()
	if e != nil {
		panic(e)
	}
	if e = screen.Init(); e != nil {
		panic(e)
	}
	style = tcell.StyleDefault.Foreground(fg_color).Background(bg_color)
	screen.SetStyle(style)
}

func Close_console() { //should be deferred!
	screen.Fini()
}

func Clear_console() {
	screen.Clear()
}

func Flush_console() {
	screen.Show()
}

//
//
//func Run_event_listener() { // should be run as go func() {}()
//	for {
//		event_queue <- termbox.PollEvent()
//	}
//}
//

func SetColor(fg int, bg int) {
	fg_color = tcell.Color(fg)
	bg_color = tcell.Color(bg)
	style = style.Foreground(fg_color).Background(bg_color)
}

func SetFgColor(fg int) {
	fg_color = tcell.Color(fg)
	style = style.Foreground(fg_color).Background(bg_color)
}

func SetBgColor(bg int) {
	bg_color = tcell.Color(bg)
	style = style.Foreground(fg_color).Background(bg_color)
}

func PutChar(c rune, x, y int) {
	screen.SetCell(x, y, style, c)
}

func PutString(s string, x, y int) {
	length := len([]rune(s))
	for i := 0; i < length; i++ {
		PutChar([]rune(s)[i], x+i, y)
	}
}

//
//
//
//func Await_keypress() rune {
//	for {
//		ev := <-event_queue
//		if ev.Type == termbox.EventKey {
//			return ev.Ch
//		}
//	}
//
//}
//
func ReadKey() string {
	ev := screen.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEscape:
			return "ESCAPE"
		case tcell.KeyEnter:
			return "ENTER"
		case tcell.KeyTab:
			return "TAB"
		default:
			return string(ev.Rune())
		}
	case *tcell.EventResize:
		screen.Sync()
		return "NONKEY_SYNC_EVENT"
	}
	return "KEY_EMPTY_WTF_HAPPENED"
}
