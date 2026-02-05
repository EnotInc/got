package main

import "fmt"

type board []byte

const (
	separator = "---+---+---\n"
	prompt    = "\n> "
	cross     = 'x'
	zero      = 'o'

	clearView    = "\033[2J"
	clearHistory = "\033[3J"
	moveToStart  = "\033[0H"

	reset  = "\033[0m"
	red    = "\033[31m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	gray   = "\033[90m"
)

type ttt struct {
	message   string
	data      []byte
	represent string
	item      byte
	place     int
	turn      int
}

func main() {
	t := ttt{
		data:  []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'},
		place: 1,
	}
	t.Draw()

	for t.turn < 9 {
		if t.isWin() {
			t.Draw()
			break
		}
		success := t.makeTurn()
		if success {
			t.turn += 1
			t.message = ""
		}
		t.Draw()
	}
}

func paint(ch byte) string {
	var color = ""
	switch ch {
	case 'x':
		color = purple
	case 'o':
		color = blue
	case '\\', '/', '-', '|':
		color = yellow
	default:
		color = gray
	}
	return fmt.Sprintf("%s%c%s", color, ch, reset)
}

func (t *ttt) isWin() bool {
	if t.data[0] == t.data[4] && t.data[4] == t.data[8] {
		t.data[0] = '\\'
		t.data[4] = '\\'
		t.data[8] = '\\'
		return true
	} else if t.data[2] == t.data[4] && t.data[4] == t.data[6] {
		t.data[2] = '/'
		t.data[4] = '/'
		t.data[6] = '/'
		return true
	}
	for i := range 3 {
		if t.data[i*3] == t.data[i*3+1] && t.data[i*3+1] == t.data[i*3+2] {
			t.data[i*3] = '-'
			t.data[i*3+1] = '-'
			t.data[i*3+2] = '-'
			return true
		} else if t.data[i] == t.data[i+1*3] && t.data[i+1*3] == t.data[i+2*3] {
			t.data[i] = '|'
			t.data[i+1*3] = '|'
			t.data[i+2*3] = '|'
			return true
		}
	}
	return false
}

func (t *ttt) makeTurn() bool {
	_, err := fmt.Scan(&t.place)
	if err != nil || t.place > 9 || t.place < 1 {
		//panic(fmt.Sprintf("I just gonna break the game, coz im too lazy to handle an error. And yeah, here is an error: %s", err))
		t.message = "type 1-9, not this noncense"
		return false
	}
	switch t.turn%2 == 0 {
	case true:
		t.item = cross
	case false:
		t.item = zero
	}
	if t.data[t.place-1] == 'x' || t.data[t.place-1] == 'o' {
		t.message = "this place is taken, choose another one"
		return false
	} else {
		t.data[t.place-1] = t.item
	}
	return true
}

func (t *ttt) Draw() {
	fmt.Print(clearView, clearHistory, moveToStart)
	t.represent = red + t.message + reset + "\n"

	for i := range 3 {
		t.represent += fmt.Sprintf(" %s | %s | %s \n", paint(t.data[3*i]), paint(t.data[3*i+1]), paint(t.data[3*i+2]))
		if i != 2 {
			t.represent += fmt.Sprint(separator)
		}
	}
	t.represent += fmt.Sprint(prompt)
	fmt.Print(t.represent)
}
