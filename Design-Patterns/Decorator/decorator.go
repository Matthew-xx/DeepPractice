package Decorator

import (
	"log"
	"math"
	"time"
)

//如果要给几十上百个函数加日志，

//定义函数pi的签名，函数传入int参数返回float64
type piFunc func(int) float64

//WrapLogger（自定义日志函数）装饰函数pi（该函数类型）
func WrapLogger(fun piFunc,looger *log.Logger) piFunc {
	//return返回一个匿名函数（该函数的结构就是pifunc结构），里面的逻辑自定义
	return func(n int) float64 {
		fn := func(n int) (result float64) {
			defer func(t time.Time) {
				looger.Printf("took=%v,n=%v,result=%v",time.Since(t),n,result)
			}(time.Now())
			return fun(n)
		}
		return fn(n)
	}
}
func Pi(n int) (result float64) {
	ch := make(chan float64)
	for k:=0;k<n;k++ {
		go func(ch chan float64,k float64) {
			ch <- 4*math.Pow(-1,k)/(2*k+1)
		}(ch, float64(k))
	}
	for i:=0;i<n;i++ {
		result += <- ch
	}
	return
}
