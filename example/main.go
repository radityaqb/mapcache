package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/radityaqb/mapcache"
)

func main() {
	// loadTest(900000)
	example()
}

func example() {
	ctx := context.Background()

	arr := []int{100, 200, 300}
	for i, x := range arr {
		mapcache.Save(ctx, "pkg_name", i, x)
	}

	for i, _ := range arr {
		z, _ := mapcache.Load(ctx, "pkg_name", i)
		fmt.Println(z)
	}

}

func loadTest(n int) {
	ctx := context.Background()

	arr := randomArray(n)

	for i, x := range arr {
		go mapcache.Save(ctx, "int", i, x)

		go mapcache.Load(ctx, "int", i)
	}
}

func randomArray(len int) []int {
	a := make([]int, len)
	for i := 0; i <= len-1; i++ {
		a[i] = rand.Intn(len)
	}
	return a
}
