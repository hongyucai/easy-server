package main
import (
	"fmt"
)
type Coordinate struct {
	X, Y float32
	print bool
}

func (coo *Coordinate) GetCoordinate() {
	if coo.print == true{
		fmt.Printf("(%.2f,%.2f)\n", coo.X, coo.Y)
	}
	return
}
func (coo *Coordinate) chanelVule()  {
	//coo.X<- float32(1.00)

}
//值拷贝对象方法
func (coo Coordinate) SetPosition01(a float32,b float32) {
	coo.X = a
	coo.Y = b
	coo.GetCoordinate()
}

//指针变量对象方法
func (coo *Coordinate) SetPosition02(a float32,b float32) {

	coo.X = a
	coo.Y = b
	coo.GetCoordinate()
}
func main(){
	p0 := Coordinate{1, 2,true}
	fmt.Print("SetPosition01调用前:")
	p0.GetCoordinate()
	fmt.Print("SetPosition01调用后:")
	p0.print = false
	p0.SetPosition01(0, 0)
	p0.GetCoordinate()
	fmt.Print("SetPosition01函数内调用:")
	p0.print = true

	fmt.Print("SetPosition02调用前:")
	p0.GetCoordinate()
	fmt.Print("SetPosition02调用后:")
	p0.SetPosition02(0, 0)
}