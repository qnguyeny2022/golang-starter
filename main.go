package main

import (
"fmt"
"os"
"strconv"
)

func main() {
input := os.Args
if len(input) == 2 {
input := os.Args[1]
num, err := strconv.Atoi(input)
if num > 10 || num < 0 || err != nil {
fmt.Println("ERROR!")
} else {
for i := 0; i < num; i++ {
fmt.Println("friendly-trout")
}
}
} else {
fmt.Println("ERROR!")
}
}
