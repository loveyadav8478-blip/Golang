package main

import "fmt"

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func getLan() (string, string, int) {
	return "Java", "Golang", 2
}

func processIt(fn func(a int) int) int {
	return fn(1)
}

func returnFun() func(a int) int {
	fn := func(a int) int {
		return 122
	}
	return fn
}

func main() {
	// test("Lane,", " happy birthday!")
	// test("Zuck,", " hope that Metaverse thing works out")
	// test("Go", " is fantastic")

	lang1, lang2, v := getLan()

	fmt.Printf("%s %s %v\n", lang1, lang2, v)

	fn := func(a int) int {
		return 2
	}
	fmt.Println(processIt(fn))

	fnn := returnFun()
	fmt.Println(fnn(1))
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}
