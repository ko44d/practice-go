package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx1, cancel1 := context.WithCancel(context.Background())
	fmt.Println("キャンセル確認")
	child(ctx1)
	cancel1()
	fmt.Println("キャンセル確認")
	child(ctx1)

	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel2()
	go func() { fmt.Println("別のゴルーチン") }()
	fmt.Println("停止")
	<-ctx2.Done()
	fmt.Println("再起動")

	ctx3, cancel3 := context.WithCancel(context.Background())
	task := make(chan int)
	go func() {
		for {
			select {
			case <-ctx3.Done():
				return
			case i := <-task:
				fmt.Println("get", i)
			default: // defaultが定義されていない場合、いずれかのcase条件が満たされるまで処理がブロックされる
				fmt.Println("キャンセルされていない")
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()
	time.Sleep(time.Second)
	for i := 0; 5 > i; i++ {
		task <- i
	}
	cancel3()

	ctx4 := context.Background()
	fmt.Printf("traceId = %q\n", GetTraceId(ctx4))
	ctx4 = SetTraceId(ctx4, "test-id")
	fmt.Printf("traceId = %q\n", GetTraceId(ctx4))
}
func child(ctx context.Context) {
	if err := ctx.Err(); err != nil {
		fmt.Println("キャンセルされた")
		return
	}

	fmt.Println("キャンセルされていない")
}

type TraceId string

const ZeroTraceId = ""

type traceIdKey struct {
}

func SetTraceId(ctx context.Context, tid TraceId) context.Context {
	return context.WithValue(ctx, traceIdKey{}, tid)
}

func GetTraceId(ctx context.Context) TraceId {
	if v, ok := ctx.Value(traceIdKey{}).(TraceId); ok {
		return v
	}
	return ZeroTraceId
}
