package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// root := context.Background()

	// fmt.Println(time.Now())
	// fmt.Println(time.Now().Add(3 * time.Second))

	// child, _ := context.WithDeadline(root, time.Now().Add(10  * time.Second))

	// for {

	// 	if child.Err() == context.DeadlineExceeded {
	// 		fmt.Println("finished")z
	// 		break
	// 	}
	// }


	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2  * time.Second))

	ctx = context.WithValue(ctx, "x", "hello")

	fn(ctx)
	fn(ctx)
	fn(ctx)
	cancel()

	fn(ctx)
	fn(ctx)
	fn(ctx)
	fn(ctx)
}

func fn(ctx context.Context) {

	if ctx.Err() != context.Canceled {
		fmt.Println(ctx.Value("x"))
	}
}
