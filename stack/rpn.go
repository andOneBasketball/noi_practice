/*
Reverse Polish notation
https://zq99299.github.io/dsalg-tutorial/dsalg-java-hsp/05/05.html
逆波兰式
中缀表达式转后缀表达式并计算

初始化两个栈：

运算符栈：s1
中间结果栈：s2
从左到右扫描中缀表达式

遇到操作数时，将其压入 s2

遇到运算符时

比较 它 与 s1 栈顶运算符的优先级：

如果 s1 为空，或则栈顶运算符号为 ( ，则将其压入符号栈 s1

如果：优先级比栈顶运算符 高，也将其压入符号栈 s1

如果：优先级比栈顶运算符 低 或 相等，将 s1 栈顶的运算符 弹出，并压入到 s2 中

再重复第 4.1 步骤，与新的栈顶运算符比较（因为 4.3 将 s1 栈顶运算符弹出了）

这里重复的步骤在实现的时候有点难以理解，下面进行解说：

如果 s1 栈顶符号 优先级比 当前符号 高或则等于，那么就将其 弹出，压入 s2 中（循环做，是只要 s1 不为空）

如果栈顶符号为 (，优先级是 -1，就不会弹出，就跳出循环了

跳出循环后，则将当前符号压入 s1

遇到括号时：

如果是左括号 ( ：则直接压入 s1

如果是右括号 )：

则依次弹出 s1 栈顶的运算符，并压入 s2，直到遇到 左括号 为止，此时将这一对括号 丢弃

重复步骤 2 到 5，直到表达式最右端

将 s1 中的运算符依次弹出并压入 s2

依次弹出 s2 中的元素并输出，结果的 逆序 即为：中缀表达式转后缀表达式

*/

package main

import (
	"fmt"
	"strconv"
)

// b1 优先 b2，栈顶为 ( 场景需要入栈
func symbolPriority(b1 byte, b2 byte) bool {
	if (b1 == '*' || b1 == '/') && (b2 == '*' || b2 == '/') {
		return false
	}
	if (b1 == '*' || b1 == '/') && (b2 == '+' || b2 == '-') {
		return true
	}
	if (b1 == '+' || b1 == '-') && (b2 == '+' || b2 == '-') {
		return false
	}
	return false
}

// 核心：栈顶优先级高则出栈，否则入栈；栈顶左括号入栈
func rpn(str string) []string {
	symbols := make([]byte, 0, len(str))
	result := make([]string, 0, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			s := ""
			// 时刻牢记最先判断是否越界
			for i < len(str) && str[i] >= '0' && str[i] <= '9' {
				s += string(str[i])
				i++
				//fmt.Printf("str[%d]=%c, s:%s\n", i, str[i], s)
			}
			// 以防和 for 中的 i++ 冲突
			i--
			result = append(result, s)
			//fmt.Printf("res1: %s\n", result)
		} else if str[i] == '+' || str[i] == '-' || str[i] == '*' || str[i] == '/' {
			// 栈顶优先级低或者栈顶为 (
			if len(symbols) == 0 || symbols[len(symbols)-1] == '(' || symbolPriority(str[i], symbols[len(symbols)-1]) {
				symbols = append(symbols, str[i])
			} else {
				/*
					1*2+2*1*3+3*4+5*6
					error: [1 2 * 2 1 * 3 * 3 4 * 5 6 * + + +]
					right: [1 2 * 2 1 * 3 * + 3 4 * + 5 6 * +]
					result = append(result, string(symbols[len(symbols)-1]))
					symbols = symbols[0 : len(symbols)-1]
				*/
				// 一级注意：需要循环比较一直出栈到栈顶优先级高于自身或者空栈
				for len(symbols) > 0 && !symbolPriority(str[i], symbols[len(symbols)-1]) {
					result = append(result, string(symbols[len(symbols)-1]))
					symbols = symbols[0 : len(symbols)-1]
				}

				symbols = append(symbols, str[i])
			}
			//fmt.Printf("res2: %s, symbols: %s\n", result, string(symbols))
		} else if str[i] == '(' {
			symbols = append(symbols, str[i])
			//fmt.Printf("symbols: %s\n", symbols)
		} else if str[i] == ')' {
			for symbols[len(symbols)-1] != '(' {
				result = append(result, string(symbols[len(symbols)-1]))
				symbols = symbols[0 : len(symbols)-1]
			}

			// ( 出栈
			symbols = symbols[0 : len(symbols)-1]
			//fmt.Printf("res3: %s, symbols: %s\n", result, string(symbols))
		}
	}
	for i := len(symbols) - 1; i >= 0; i-- {
		result = append(result, string(symbols[i]))
	}
	return result
}

func isSymbol(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return true
	}
	return false
}

// 根据 rpn 表达式计算结果
func calc(str []string) int {
	stack := make([]string, 0, len(str))
	result := 0
	for _, s := range str {
		// 注意先入栈，后计算
		stack = append(stack, s)
		if len(stack) > 0 && isSymbol(stack[len(stack)-1]) {
			symbol := stack[len(stack)-1]
			i2, _ := strconv.Atoi(stack[len(stack)-2])
			i1, _ := strconv.Atoi(stack[len(stack)-3])
			if symbol == "+" {
				result = i1 + i2
			} else if symbol == "-" {
				result = i1 - i2
			} else if symbol == "*" {
				result = i1 * i2
			} else if symbol == "/" {
				result = i1 / i2
			}
			stack = stack[0 : len(stack)-3]
			stack = append(stack, fmt.Sprintf("%d", result))
		}
	}
	return result
}

// 判断表达式的有效性，待实现
func validExpression(str string) bool {
	/*
		leftBrackets := make([]byte, 0, len(str))
		symbols := make([]byte, 0, len(str))
	*/
	nums := make([]string, 0, len(str))
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			sNum := ""
			for i < len(str) && str[i] >= '0' && str[i] <= '9' {
				sNum += string(str[i])
				i++
			}
			nums = append(nums, sNum)
		}
	}
	return true
}

// 暂未实现有效表达式的检测
func main() {
	fmt.Printf("%s\n", rpn("1+((2+3)*4)-5"))
	fmt.Printf("%s\n", rpn("11+((21+323)*41)-5"))
	fmt.Printf("%s\n", rpn("1*2+2*1*3+3*4+5*6"))

	fmt.Printf("%d\n", calc(rpn("1+((2+3)*4)-5")))
	fmt.Printf("%d\n", calc(rpn("11+((21+323)*41)-5")))

	fmt.Printf("%s\n", rpn("1++((2+3)*4)-5"))
	//fmt.Printf("%s\n", rpn("1++(2+3)*4)-5"))
}
