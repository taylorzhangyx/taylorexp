package main

import "fmt"

type L  struct {
  N int
}

type LL []L

func main(){
t := LL{}

for i:=0; i<10; i++ {
  t = append(t, L{i})

}
fmt.Println(t)
}
