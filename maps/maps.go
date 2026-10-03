package main

import (
	"fmt"

	// "golang.org/x/tools/go/callgraph/cha"
)

func main() {
	m := make(map[string]string)
	m["Love"] = "Yadav"
	m["Love"] = "YDV"

	fmt.Println(m["Love"])
	
	

	mp := make(map[string]int)

	arr := []string{"Hello","My","Name","Is","Charlie","Charlie","Charlie"}

	for i:=0; i<len(arr); i++{
		mp[arr[i]]++;
	}

	s := "1a2ddd33322dvv"
	charCount := make(map[rune]int)
	for _, c := range s{
		charCount[c]++
		
	}
	

	for k,v := range charCount{
		fmt.Printf("%c: %d\n",k,v)
	}
	println()



	// clear(mp)
	// delete(mp,"Love")
	// fmt.Println(charCount)
	fmt.Println(mp)
}