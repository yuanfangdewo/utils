package array

import (
	"fmt"
	"utils/colorOutput"
)

/**
 * test - 计算两个字符串数组差集
 */
func DifferenceTest(){
	//round 1
	t1 := []string{"a", "b", "c"}
	t2 := []string{"a", "b"}
	r1 := []string{"c"}

	res := Difference(t1, t2)
	if ArrayEqual(res, r1) != true{
		panic("round 1 fail: expect ['c'] \n")
	}
	colorOutput.Green(fmt.Sprintf("%v\n", r1))

	//round 2
	t1 = []string{"a", "b", "c"}
	t2 = []string{"a"}
	r1 = []string{"c", "b"}

	res = Difference(t1, t2)
	if ArrayEqual(res, r1) != true{
		panic(fmt.Sprint("fail: expect %v , got %v \n", r1, res))
	}
	colorOutput.Green(fmt.Sprintf("%v\n", r1))

	//round 2
	t1 = []string{"a", "b", "c"}
	t2 = []string{}
	r1 = []string{"a", "c", "b"}

	res = Difference(t1, t2)
	if ArrayEqual(res, r1) != true{
		panic(fmt.Sprint("fail: expect %v , got %v \n", r1, res))
	}
	colorOutput.Green(fmt.Sprintf("%v\n", r1))

	//round 3
	t1 = []string{"a", "b", "c"}
	t2 = []string{"d", "e"}
	r1 = []string{"a", "c", "b", "d", "e"}

	res = Difference(t1, t2)
	if ArrayEqual(res, r1) != true{
		panic(fmt.Sprint("fail: expect %v , got %v \n", r1, res))
	}
	colorOutput.Green(fmt.Sprintf("%v\n", r1))

	//round 4
	t1 = []string{}
	t2 = []string{}
	r1 = []string{}

	res = Difference(t1, t2)
	if ArrayEqual(res, r1) != true{
		panic(fmt.Sprint("fail: expect %v , got %v \n", r1, res))
	}
	colorOutput.Green(fmt.Sprintf("%v\n", r1))

}
