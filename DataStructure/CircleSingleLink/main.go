package main

import "fmt"

//定义猫的结构体节点
type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, catNode *CatNode) {
	//判断是不是第一只猫
	if head.next == nil {
		head.no = catNode.no
		head.name = catNode.name
		head.next = head
		fmt.Println(catNode, "已经加入了循环链表")
		return
	}
	//先定义一个辅助节点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入到链表中
	temp.next = catNode
	catNode.next = head
	fmt.Println(catNode, "已经加入了循环链表")
}

func ListCat(head *CatNode) {
	//辅助节点
	temp := head

	//先判断链表是否为空
	if temp.next == nil {
		fmt.Println("空链表...")
		return
	}

	for {
		fmt.Printf("[%d\t%s\t]\n", temp.no, temp.name)
		if temp.next == head {
			fmt.Println("输出完毕")
			break
		}
		temp = temp.next
	}
}

func DelCat(head *CatNode, no int) *CatNode {
	temp := head
	helper := head

	//空链表
	if temp.next == nil {
		fmt.Println("空链表...")
		return head
	}

	//如果只有一个节点
	if temp.next == head {
		temp.next = nil
		return head
	}

	//如果有两个以上的节点
	//让helper指向最后一个节点
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	flag := true
	for {
		if temp.next == head { //已经到最后一个了 最后一个没比较
			break
		}
		if temp.no == no {
			if temp == head { //说明删除头结点
				head = head.next
			}
			helper.next = temp.next
			flag = false
		}
		temp = temp.next
		helper = helper.next
	}
	if flag {
		if temp.no == no {
			helper.next = temp.next
		}
	}

	return head

}

func main() {
	//在这里我们初始化一个猫的头结点
	head := &CatNode{}

	//创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "Tom",
	}

	cat2 := &CatNode{
		no:   2,
		name: "jerry",
	}

	cat3 := &CatNode{
		no:   3,
		name: "nuum",
	}

	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)

	head = DelCat(head, 1)

	ListCat(head)
}
