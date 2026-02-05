package main

import "fmt"

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
	board   []byte
	message string
	item    byte
	place   int
	turn    int
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
	if t.board[0] == t.board[4] && t.board[4] == t.board[8] {
		t.board[0] = '\\'
		t.board[4] = '\\'
		t.board[8] = '\\'
		return true
	} else if t.board[2] == t.board[4] && t.board[4] == t.board[6] {
		t.board[2] = '/'
		t.board[4] = '/'
		t.board[6] = '/'
		return true
	}
	for i := range 3 {
		if t.board[i*3] == t.board[i*3+1] && t.board[i*3+1] == t.board[i*3+2] {
			t.board[i*3] = '-'
			t.board[i*3+1] = '-'
			t.board[i*3+2] = '-'
			return true
		} else if t.board[i] == t.board[i+1*3] && t.board[i+1*3] == t.board[i+2*3] {
			t.board[i] = '|'
			t.board[i+1*3] = '|'
			t.board[i+2*3] = '|'
			return true
		}
	}
	return false
}

func (t *ttt) makeTurn() bool {
	_, err := fmt.Scan(&t.place)
	if err != nil || t.place > 9 || t.place < 1 {
		t.message = "type 1-9, not this noncense"
		return false
	}
	switch t.turn%2 == 0 {
	case true:
		t.item = cross
	case false:
		t.item = zero
	}
	if t.board[t.place-1] == 'x' || t.board[t.place-1] == 'o' {
		t.message = "this place is taken, choose another one"
		return false
	} else {
		t.board[t.place-1] = t.item
	}
	return true
}

func (t *ttt) Draw() {
	fmt.Print(clearView, clearHistory, moveToStart)
	represent := red + t.message + reset + "\n"

	for i := range 3 {
		represent += fmt.Sprintf(" %s | %s | %s \n", paint(t.board[3*i]), paint(t.board[3*i+1]), paint(t.board[3*i+2]))
		if i != 2 {
			represent += fmt.Sprint(separator)
		}
	}
	represent += fmt.Sprint(prompt)
	fmt.Print(represent)
}

func main() {
	t := ttt{
		board: []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'},
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
