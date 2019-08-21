package main

// 实现匹配'.'和'*'正则表达式的函数
func fun1(a string, p string) bool {
	if a == "" || p == "" {
		return false
	}
	return core(a, p, 0, 0)
}
func core(a, p string, i, j int) bool {
	if i == len(a)-1 && j == len(p)-1 {
		return true
	}
	if i < len(a)-1 && j == len(p)-1 {
		return false
	}
	if p[j+1] == '*' {
		if p[j] == a[i] || p[j] == '.' && i != len(a)-1 {
			return core(a, p, i+1, j+2) || // 进入下一状态
				core(a, p, i+1, j) || // 保持当前状态即判断是否还有相同字符
				core(a, p, i, j+2) // 忽略'a*',进入下一状态
		} else {
			return core(a, p, i, j+2)
		}
	}
	if a[i] == p[j] || (p[j] == '.' && i != len(a)-1) {
		return core(a, p, i+1, j+1)
	}
	return false
}
