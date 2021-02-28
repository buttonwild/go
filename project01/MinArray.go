package main

import "fmt"

func main() {
	a := []int{1, 2, 1}
	fmt.Println(minArray(a))
}

func minArray(numbers []int) int {
	// 异常一{1，1}
	// 修改循环判断条件
	// 判断的是low!=hight
	// 但是没有给如果不满足循环后的输出
	// 异常二{3，1}
	// 更换区间错误
	// 异常三{3，1}
	// mid==low
	// 说明low和hight相邻
	// low=hight
	// 异常四
	// 将low=hight放在循环第一节会导致全顺序排列的判断失败，需要将其放入子判断中
	// 异常五
	// 当low>hight&&mid==hight
	// 知道else的作用，就是当if修改了数据，之后的判断又刚好用到该数据，就死亡了
	// 使用continue跳出本次循环
	// 异常六{1.2，1}
	// 和四是同一个问题，只需要low++，想太多
	var low, hight, mid int
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	low = 0
	hight = len(numbers) - 1
	mid = low + (hight-low)/2
	for low != hight {
		if numbers[low] < numbers[hight] {
			return numbers[low]
		}
		if numbers[low] == numbers[hight] {
			low++
		}
		if numbers[low] > numbers[hight] {
			if numbers[mid] <= numbers[hight] {
				hight = mid
				mid = low + (hight-low)/2
			}
			if numbers[mid] > numbers[hight] {
				low = mid
				mid = low + (hight-low)/2
			}
			if mid == low {
				low = hight
			}
		}
	}
	if low == hight {
		return numbers[low]
	}
	return -1
}
