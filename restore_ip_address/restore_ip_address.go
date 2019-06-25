package main

import "fmt"

/*
利用循环递归解决子问题。对于每个段内数来说，最多3位最少1位，
所以在每一层可以循环3次，来尝试填段。因为IP地址最多4个分段，当层数是3的时候说明已经尝试填过3个段了，那么把剩余没填的数段接到结尾即可。

这个过程中要保证的是填的数是合法的，最后拼接的剩余的数也是合法的。

注意开头如果是0的话要特殊处理，如果开头是0，判断整个串是不是0，不是的话该字符就是非法的。因为001，01都是不对的。
*/

func restoreIpAddresses(s string) []string {
	n := len(s)
	if n < 4 || n > 12 {
		return []string{}
	}

	res := []string{}
	combination := make([]string, 4)

	var dfs func(int, int)
	dfs = func(idx, begin int) {
		if idx == 3 {
			temp := s[begin:]
			if isOK(temp) {
				combination[3] = temp
				res = append(res, IP(combination))
			}
			return
		}

		// 剩余 IP 段，最多需要的宽度
		maxRemain := 3 * (3 - idx)

		for end := begin + 1; end <= n-(3-idx); end++ {
			if end+maxRemain < n {
				// 后面的 IP 段 至少有一个超过了 3 个字符
				// 说明此IP段短了
				continue
			}

			if end-begin > 3 {
				// 此 IP 段长度，超过 3 位了
				break
			}

			temp := s[begin:end]
			if isOK(temp) {
				combination[idx] = temp
				dfs(idx+1, end)
			}
		}
	}

	dfs(0, 0)

	return res
}

// IP 返回 s 代表的 IP 地址
func IP(s []string) string {
	return fmt.Sprintf("%s.%s.%s.%s", s[0], s[1], s[2], s[3])
}

func isOK(s string) bool {
	// "0"  可以
	// "01" 不可以
	if len(s) > 1 && s[0] == '0' {
		return false
	}

	if len(s) < 3 {
		return true
	}

	// len(s) == 3
	switch s[0] {
	case '1':
		return true
	case '2':
		if '0' <= s[1] && s[1] <= '4' {
			return true
		}
		if s[1] == '5' && '0' <= s[2] && s[2] <= '5' {
			return true
		}
	}

	return false
}
