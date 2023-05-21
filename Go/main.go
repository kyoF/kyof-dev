package main

import (
	"fmt"
	"math"
)

func main() {
	// 切り捨て
	fmt.Println(math.Trunc(3.14))
	fmt.Println(math.Trunc(-3.14))
	fmt.Println(math.Trunc(-3.8))
	// 切り下げ
	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Floor(-3.5))
	// 切り上げ
	fmt.Println(math.Ceil(3.5))
	fmt.Println(math.Ceil(-3.5))

	// // 絶対値
	// fmt.Println(math.Abs(100))
	// fmt.Println(math.Abs(-100))
	// // 累乗
	// fmt.Println(math.Pow(0, 2))
	// fmt.Println(math.Pow(2, 2))
	// // 平方根、立法根
	// fmt.Println(math.Sqrt(2))
	// fmt.Println(math.Sqrt(8))
	// // 最大値、最小値
	// fmt.Println(math.Max(1, 3))
	// fmt.Println(math.Min(1, 3))

	// // 円周率
	// fmt.Println(math.Pi)
	// // 2のルート
	// fmt.Println(math.Sqrt2)

	// fmt.Println(math.MaxFloat32)
	// fmt.Println(math.SmallestNonzeroFloat32)
	// fmt.Println(math.MaxFloat64)
	// fmt.Println(math.SmallestNonzeroFloat64)
	// fmt.Println(math.MaxInt64)
	// fmt.Println(math.MinInt64)
}
