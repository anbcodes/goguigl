package gui

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Screen struct {
	window                                           *glfw.Window
	labels                                           []*Label
	buttons                                          []*Button
	entrys                                           []*Entry
	text                                             *Text
	Xpos, Ypos                                       float64
	fontpngpath, fontjsonpath, buttonpath, entrypath string
}

func (screen *Screen) InitGui(fontpngpath, fontjsonpath, buttonpath, entrypath string) {
	screen.fontpngpath = fontpngpath
	screen.fontjsonpath = fontjsonpath
	screen.buttonpath = buttonpath
	screen.entrypath = entrypath
	screen.text = NewText(screen)

}

func NewScreen(window *glfw.Window) *Screen {
	s := Screen{}
	s.window = window
	return &s
}
func (screen *Screen) Update() {
	for _, label := range screen.labels {
		screen.text.Draw(label.Text, label.X, label.Y, label.Size, false, false, screen.window)
	}
	for _, button := range screen.buttons {
		button.Draw(screen)
		screen.text.Draw(button.Text, button.X+button.W/2, button.Y+button.H/2, button.TextSize(screen), true, true, screen.window)
	}
}
func (screen *Screen) MouseButtonCallback() func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	return func(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
		wd, ht := FramebufferSize(w)
		x, y := screen.Xpos/float64(wd)*2-1, screen.Ypos/float64(ht)*2-1
		for _, b := range screen.buttons {
			if b.isInside(x, y) && button == glfw.MouseButtonLeft && action == glfw.Press {
				b.Command()
			}
		}
	}
}
func (screen *Screen) CursorPosCallback() func(w *glfw.Window, xpos, ypos float64) {
	return func(w *glfw.Window, xpos, ypos float64) {
		screen.Xpos = xpos
		screen.Ypos = ypos
		wd, ht := FramebufferSize(w)
		x, y := screen.Xpos/float64(wd)*2-1, screen.Ypos/float64(ht)*2-1
		for _, b := range screen.buttons {
			b.mouseover = b.isInside(x, y)
		}
	}
}
