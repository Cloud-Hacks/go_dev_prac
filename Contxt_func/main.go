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
	time.Sleep(5 * time.Millisecond)

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
			return 0, nil
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
	fmt.Println(time.Since(strt))
}
