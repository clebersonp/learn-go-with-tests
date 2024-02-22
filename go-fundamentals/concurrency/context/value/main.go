package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "language", "Go")
	foo(ctx, "language")

	ctx = context.WithValue(ctx, "dog", "Gaston")
	foo(ctx, "dog")

	foo(ctx, "color")

	fmt.Printf("%T\n", ctx)
	fmt.Printf("%v\n", ctx)

}

func foo(ctx context.Context, key string) {
	if v := ctx.Value(key); v != nil {
		fmt.Printf("found key %q with value %q\n", key, v)
		return
	}
	fmt.Printf("Value not found for key %q\n", key)
}
