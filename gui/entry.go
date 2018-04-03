package gui

type Entry struct {
	Text string
	X    int
	Y    int
	Bg   string
}

func NewEntry(screen *Screen, text string, x, y int, bg string) *Entry {
	e := Entry{}
	if bg == "" {
		e.Bg = "#ffffff"
	} else {
		e.Bg = bg
	}
	e.Text = text
	e.X = x
	e.Y = y
	screen.entrys = append(screen.entrys, &e)
	return &e
}
