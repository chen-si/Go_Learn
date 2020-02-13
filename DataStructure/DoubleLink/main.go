package main

import "fmt"

//定义一个HeroNode
type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode
	next     *HeroNode
}

//编写第一种插入方法，在最后插入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//1、先找到链表的最后节点
	//2、传建一个辅助节点
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next //让temp不断指向下一个节点
	}
	temp.next = newHeroNode //将newHeroNode加入到链表的结尾
	newHeroNode.pre = temp
}

//根据英雄的编号从小到大插入
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//1、先找到链表的适当的
	//2、创建一个辅助节点

	temp := head
	flag := true

	for {
		if temp.next == nil { //到链表的最后了
			break
		} else if temp.next.no > newHeroNode.no {
			//说明newHeroNode就应该插入到此后
			break
		} else if temp.next.no == newHeroNode.no {
			flag = false
			break
			//说明我们的链表中已经有这个人
		}
		temp = temp.next
	}
	if !flag {
		fmt.Printf("对不起，已经存在%d号英雄\n", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		newHeroNode.pre = temp
		if temp.next != nil {
			temp.next.pre = newHeroNode
		}
		temp.next = newHeroNode
	}
}

//显示链表的所有节点信息
func ListHeroLink(head *HeroNode) {
	//辅助节点
	temp := head

	//先判断链表是否为空
	if temp.next == nil {
		fmt.Println("空链表...")
		return
	}

	for {
		if temp.next == nil {
			fmt.Println("输出完毕")
			break
		}
		fmt.Printf("[%d\t%s\t%s\t]\n", temp.next.no, temp.next.name,
			temp.next.nickname)
		temp = temp.next
	}
}

//逆序打印
func ListHeroLink2(head *HeroNode) {
	//辅助节点
	temp := head
	//先判断链表是否为空
	if temp.next == nil {
		fmt.Println("空链表...")
		return
	}

	for temp.next != nil {
		temp = temp.next
	}

	for {
		if temp == head {
			fmt.Println("输出完毕")
			break
		}
		fmt.Printf("[%d\t%s\t%s\t]\n", temp.no, temp.name,
			temp.nickname)
		temp = temp.pre
	}
}

func DelHeroNode(head *HeroNode, no int) {
	temp := head
	for {
		if temp.no == no {
			break
		} else if temp.next == nil {
			fmt.Println("未找到此项")
			return
		}
		temp = temp.next
	}
	temp.pre.next = temp.next
	if temp.next != nil {
		temp.next.pre = temp.pre
	}
	temp.next = nil
	temp.pre = nil
}

func main() {
	//1、先创建一个头结点
	head := &HeroNode{}

	//2、创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}

	hero4 := &HeroNode{
		no:       4,
		name:     "吴用",
		nickname: "智多星",
	}

	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero4)
	InsertHeroNode2(head, hero2)
	InsertHeroNode2(head, hero3)

	ListHeroLink(head)
	ListHeroLink2(head)

	DelHeroNode(head, 4)

	fmt.Println("-----------------------------")
	ListHeroLink2(head)
	ListHeroLink(head)

	// DelHeroNode(head, 1)
	// ListHeroLink(head)
}
