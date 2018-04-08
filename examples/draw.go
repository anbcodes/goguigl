package main

import (
	"runtime"

	"github.com/anbcodes/goguigl/gui"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func windowSizeCallback(w *glfw.Window, wd, ht int) {
	fwidth, fheight := gui.FramebufferSize(w)
	gl.Viewport(0, 0, int32(fwidth), int32(fheight))
}
func hi() {
	println("WORKED!!!")
}
func hi2() {
	println("WORKED!!!2")
}
func main() {
	// println(gl.LESS)
	runtime.LockOSThread()
	w := gui.InitGlfw(700, 500, "launcher")
	gui.InitOpenGL()
	w.SetSizeCallback(windowSizeCallback)
	screen := gui.NewScreen(w)
	screen.InitGui("extras/font/font.png", "extras/font/font.json", "extras/button.png", "extras/entry.png", 0.3)
	guimousebuttoncallback := screen.MouseButtonCallback()
	guicursorposcallback := screen.CursorPosCallback()
	guikeycallback := screen.KeyCallBack()
	w.SetMouseButtonCallback(guimousebuttoncallback)
	w.SetCursorPosCallback(guicursorposcallback)
	w.SetKeyCallback(guikeycallback)
	gui.NewButton(screen, "pick profile", -0.7, -0.7, 0.3, 0.1, 0.05, hi)
	gui.NewButton(screen, "new profile", -1, -0.7, 0.3, 0.1, 0.05, hi2)
	// gui.NewButton(screen, "hellohello", 0, 0.5, 1, 1, 0.1, hi)

	// gui.NewLabel(screen, "Hello World", 0.5, 0.5, 0.1)
	// gl.DepthFunc(gl.LEQUAL)
	for {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		// b.Draw()
		screen.Update()
		glfw.PollEvents()
		w.SwapBuffers()
	}
}
