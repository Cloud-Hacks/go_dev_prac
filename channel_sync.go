package main

import "fmt"

func main() {
	req := make(chan string, 5)
	resp := make(chan string, 5)
	err := make(chan error)

	go func() {
		for {
			select {
			case request := <-req:
				if request == "Plug added" {
					resp <- "Added"
				} else if request == "Plug deleted" {
					resp <- "Deleted"
				} else {
					err <- fmt.Errorf("Error caught")
				}
			}
		}
	}()

	req <- "Plug added"
	req <- "Plug deleted"
	req <- "Plug updated"

	for i := 0; i < 3; i++ {
		select {
		case tmp := <-resp:
			fmt.Println(tmp)
		case er := <-err:
			fmt.Println(er)
		}
	}

	close(req)
	close(resp)
}
