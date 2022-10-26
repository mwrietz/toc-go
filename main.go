package main

import (
	"bufio"
	"fmt"
	"github.com/gookit/color"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// setup output colors
	blue := color.FgBlue.Render
	red := color.FgRed.Render
	green := color.FgGreen.Render
	magenta := color.FgMagenta.Render
	yellow := color.FgYellow.Render
	cyan := color.FgCyan.Render

	// clear screen
	TuiClear()

	// print executable name and version
	path := os.Args[0]
	var s = filepath.Base(path)
	fmt.Println(green(s), ": v0.1.0\n")
    TuiPageCheck()

	// get a list of files in cwd and subdirs 
	filelist := FileListTREE()

	for _, file := range filelist {
        if !strings.HasPrefix(file, ".") {
            fmt.Println(blue(file))
        }
        TuiPageCheck()

		linecount := 0
		var multiline_import = false
		var multiline_type = false

		if strings.HasSuffix(file, ".go") {
			f, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			scanner := bufio.NewScanner(f)

			//var multiline_import = false
			//var multiline_type = false

			for scanner.Scan() {
				linecount += 1
				buffer := scanner.Text()

				if strings.HasPrefix(buffer, "func") {
					lc := fmt.Sprintf("%6d", linecount)
					fmt.Println(red(lc), ":", strings.TrimSuffix(buffer, "{"))
                    TuiPageCheck()
				}

				if strings.HasPrefix(buffer, "package") {
					lc := fmt.Sprintf("%6d", linecount)
					fmt.Println(red(lc), ":", yellow(buffer))
                    TuiPageCheck()
				}

				if strings.HasPrefix(buffer, "import") {
					lc := fmt.Sprintf("%6d", linecount)
					fmt.Println(red(lc), ":", cyan(buffer))
                    TuiPageCheck()
					if strings.Contains(buffer, "(") {
						multiline_import = true
					}
				}
				if multiline_import == true {
					if !strings.Contains(buffer, "(") {
						lc := fmt.Sprintf("%6d", linecount)
						fmt.Println(red(lc), ":", cyan(buffer))
                        TuiPageCheck()
					}
					if strings.Contains(buffer, ")") {
						multiline_import = false
					}
				}

				if strings.HasPrefix(buffer, "type") {
					lc := fmt.Sprintf("%6d", linecount)
					fmt.Println(red(lc), ":", magenta(buffer))
                    TuiPageCheck()
					multiline_type = true
				}
				if multiline_type == true {
					if !strings.Contains(buffer, "{") {
						lc := fmt.Sprintf("%6d", linecount)
						fmt.Println(red(lc), ":", magenta(buffer))
                        TuiPageCheck()
					}
					if strings.Contains(buffer, "}") {
						multiline_type = false
					}
				}
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
            fmt.Println()
            TuiPageCheck()
		}
	}
}

