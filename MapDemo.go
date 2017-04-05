package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("Assigned value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("After delete", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
