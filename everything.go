package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	// aVariables()
	// cTypes()
	// dComplexTypes()
	// eWriteToFile()
	// fFlowControl()
	// gLoops()
	hHTTPRequests()
}

func aVariables() {
	var x int // declaration
	x = 3     // assignment
	y := 4    // short declaration - infers the type,declare and assign

	sum, prod := bFunctions(x, y)
	fmt.Println("Sum:", sum, " Prod:", prod)
}

func bFunctions(x, y int) (sum, prod int) {
	// function parameters are passed like this
	// x, y are the input arguments of type int
	// sum, prod are the output signatures of type int
	return x + y, x * y
}

func bFunctionsNamedReturn(x, y int) (z int) {
	z = x * y
	return // z is implicit as only single parameter is present
}

func cTypes() {
	// datatypes available in Go
	str := "Learn Go!"
	str2 := `A raw string can include
	line breaks.`
	g := 'Σ'
	f := 3.14195
	c := 3 + 4i

	var u uint = 7
	var pi float32 = 22. / 7

	n := byte('\n')

	fmt.Println(str, str2, g, f, c, u, pi, n)
}

func dComplexTypes() {
	//complex datatypes in Go

	//arrays (fixed size)
	var a4 [4]int                    // fixed array
	a5 := [...]int{3, 1, 5, 10, 100} //array initialized with fixed size of five

	//slices (dynamic size)
	s3 := []int{4, 5, 9} // dynamic sized array dec and init
	s4 := make([]int, 4) //assigns four sized slice initialized to 0
	var d2 [][]float64
	bs := []byte("a slice")

	s3 = append(s3, 4, 5, 6)
	s3 = append(s3, []int{7, 8, 9}...)

	fmt.Println(s3)

	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1

	//unused variables are errors in Go
	//underscore
	_, _, _, _, _ = a4, a5, s4, d2, bs

	fmt.Println(a4, a5, s4, d2, bs)
}

func eWriteToFile() {
	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "This is an example to write to file")
	file.Close()
}

func fFlowControl() {
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

func gLoops() {
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

//structure for unmarshalling json
type obj struct {
	Positive float32
	Negative float32
}
type resultBody struct {
	Status      string
	Predictions []obj
}

type requestBody struct {
	Text []string `json:"text"`
}

func hHTTPRequests() {
	url := "http://localhost:5000/model/predict"

	//preparing the object
	reqobj := requestBody{}
	reqobj.Text = append(reqobj.Text, "Kubernetes is a great orchestration service", "Uber is not the best taxi service")

	//json conversion
	data, _ := json.Marshal(reqobj)

	//convert to byte[]
	payload := strings.NewReader(string(data))
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var finobj resultBody
	json.Unmarshal([]byte(body), &finobj)

	fmt.Println(finobj.Predictions)
}
