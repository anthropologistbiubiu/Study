package main

import "fmt"

//0020. 有效的括号

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。
//
//示例 1:
//
//输入: "()"
//输出: true
//示例 2:
//
//输入: "()[]{}"
//输出: true
//示例 3:
//
//输入: "(]"
//输出: false
//示例 4:
//
//输入: "([)]"
//输出: false
//示例 5:
//
//输入: "{[]}"
//输出: true

func isValid(s string) bool {
	stack := make([]rune, len(s))
	top := -1
	for _, v := range s {
		fmt.Printf("%+v\n", v)
		fmt.Printf("%+T\n", v)
		if v == '(' || v == '[' || v == '{' {
			fmt.Println(stack, top)
			top++
			stack[top] = v
		} else {
			if top == -1 {
				return false
			}
			switch v {
			case ')':
				item := stack[top]
				top--
				if item != '(' {
					return false
				}
			case ']':
				item := stack[top]
				top--
				if item != '[' {
					return false
				}
			case '}':
				item := stack[top]
				top--
				if item != '{' {
					return false
				}
			}
		}
		switch v {

		}
	}
	return true
}

func main() {

	s := "{[()]}"
	fmt.Println(isValid(s))

}
