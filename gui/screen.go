package gui

import (
	"fmt"
	"math"

	"github.com/go-gl/glfw/v3.2/glfw"
)

type Screen struct {
	Window                                           *glfw.Window
	labels                                           []*Label
	buttons                                          []*Button
	entrys                                           []*Entry
	text                                             *Text
	Xpos, Ypos, MaxTextSize                          float64
	fontpngpath, fontjsonpath, buttonpath, entrypath string
	shift                                            bool
	keys                                             string
	keysShift                                        string
	Key                                              string
}

func (screen *Screen) InitGui(fontpngpath, fontjsonpath, buttonpath, entrypath string, maxtextsize float64) {
	screen.fontpngpath = fontpngpath
	screen.fontjsonpath = fontjsonpath
	screen.buttonpath = buttonpath
	screen.entrypath = entrypath
	screen.MaxTextSize = maxtextsize
	screen.text = NewText(screen)
	screen.keys = "`1234567890-=qwertyuiop[]\\asdfghjkl;'zxcvbnm,./ "
	screen.keysShift = "~!@#$%^&*()_+QWERTYUIOP{}|ASDFGHJKL:\"ZXCVBNM<>? "
	println(len(screen.keys), len(screen.keysShift))
}
func (screen *Screen) Clear() {
	screen.entrys = nil
	screen.buttons = nil
	screen.labels = nil
}
func NewScreen(window *glfw.Window) *Screen {
	s := Screen{}
	s.Window = window
	return &s
}
func (screen *Screen) Update() {
	for _, label := range screen.labels {
		if !label.Hide {
			screen.text.draw(label.Text, label.X, label.Y, label.Size, false, false, screen.Window)
		}
	}
	for _, button := range screen.buttons {
		button.draw(screen)
		if !button.Hide {
			screen.text.draw(button.Text, button.X+button.W/2, button.Y+button.H/2, button.textSize(screen), true, true, screen.Window)
		}
	}
	for _, entry := range screen.entrys {
		if !entry.Hide {
			entry.draw(screen)
			screen.text.draw(entry.Text, entry.X, entry.Y+entry.H/2, math.Min(entry.textSize(screen), screen.MaxTextSize), false, true, screen.Window)
		}
	}
}

func KeyName(key glfw.Key) string {
	name := glfw.GetKeyName(key, 0)
	if name != "" && name != " " {
		return name
	}
	switch key {
	case glfw.KeySpace:
		return "Space"
	case glfw.KeyLeftShift:
		return "LShift"
	case glfw.KeyRightShift:
		return "RShift"
	case glfw.KeyLeftControl:
		return "LControl"
	case glfw.KeyRightControl:
		return "RControl"
	case glfw.Key(0):
		return "LButton"
	case glfw.Key(1):
		return "RButton"
	}
	return "Unknown"
}

func (screen *Screen) MouseButtonCallback() func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	return func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		wd, ht := FramebufferSize(w)
		x, y := screen.Xpos/float64(wd)*2-1, -(screen.Ypos/float64(ht)*2 - 1)
		for _, b := range screen.buttons {
			if b.isInside(x, y) && button == glfw.MouseButtonLeft && action == glfw.Press {
				b.Command()
			}
		}
		for _, e := range screen.entrys {
			if e.KeyHolder && e.Focus {
				if action == glfw.Press {
					if button == glfw.MouseButtonRight {
						e.Key = glfw.Key(button)
						e.Text = KeyName(e.Key)
						e.Focus = false
					} else if button == glfw.MouseButtonLeft {
						e.Key = glfw.Key(button)
						e.Text = KeyName(e.Key)
						e.Focus = false
					}
				}
			}
			if e.isInside(x, y) && button == glfw.MouseButtonLeft && action == glfw.Press {
				e.Focus = true
			} else if button == glfw.MouseButtonLeft && action == glfw.Press {
				e.Focus = false
			}
		}
	}
}
func (screen *Screen) CursorPosCallback() func(w *glfw.Window, xpos, ypos float64) {
	return func(w *glfw.Window, xpos, ypos float64) {
		screen.Xpos = xpos
		screen.Ypos = ypos
		wd, ht := FramebufferSize(w)
		x, y := screen.Xpos/float64(wd)*2-1, -(screen.Ypos/float64(ht)*2 - 1)
		for _, b := range screen.buttons {
			b.mouseover = b.isInside(x, y)
		}
		for _, e := range screen.entrys {
			e.mouseover = e.isInside(x, y)
		}
	}
}
func (screen *Screen) KeyCallBack() func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	return func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		text := func(k string) {
			for _, e := range screen.entrys {
				if e.KeyHolder {
					if e.Focus {
						if k == "delete" {
							e.Text = "delete"
							e.Key = key
						} else if k == "enter" {
							e.Text = "enter"
							e.Key = key
						} else if k == "escape" {
							e.Focus = false
						} else {
							e.Text = k
							e.Key = key
						}
						e.Focus = false
					}
				} else {
					if e.Focus {
						if k == "delete" {
							e.Text = e.Text[0:max(len(e.Text)-1, 0)]
						} else if k == "enter" && e.Command != nil {
							e.Command()
						} else if k == "escape" {
							e.Focus = false
						} else if k == "enter" {
						} else {
							e.Text += k
						}
					}
				}
			}
		}
		// println(glfw.GetKeyName(key, scancode))
		shiftp := mods&glfw.ModShift == glfw.ModShift
		// altp := mods&glfw.ModAlt == glfw.ModAlt
		// superp := mods&glfw.ModSuper == glfw.ModSuper
		// ctrlp := mods&glfw.ModControl == glfw.ModControl
		h := ""
		switch action {
		case glfw.Press:
			switch key {
			case glfw.Key1:
				h = "1"
			case glfw.Key2:
				h = "2"
			case glfw.Key3:
				h = "3"
			case glfw.Key4:
				h = "4"
			case glfw.Key5:
				h = "5"
			case glfw.Key6:
				h = "6"
			case glfw.Key7:
				h = "7"
			case glfw.Key8:
				h = "8"
			case glfw.Key9:
				h = "9"
			case glfw.Key0:
				h = "0"
			case glfw.KeyQ:
				h = "q"
			case glfw.KeyW:
				h = "w"
			case glfw.KeyE:
				h = "e"
			case glfw.KeyR:
				h = "r"
			case glfw.KeyT:
				h = "t"
			case glfw.KeyY:
				h = "y"
			case glfw.KeyU:
				h = "u"
			case glfw.KeyI:
				h = "i"
			case glfw.KeyO:
				h = "o"
			case glfw.KeyP:
				h = "p"
			case glfw.KeyA:
				h = "a"
			case glfw.KeyS:
				h = "s"
			case glfw.KeyD:
				h = "d"
			case glfw.KeyF:
				h = "f"
			case glfw.KeyG:
				h = "g"
			case glfw.KeyH:
				h = "h"
			case glfw.KeyJ:
				h = "j"
			case glfw.KeyK:
				h = "k"
			case glfw.KeyL:
				h = "l"
			case glfw.KeyZ:
				h = "z"
			case glfw.KeyX:
				h = "x"
			case glfw.KeyC:
				h = "c"
			case glfw.KeyV:
				h = "v"
			case glfw.KeyB:
				h = "b"
			case glfw.KeyN:
				h = "n"
			case glfw.KeyM:
				h = "m"
			case glfw.KeyEnter:
				h = "enter"
			case glfw.KeyDelete:
				h = "delete"
			case glfw.KeyBackspace:
				h = "delete"
			case glfw.KeySpace:
				h = " "
			case glfw.KeyApostrophe:
				h = "'"
			case glfw.KeyKP0:
				h = "0"
			case glfw.KeyKP1:
				h = "1"
			case glfw.KeyKP2:
				h = "2"
			case glfw.KeyKP3:
				h = "3"
			case glfw.KeyKP4:
				h = "4"
			case glfw.KeyKP5:
				h = "5"
			case glfw.KeyKP6:
				h = "6"
			case glfw.KeyKP7:
				h = "7"
			case glfw.KeyKP8:
				h = "8"
			case glfw.KeyKP9:
				h = "9"
			case glfw.KeyKPEnter:
				h = "enter"
			case glfw.KeyKPEqual:
				h = "="
			case glfw.KeyEqual:
				h = "="
			case glfw.KeyComma:
				h = ","
			case glfw.KeySlash:
				h = "/"
			case glfw.KeyMinus:
				h = "-"
			case glfw.KeyEscape:
				h = "escape"
			case glfw.KeyPeriod:
				h = "."
			case glfw.KeyBackslash:
				h = "\\"
			case glfw.KeySemicolon:
				h = ";"
			case glfw.KeyGraveAccent:
				h = "`"
			case glfw.KeyLeftBracket:
				h = "["
			case glfw.KeyRightBracket:
				h = "]"
			case glfw.KeyLeftShift:
				h = "shift"
			case glfw.KeyRightShift:
				h = "shift"
			case glfw.KeyRightControl:
				h = "ctrl"
			case glfw.KeyLeftControl:
				h = "ctrl"
			case glfw.KeyLeftAlt:
				h = "alt"
			case glfw.KeyRightAlt:
				h = "alt"
			case glfw.KeyRightSuper:
				h = "super"
			case glfw.KeyLeftSuper:
				h = "super"
			}
		case glfw.Release:
			switch key {
			case glfw.Key1:
				h = "1"
			case glfw.Key2:
				h = "2"
			case glfw.Key3:
				h = "3"
			case glfw.Key4:
				h = "4"
			case glfw.Key5:
				h = "5"
			case glfw.Key6:
				h = "6"
			case glfw.Key7:
				h = "7"
			case glfw.Key8:
				h = "8"
			case glfw.Key9:
				h = "9"
			case glfw.Key0:
				h = "0"
			case glfw.KeyQ:
				h = "q"
			case glfw.KeyW:
				h = "w"
			case glfw.KeyE:
				h = "e"
			case glfw.KeyR:
				h = "r"
			case glfw.KeyT:
				h = "t"
			case glfw.KeyY:
				h = "y"
			case glfw.KeyU:
				h = "u"
			case glfw.KeyI:
				h = "i"
			case glfw.KeyO:
				h = "o"
			case glfw.KeyP:
				h = "p"
			case glfw.KeyA:
				h = "a"
			case glfw.KeyS:
				h = "s"
			case glfw.KeyD:
				h = "d"
			case glfw.KeyF:
				h = "f"
			case glfw.KeyG:
				h = "g"
			case glfw.KeyH:
				h = "h"
			case glfw.KeyJ:
				h = "j"
			case glfw.KeyK:
				h = "k"
			case glfw.KeyL:
				h = "l"
			case glfw.KeyZ:
				h = "z"
			case glfw.KeyX:
				h = "x"
			case glfw.KeyC:
				h = "c"
			case glfw.KeyV:
				h = "v"
			case glfw.KeyB:
				h = "b"
			case glfw.KeyN:
				h = "n"
			case glfw.KeyM:
				h = "m"
			case glfw.KeyEnter:
				h = "enter"
			case glfw.KeyDelete:
				h = "delete"
			case glfw.KeyBackspace:
				h = "delete"
			case glfw.KeySpace:
				h = " "
			case glfw.KeyApostrophe:
				h = "'"
			case glfw.KeyKP0:
				h = "0"
			case glfw.KeyKP1:
				h = "1"
			case glfw.KeyKP2:
				h = "2"
			case glfw.KeyKP3:
				h = "3"
			case glfw.KeyKP4:
				h = "4"
			case glfw.KeyKP5:
				h = "5"
			case glfw.KeyKP6:
				h = "6"
			case glfw.KeyKP7:
				h = "7"
			case glfw.KeyKP8:
				h = "8"
			case glfw.KeyKP9:
				h = "9"
			case glfw.KeyKPEnter:
				h = "enter"
			case glfw.KeyKPEqual:
				h = "="
			case glfw.KeyEqual:
				h = "="
			case glfw.KeyComma:
				h = ","
			case glfw.KeySlash:
				h = "/"
			case glfw.KeyMinus:
				h = "-"
			case glfw.KeyEscape:
				h = "escape"
			case glfw.KeyPeriod:
				h = "."
			case glfw.KeyBackslash:
				h = "\\"
			case glfw.KeySemicolon:
				h = ";"
			case glfw.KeyGraveAccent:
				h = "`"
			case glfw.KeyLeftBracket:
				h = "["
			case glfw.KeyRightBracket:
				h = "]"
			case glfw.KeyLeftShift:
				h = "shift"
			case glfw.KeyRightShift:
				h = "shift"
			case glfw.KeyRightControl:
				h = "ctrl"
			case glfw.KeyLeftControl:
				h = "ctrl"
			case glfw.KeyLeftAlt:
				h = "alt"
			case glfw.KeyRightAlt:
				h = "alt"
			case glfw.KeyRightSuper:
				h = "super"
			case glfw.KeyLeftSuper:
				h = "super"
			}
			screen.Key = fmt.Sprintf("release %v", h)
		}

		for f := range screen.keys {
			if string(screen.keys[f]) == h && shiftp {
				text(string(screen.keysShift[f]))
			} else if string(screen.keys[f]) == h {
				text(string(screen.keys[f]))
			}

		}
		if h == "delete" {
			text("delete")
		}
		if h == "enter" {
			text("enter")
		}
		screen.Key = h
	}
}
