package src

import (
	"fmt"
	"os"
)

// WriteToFile -> write to file from golang
func WriteToFile() {
	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "This is an example to write to file")
	file.Close()
}
