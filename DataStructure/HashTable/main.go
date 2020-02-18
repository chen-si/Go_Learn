package main

import (
	"fmt"
	"os"
)

//定义emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (this *Emp) ShowMe() {
	fmt.Printf("链表%d 找到雇员ID = %d 名字 %s \n", this.Id%7, this.Id, this.Name)
}

//定义EmpLink
//EmpLink不带表头 即第一个节点就存放信息
type EmpLink struct {
	Head *Emp
}

//添加员工的方法
func (this *EmpLink) Insert(emp *Emp) {
	cur := this.Head   //这是辅助指针
	var pre *Emp = nil //这也是一个辅助指针
	//如果当前的EmpLink就是一个空链表
	if cur == nil {
		this.Head = emp
		return
	}
	//如果不是空链表 给emp找到对应的位置插入
	//思路： 让cur和emp比较 pre保持在cur前面
	for {
		if cur != nil {
			if cur.Id >= emp.Id {
				//找到位置
				break
			}
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}

	//退出时，我们看下是否将emp添加到链表最后
	pre.Next = emp
	emp.Next = cur
}

//显示当前链表的信息
func (this *EmpLink) ShowLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表%d 没有雇员\n", no)
		return
	}
	//遍历当前链表并显示数据
	cur := this.Head
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员ID = %d 名字 %s ->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

//根据id查找对应的雇员
func (this *EmpLink) FindById(id int) *Emp {
	cur := this.Head
	for {
		if cur == nil {
			return nil
		}
		if cur.Id == id {
			return cur
		}
		cur = cur.Next
	}
}

//定义HashTable 一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

//给HashTable添加一个 Insert 方法
func (this *HashTable) Insert(emp *Emp) {
	//使用散列函数 确定将该雇员添加到那个链表
	linkno := this.HashFun(emp.Id)
	//使用对应的链表添加
	this.LinkArr[linkno].Insert(emp)
}

//编写一个方法显示HashTable的所有雇员
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

//编写一个用于散列方法
func (this *HashTable) HashFun(id int) int {
	return id % 7
}

//编写一个方法完成查找
func (this *HashTable) FindById(id int) *Emp {
	linkno := this.HashFun(id)
	return this.LinkArr[linkno].FindById(id)
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("===============雇员系统===========")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show  表示显示雇员")
		fmt.Println("find  表示查找雇员")
		fmt.Println("exit  表示退出系统")
		fmt.Println("请输入选择：")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("请输入雇员id:")
			fmt.Scanln(&id)
			fmt.Println("请输入雇员名字:")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Println("请输入id号：")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Println(id, "不存在")
			} else {
				//编写一个方法显示雇员信息
				emp.ShowMe()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}

	}
}
