package main 
import "fmt"

type person struct{
	name string
	age int
}

func main(){

	fmt.Println("\n", person{"sheikh",25})
	fmt.Println(person{name:"hapheej",age:25})
	fmt.Println(&person{name:"hapheej",age:25})

	human := person{"sheikh",25}
	fmt.Println(&human)

	man := &human
	fmt.Println(man)

}