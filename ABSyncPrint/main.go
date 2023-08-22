package ABSyncPrint

import "fmt"

func main() {
	var (
		aCh = make(chan struct{})
		bCh = make(chan struct{})
	)

	for i:= 0 ;i<50 ;i++ {
		go func ()  {
			
			fmt.Println("A")
		}
	}

	for i:=0 ;i<50 ;i++ {
		go func ()  {
			fmt.Println("B")
		}
	}

	SELECT {}
}
