package gui

type Entry struct {
	Text string
	X    int
	Y    int
}

func NewEntry(screen *Screen, text string, x, y int) *Entry {
	e := Entry{}
	e.Text = text
	e.X = x
	e.Y = y
	screen.entrys = append(screen.entrys, &e)
	return &e
}
