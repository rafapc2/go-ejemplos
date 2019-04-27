package main

import "fmt"

func add(a int,b int ) int{
     return a +b
}

func f(mensaje string, count int ){
  for indice := 0; indice < count; indice++ {
    fmt.Println(mensaje, " = ", indice)
  }
}

func main()  {
  fmt.Println("Hello Go")

  var  n1 = 50
  var  n2 = 90
  resultado := add(n1 , n2)
  fmt.Println(n1 , " + " , n2 , " = ", resultado)


  go f("paso2",n2)
  go f("paso1",n1)
  fmt.Scanln()
  fmt.Println("fin")

}
