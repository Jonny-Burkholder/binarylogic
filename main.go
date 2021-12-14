package main

import (
	"fmt"
	"time"

	"github.com/Jonny-Burkholder/binarylogic/blogic"
)

func main() {

	a := blogic.NewBnum()
	time.Sleep(1)
	b := blogic.NewBnum()

	fmt.Printf("First Number: %v\n", a.Number)
	fmt.Printf("Second Number: %v\n", b.Number)
	fmt.Printf("Sum: %v\n", blogic.Add(a, b).Number)

}
