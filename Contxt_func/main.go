package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Res struct {
	v  int
	er error
}

// context allows you to maintain the consistency among the fetch requests to third party
func fetchThirdPartyDataAnything() (int, error) {
	ctx := context.WithValue(context.Background(), "todo", 5)
	t := ctx.Value("todo")
	fmt.Println(t)
	time.Sleep(50 * time.Millisecond)

	return 777, nil
}

func fetchDataAny(ctx context.Context, Uid int) (int, error) {
	ctx, canc := context.WithTimeout(ctx, time.Millisecond*200)
	defer canc()

	resp := make(chan Res)

	go func() {
		val, err := fetchThirdPartyDataAnything()
		resp <- Res{
			v:  val,
			er: err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("error fetching from 3rd party")
		case res := <-resp:
			// fmt.Println(resp)
			return res.v, res.er
		case <-resp:
			// fmt.Println(resp, <-resp)
			return Uid, nil
		}
	}

}

func main() {
	strt := time.Now()
	ctx := context.Background()
	val, err := fetchDataAny(ctx, 10)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("res %d \n", val)

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("time exceeded")
		default:
			fmt.Println("channel created")
		}
	}()

	time.Sleep(2 * time.Second)

	fmt.Println(time.Since(strt))
}
