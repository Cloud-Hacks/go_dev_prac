package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func fetchThirdPartyDataAnything() (int, error) {
	time.Sleep(50 * time.Second)

	return 777, nil
}

func fetchDataAny(ctx context.Context, Uid int) (int, error) {
	val, err := fetchThirdPartyDataAnything()

	if err != nil {
		return 0, err
	}

	return val, nil

}

func main() {
	strt := time.Now()
	ctx := context.Background()
	val, err := fetchDataAny(ctx, 10)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("res %d", val)
	fmt.Println(time.Since(strt))
}
