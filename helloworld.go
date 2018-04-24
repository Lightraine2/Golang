package main
//btw, single line comments look like this
import ("fmt"
        "math/rand")

func add(x float64,y float64) float64 {
    return x+y
}

func main() {
    fmt.Println("Welcome to Go!")
    fmt.Println("A random number from 1-100 is:",rand.Intn(100))
      var num1 float64 = 5.6
      var num2 float64 = 9.6

      fmt.Println(add(num1,num2))
}
