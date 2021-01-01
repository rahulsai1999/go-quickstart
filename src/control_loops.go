package src

import "fmt"

// FlowControl -> flow control explained
func FlowControl() {
	if 5 > 3 {
		fmt.Println("True in if clause")
	} else if 3 > 2 {
		fmt.Println("If else clause")
	} else {
		fmt.Println("Else Clause")
	}

	x := 42
	switch x {
	case 0:
	case 1:
	case 42:
		fmt.Println("Reached 42")
	case 43:
		fmt.Println("Unreached")
	default:
		fmt.Println("Default case is optional")
	}

}

// Loops -> examples for loops in Golang
func Loops() {
	for x := 0; x < 5; x++ {
		fmt.Println("Print 5 times")
	}

	// You can use range to iterate over an array, a slice, a string, a map, or a channel.
	// range returns one (channel) or two values (array, slice, string and map).

	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		// for each pair in the map, print key and value
		fmt.Printf("key=%s, value=%d\n", key, value)
	}
	// If you only need the value, use the underscore as the key
	for _, name := range []string{"Bob", "Bill", "Joe"} {
		fmt.Printf("Hello, %s\n", name)
	}
}
