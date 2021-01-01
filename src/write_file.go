package src

import (
	"fmt"
	"os"
)

func eWriteToFile() {
	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "This is an example to write to file")
	file.Close()
}
