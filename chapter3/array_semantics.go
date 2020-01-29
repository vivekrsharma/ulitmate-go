package chapter3

import "fmt"

func main() {
	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1]= "Banana"
	fruits[2]="Cherry"
	fruits[3]="Grape"
	fruits[4]="Plum"

	// Pointer semantics for accessing fruits.
	// Since we are not making any copy, the "CHANGED" is reflected in the
	// iteration of the for loop.
	for i:=0;i<len(fruits);i++ {
		fruits[1] = "CHANGED"
		fmt.Printf(" %d [%s]",i, fruits[i])
	}

	fmt.Println("*********************")
	// reset value for next loop.
	fruits[1] = "Banana"

	// Value semantics for accessing fruits.
	// Since we use value semanics, it makes a copy of fruit and the changed value
	// isn't reflected.
	for i,fruit := range fruits {
		fruits[1] = "CHANGED"
		fmt.Printf(" Index[%d]   Value [%s] ValueAddress [%p] IndexAddress [%p] \n",i,fruit,&fruit, &fruits[i])
	}


}
