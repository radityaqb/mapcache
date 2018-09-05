package mapcache

import (
	"context"
	"testing"
)

func Benchmark(b *testing.B) {
	ctx := context.Background()

	InitTTL(ctx, "test", 1)

	for n := 0; n < b.N; n++ {
		Save(ctx, "test", n, n)
	}

	for n := 0; n < b.N; n++ {
		Load(ctx, "test", n)
	}
}
