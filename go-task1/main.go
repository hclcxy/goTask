package main

import (
	"fmt"
	"sort"
)

func main() {

	// 1.   只出现一次的数字：
	nums := []int{5, 1, 2, 1, 2}                            // 示例输入
	fmt.Println("1.", nums, "只出现一次数字：", singleNumber(nums)) // Output: 5

	// 2.示例输入：回文数
	exampleNum := 121                                                // 示例输入
	fmt.Println("2.", exampleNum, "判断回文数", isPalindrome(exampleNum)) // Output: true
	// 3.判断字符是否有效
	validStr := ")[]{}" // 示例输入
	fmt.Println("3.", validStr, "判断字符是否有效", isValid(validStr))

	//4.最长公共前缀
	strs := []string{"flower", "flow", "flight"} // 示例输入
	fmt.Println("4.", strs, "最长公共前缀：", longestCommonPrefix(strs))

	// 5.删除排序数组中的重复项
	sortedNums := []int{0, 0, 1, 1, 2, 2, 3, 3, 4} // 示例输入
	num, newSortedNums := removeDuplicates(sortedNums)
	fmt.Println("5.", sortedNums, "删除重复项后的长度：", num, "删除重复项后的数组：", newSortedNums)

	//6.加一
	digits := []int{1, 2, 5}                              // 示例输入
	fmt.Println("6.", digits, "加一后的结果：", plusOne(digits)) // Output: [1, 2, 4]

	// 7. 合并区间
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}   // 示例输入
	fmt.Println("7.", intervals, "合并后的区间：", merge(intervals)) // Output: [[1,6],[8,10],[15,18]]
	// 8. 两数之和
	nums2 := []int{2, 7, 11, 15}                                                // 示例输入
	target := 9                                                                 // 示例目标值
	fmt.Println("8.", nums2, "目标值：", target, "两数之和的索引：", twoSum(nums2, target)) // Output: [0,1]
}

// 1.只出现一次的数字：
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。
// 可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func singleNumber(nums []int) int {
	m := make(map[int]int)

	// 统计每个数字出现的次数
	for _, num := range nums {
		m[num]++
	}

	// 找到只出现一次的数字
	for num, count := range m {
		if count == 1 {
			return num
		}
	}

	return 0 // 如果没有找到，返回0（理论上不会到达这里）
}

// 2.回文数

// 考察：数字操作、条件判断
// 题目：判断一个整数是否是回文数
func isPalindrome(x int) bool {
	// 负数和末尾为0的数字（但不包括0本身）都不是回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，可以通过 revertedNumber/10 去除中间数字
	return x == revertedNumber || x == revertedNumber/10
}

// 3.有效的括号

// 考察：字符串处理、栈的使用

// 题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if open, exists := mapping[char]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1] // 弹出栈顶元素
		} else {
			stack = append(stack, char) // 压入栈
		}
	}

	return len(stack) == 0 // 如果栈为空，则括号有效
}

// 4.最长公共前缀

// 考察：字符串处理、循环嵌套

// 题目：查找字符串数组中的最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0] // 假设第一个字符串是最长公共前缀

	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 && (len(strs[i]) < len(prefix) || strs[i][:len(prefix)] != prefix) {
			prefix = prefix[:len(prefix)-1] // 缩短前缀
		}
	}

	return prefix
}

// 5.删除排序数组中的重复项

// 考察：数组/切片操作

// 题目：给定一个排序数组，你需要在原地删除重复出现的元素
func removeDuplicates(nums []int) (int, []int) {
	if len(nums) == 0 {
		return 0, []int{}
	}

	// 使用双指针法
	writeIndex := 1 // 写入位置，从第二个元素开始

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] { // 如果当前元素与前一个元素不同
			nums[writeIndex] = nums[i] // 写入新位置
			writeIndex++               // 移动写入位置
		}
	}

	return writeIndex, nums[:writeIndex] // 返回新数组的长度和新数组
}

// 6.加一

// 考察：数组操作、进位处理

// 题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++ // 如果当前位小于9，直接加1
			return digits
		}
		digits[i] = 0 // 如果当前位是9，变为0并继续进位
	}

	// 如果所有位都是9，需要在最前面添加一个1
	return append([]int{1}, digits...)
}

// 7. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
// 可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
// 将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按照区间的起始位置进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]} // 初始化合并后的区间

	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1] // 获取最后一个合并的区间
		current := intervals[i]       // 当前区间

		if current[0] <= last[1] { // 如果有重叠
			last[1] = max(last[1], current[1]) // 合并区间
		} else {
			merged = append(merged, current) // 没有重叠，添加当前区间
		}
	}

	return merged
}

//
// 基础
// 两数之和

// 考察：数组遍历、map使用

// 题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func twoSum(nums []int, target int) []int {
	m := make(map[int]int) // 创建一个 map 来存储数字和它们的索引

	for i, num := range nums {
		complement := target - num            // 计算补数
		if j, found := m[complement]; found { // 检查补数是否在 map 中
			return []int{j, i} // 如果找到，返回索引
		}
		m[num] = i // 将当前数字和它的索引存入 map
	}

	return nil // 如果没有找到，返回 nil
}
