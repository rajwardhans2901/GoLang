package main
import "fmt"

func main() {

	i:= 0       				//Initialize
	for i < 3 {					//condition
		fmt.Println("i: ",i)
		i++						//increment
	}

	for j:=3; j> 0; j-- {		//reverse loop with all conditions and decrements
		fmt.Println("j: ",j)
	}

	for i:= range 4 {			//use of range
		fmt.Println("range", i)
	}

	k:= 0
	for {						//infinite loop 
		fmt.Println(k)
		k++
		if k == 1 {
			break				//break condition to stop the loop
		}
	} 
}