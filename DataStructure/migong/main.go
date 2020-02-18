package main

import "fmt"

//编写一个函数，完成小鼠找路
//Map *[8][7]int 保证是同一个地图
//i j表示对地图的哪一个点进行测试
func SetWay(Map *[8][7]int, i int, j int) bool {
	//分析出什么情况下找到了出路
	if Map[6][5] == 2 {
		return true
	} else {
		//说明要继续找
		if Map[i][j] == 0 { //如果这个点是可以探测的
			//假设这个点是可以通的 但是需要探测
			Map[i][j] = 2
			if SetWay(Map, i+1, j) { //下
				return true
			} else if SetWay(Map, i, j+1) { //右
				return true
			} else if SetWay(Map, i-1, j) { //上
				return true
			} else if SetWay(Map, i, j-1) { //左
				return true
			} else {
				Map[i][j] = 3
				return false
			}
		} else { //说明这点不能探测是墙
			return false
		}
	}
}

func main() {
	//规则
	//如果元素的值为1 就是墙
	//如果元素的值为0 就是未探测的路径
	//如果元素的值为2 是通路
	//如果元素的值为3 是走不通的路
	var Map [8][7]int

	//先把地图的 最上最下设置为1
	for i := 0; i < 7; i++ {
		Map[0][i] = 1
		Map[7][i] = 1
	}
	//先把地图的 最左最右设置为1
	for i := 0; i < 8; i++ {
		Map[i][0] = 1
		Map[i][6] = 1
	}
	Map[3][1] = 1
	Map[3][2] = 1
	Map[3][4] = 1
	Map[4][4] = 1
	Map[5][4] = 1
	Map[6][4] = 1

	//输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(Map[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println()

	//测试
	SetWay(&Map, 1, 1)

	//探测完毕的地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(Map[i][j], " ")
		}
		fmt.Println()
	}

}
