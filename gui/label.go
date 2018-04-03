package gui

type Label struct {
	Text string
	X    float64
	Y    float64
	Size float64
}

func NewLabel(screen *Screen, text string, x, y, size float64) *Label {
	l := Label{}
	l.Text = text
	l.X = x
	l.Y = y
	l.Size = size
	screen.labels = append(screen.labels, &l)
	return &l
}
