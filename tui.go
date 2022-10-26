package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/gookit/color"
	"golang.org/x/term"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func TuiClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func TuiPageCheck() {
	_, height := TuiSize()
	_, row := TuiCursorPos()
	//row := 10
	//fmt.Println("row: ", row, "height: ", height)
	if row > (height - 5) {
		fmt.Println()
		TuiPause()
		TuiClear()
	}
}

func TuiPause() {
	green := color.FgGreen.Render
	width, _ := TuiSize()
	for i := 1; i < (width/2 - 15); i++ {
		fmt.Print(" ")
	}
	fmt.Println(green("press Enter key to continue..."))
	fmt.Scanln()
}

func TuiSize() (int, int) {
	width, height, err := term.GetSize(0)
	if err != nil {
		os.Exit(1)
	}
	return width, height
}

/*
Get the terminal cursor's position
- Set the terminal to raw mode
- Use the ANSI escape sequence to get the cursor position
- Read in the result
- Set the terminal back to cooked mode
*/

func TuiCursorPos() (int, int) {

	// Set the terminal to raw mode (to be undone with `-raw`)
	rawMode := exec.Command("/bin/stty", "raw")
	rawMode.Stdin = os.Stdin
	_ = rawMode.Run()
	rawMode.Wait()

	// same as $ echo -e "\033[6n"
	cmd := exec.Command("echo", fmt.Sprintf("%c[6n", 27))
	randomBytes := &bytes.Buffer{}
	cmd.Stdout = randomBytes

	// Start command asynchronously
	_ = cmd.Start()

	// capture keyboard output from echo command
	reader := bufio.NewReader(os.Stdin)
	cmd.Wait()

	// by printing the command output, we are triggering input
	fmt.Print(randomBytes)
	fmt.Printf("\033[%dA", 1)
	// capture the triggered stdin from the print
	text, _ := reader.ReadSlice('R')

	// Set the terminal back from raw mode to 'cooked'
	rawModeOff := exec.Command("/bin/stty", "-raw")
	rawModeOff.Stdin = os.Stdin
	_ = rawModeOff.Run()
	rawModeOff.Wait()

	// check for the desired output
	col, row := 0, 0
	if strings.Contains(string(text), ";") {
		re := regexp.MustCompile(`\d+;\d+`)
		line := re.FindString(string(text))
		split := strings.Index(line, ";")
		if split > -1 {
			ystr := line[:split]
			xstr := line[split+1:]
			y0, err := strconv.Atoi(ystr)
			if err == nil {
				row = y0
			}
			x0, err := strconv.Atoi(xstr)
			if err == nil {
				col = x0
			}
		} else {
			fmt.Println("Index not found")
			fmt.Println(line)
		}
	} else {
		fmt.Println("it does not work. womp womp.")
	}
	return col, row
}
