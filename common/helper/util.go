package helper

import (
	"fmt"
	"math"
	"strconv"
	"github.com/shopspring/decimal"
)

func Decimal2(d float64) float64 {
	f, err := strconv.ParseFloat(fmt.Sprintf("%.2f", d), 64)
	if err==nil {
		return f
	}
	return 0
}
func Decimal3(d float64) float64 {
	f, err := strconv.ParseFloat(fmt.Sprintf("%.3f", d), 64)
	if err==nil {
		return f
	}
	return 0
}
func Decimal4(d float64) float64 {
	f, err := strconv.ParseFloat(fmt.Sprintf("%.4f", d), 64)
	if err==nil {
		return f
	}
	return 0
}
func Decimal9(d float64) float64 {
	f, err := strconv.ParseFloat(fmt.Sprintf("%.9f", d), 64)
	if err==nil {
		return f
	}
	return 0
}
func Decimal11(d float64) float64 {
	f, err := strconv.ParseFloat(fmt.Sprintf("%.11f", d), 64)
	if err==nil {
		return f
	}
	return 0
}
func Round(x float64) int64{ //四舍五入 实现原理先+0.5，然后向下取整！ 官方的math 包中提供了取整的方法，向上取整math.Ceil() ，向下取整math.Floor()
	return int64(math.Floor(x + 0/5))
}


// 加法
func Add(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Add(d2)
}

// 减法
func Sub(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Sub(d2)
}

// 乘法
func Mul(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Mul(d2)
}

// 除法
func Div(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Div(d2)
}

func DecimalNewFromString(d string) decimal.Decimal {
	v, _ := decimal.NewFromString(d)
	return v
}
// int
func Int64(d decimal.Decimal) int64{
	return d.IntPart()
}

// float
func Float64(d decimal.Decimal) float64{
	f, exact := d.Float64()
	if !exact{
		return Decimal11(f)
	}
	f1 := d.IntPart()
	if f1>0{
		return float64(f1)
	}
	return 0
}