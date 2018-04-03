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
func main() {
	// println(gl.LESS)
	runtime.LockOSThread()
	w := gui.InitGlfw()
	gui.InitOpenGL()
	w.SetSizeCallback(windowSizeCallback)
	screen := gui.NewScreen(w)
	screen.InitGui("extras/font/font.png", "extras/font/font.json", "extras/button.png", "")
	guimousebuttoncallback := screen.MouseButtonCallback()
	guicursorposcallback := screen.CursorPosCallback()
	w.SetMouseButtonCallback(guimousebuttoncallback)
	w.SetCursorPosCallback(guicursorposcallback)
	b := gui.NewButton(screen, "hiasfdslksdfajl", hi, 0, 0, 0.5, 1, 0.2)
	// gui.NewLabel(screen, "Hello World", 0.5, 0.5, 0.1)
	b.Text = "hisldsflsfdahjlsd"
	// gl.DepthFunc(gl.LEQUAL)
	for {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		// b.Draw()
		screen.Update()
		glfw.PollEvents()
		w.SwapBuffers()
	}
}
