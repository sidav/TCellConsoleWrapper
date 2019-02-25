package tcell_wrapper

import (
	"github.com/gdamore/tcell"
	"time"
)

var (
	fg_color       = tcell.ColorWhite
	bg_color       = tcell.ColorBlack
	style          tcell.Style
	screen         tcell.Screen
	CONSOLE_WIDTH  = 80
	CONSOLE_HEIGHT = 25
	wasResized     = false
	evCh           chan tcell.Event
	flushesCounter int

	mouseX, mouseY             int
	mouseVectorX, mouseVectorY int // for getting mouse coords changes
	mouseButton                string
	mouseHeld                  bool
	mouseMoved                 bool
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
	screen.EnableMouse()
	style = tcell.StyleDefault.Foreground(fg_color).Background(bg_color)
	screen.SetStyle(style)
	CONSOLE_WIDTH, CONSOLE_HEIGHT = screen.Size()
	evCh = make(chan tcell.Event, 1)
	go startAsyncEventListener()
}

func Close_console() { //should be deferred!
	screen.Fini()
}

func Clear_console() {
	screen.Clear()
}

func Flush_console() {
	flushesCounter++
	screen.Show()
}

func GetConsoleSize() (int, int) {
	return screen.Size()
}

func WasResized() bool {
	if wasResized {
		wasResized = false
		return true
	}
	return false
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

func ReadKey() string {
	for {
		for len(evCh) == 0 {
			time.Sleep(1 * time.Millisecond)
		}
		ev := <-evCh
		switch ev := ev.(type) {
		case *tcell.EventKey:
			return eventToKeyString(ev)
		case *tcell.EventResize:
			screen.Sync()
			CONSOLE_WIDTH, CONSOLE_HEIGHT = screen.Size()
			wasResized = true
		}
	}
	return "KEY_EMPTY_WTF_HAPPENED"
}

func ReadKeyAsync() string { // also reads mouse events... TODO: think of if separate mouse events reader is needed.
	if len(evCh) == 0 {
		return "NOTHING"
	}
	ev := <-evCh
	switch ev := ev.(type) {
	case *tcell.EventKey:
		return eventToKeyString(ev)
	case *tcell.EventMouse:
		mouseEventWork(ev)
	case *tcell.EventResize:
		screen.Sync()
		CONSOLE_WIDTH, CONSOLE_HEIGHT = screen.Size()
		wasResized = true
	}
	return "NON-KEY"
}

func eventToKeyString(ev *tcell.EventKey) string {
	switch ev.Key() {
	case tcell.KeyUp:
		return "UP"
	case tcell.KeyRight:
		return "RIGHT"
	case tcell.KeyDown:
		return "DOWN"
	case tcell.KeyLeft:
		return "LEFT"
	case tcell.KeyEscape:
		return "ESCAPE"
	case tcell.KeyEnter:
		return "ENTER"
	case tcell.KeyTab:
		return "TAB"
	case tcell.KeyDelete:
		return "DELETE"
	case tcell.KeyInsert:
		return "INSERT"
	case tcell.KeyEnd:
		return "END"
	case tcell.KeyHome:
		return "HOME"
	default:
		return string(ev.Rune())
	}
}

func mouseEventWork(ev *tcell.EventMouse) {
	mx, my := ev.Position()
	if mouseX != mx || mouseY != my {
		mouseVectorX = mx-mouseX
		mouseVectorY = my-mouseY
		mouseX, mouseY = mx, my
		mouseMoved = true
	}
	// PrevMouseButton = mouseButton
	switch ev.Buttons() {
	case tcell.ButtonNone:
		mouseButton = "NONE"
		mouseHeld = false
	case tcell.Button1:
		mouseHeld = mouseButton != "NONE"
		mouseButton = "LEFT"
	case tcell.Button2:
		mouseHeld = mouseButton != "NONE"
		mouseButton = "RIGHT"
	}
}

func GetMouseCoords() (int, int) {
	return mouseX, mouseY
}

func GetMouseButton() string {
	return mouseButton
}

func IsMouseHeld() bool {
	return mouseHeld
}

func WasMouseMovedSinceLastEvent() bool {
	t := mouseMoved
	mouseMoved = false
	return t
}

func GetMouseMovementVector() (int, int) {
	return mouseVectorX, mouseVectorY
}

func startAsyncEventListener() {
	for {
		ev := screen.PollEvent()
		select {
		case evCh <- ev:
		default:
		}
	}
}

func GetNumberOfRecentFlushes() int { // may be useful for searching rendering overkills and something
	t := flushesCounter
	flushesCounter = 0
	return t
}

func PrintCharactersTable() {
	for x := 0; x < CONSOLE_WIDTH; x++ {
		for y := 0; y < CONSOLE_HEIGHT; y++ {
			PutChar(rune(x+y*CONSOLE_WIDTH), x, y)
		}
	}
	Flush_console()
	ReadKey()
}
