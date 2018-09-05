package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/radityaqb/mapcache"
)

func main() {
	loadTest(900000)
	// example()
}

func example() {
	ctx := context.Background()

	mapcache.InitTTL(ctx, "pkg_name", 3)

	arr := []int{100, 200, 300}
	for i, x := range arr {
		mapcache.Save(ctx, "pkg_name", i, x)
	}

	for i, _ := range arr {
		z, _ := mapcache.Load(ctx, "pkg_name", i)
		fmt.Println(z)
	}

	mapcache.Delete(ctx, "pkg_name")
}

func loadTest(n int) {
	ctx := context.Background()

	mapcache.InitTTL(ctx, "int", 3)

	arr := randomArray(n)

	for i, x := range arr {
		go mapcache.Save(ctx, "int", i, x)

		go mapcache.Load(ctx, "int", i)

		if i%10 == 0 {
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func randomArray(len int) []int {
	a := make([]int, len)
	for i := 0; i <= len-1; i++ {
		a[i] = rand.Intn(len)
	}
	return a
}
