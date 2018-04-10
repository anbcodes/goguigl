package gui

type Label struct {
	Text   string
	X      float64
	Y      float64
	Size   float64
	Screen *Screen
	index  int
}

func (l *Label) Remove() {
	l.Screen.labels[len(l.Screen.labels)-1], l.Screen.labels[l.index] = l.Screen.labels[l.index], l.Screen.labels[len(l.Screen.labels)-1]
	l.Screen.labels = l.Screen.labels[:len(l.Screen.labels)-1]
}
func NewLabel(screen *Screen, text string, x, y, size float64) *Label {
	l := Label{}
	l.Text = text
	l.X = x
	l.Y = y
	l.Size = size
	l.Screen = screen
	l.index = len(screen.labels) - 1
	screen.labels = append(screen.labels, &l)
	return &l
}
