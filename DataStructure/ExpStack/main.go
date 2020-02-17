package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	MaxTop int //表示栈的最大可以存放数的个数
	Top    int //表示栈顶
	arr    [20]int
}

func (this *Stack) Push(val int) {
	//先判断栈是否满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("栈满！")
		return
	}
	this.Top++
	this.arr[this.Top] = val
	return
}

func (this *Stack) List() {
	if this.Top == -1 {
		fmt.Println("栈空!")
		return
	}
	for i := 0; i <= this.Top; i++ {
		fmt.Printf(" %d ->", this.arr[i])
	}
	fmt.Println()
}

func (this *Stack) Pop() (val int) {
	if this.Top == -1 {
		fmt.Println("栈空!")
		return
	}
	val = this.arr[this.Top]
	this.Top--
	return
}

//判断一个字符是不是一个运算符的方法
func (this *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

//运算的方法
func (this *Stack) cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num1
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误..")
	}
	return res
}

//编写一个方法 返回某个运算符的优先级
//* / => 1  + - => 0
func (this *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func main() {
	//数栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	//符号栈
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}

	exp := "30+30*6-4"
	//定义一个index帮助扫描表达式
	index := 0
	//配合运算，定义需要的变量
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := ""

	for {
		ch := exp[index : index+1]
		//判断扫描出来的是数字还是运算符
		temp := int([]byte(ch)[0])
		if operStack.IsOper(temp) { //说明是符号
			//如果operStack是一个空栈，直接入栈
			if operStack.Top == -1 { //空栈
				operStack.Push(temp)
			} else {
				//判断优先级 计算
				if operStack.Priority(operStack.arr[operStack.Top]) >=
					operStack.Priority(temp) {
					num1 = numStack.Pop()
					num2 = numStack.Pop()
					oper = operStack.Pop()
					result = operStack.cal(num1, num2, oper)
					//将计算结果入栈
					numStack.Push(result)
					//当前的符号压入符号栈
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}
		} else { //说明是数
			//出路多为数的思路
			//1、 定义一个keepnum string 做拼接
			//2、 每次要向index的前面字符测试一下，看看是不是数
			keepNum += ch
			//如果已经到表达式最后，直接将keepNum
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				//向index的后一位探一位 测试一下
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				}
			}
		}

		//继续是扫描
		//先判断是否结尾
		if index+1 == len(exp) {
			break
		}
		index++
	}
	for {
		if operStack.Top == -1 {
			break //退出条件
		}
		num1 = numStack.Pop()
		num2 = numStack.Pop()
		oper = operStack.Pop()
		result = operStack.cal(num1, num2, oper)
		//将计算结果入栈
		numStack.Push(result)
	}

	//如果表达式没问题 结果就是numStack最后一个数
	res := numStack.Pop()
	fmt.Printf("表达式 %s = %v\n", exp, res)
}
