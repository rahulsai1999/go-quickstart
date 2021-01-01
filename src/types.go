package src

import "fmt"

//Types -> Generic Types in Golang
func Types() {
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

//ComplexTypes -> Complex Types in Golang
func ComplexTypes() {
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
