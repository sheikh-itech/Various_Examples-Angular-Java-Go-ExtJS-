package main

import "fmt"

func main() {
   var a = "sheikh" 
   var b = "hafeez"
   fmt.Println("\nbefore swap : "+a+"---"+b)
   x, y :=  swap(a,b)
   fmt.Println("\nafter swap : "+x+"---"+y+"\n")
}

func swap(name1, name2 string) (string, string){
   
      return name2,name1
}