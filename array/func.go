package array

/**
 * 计算两个字符串数组差集
 */
func Difference(a, b []string) []string {
	mb := map[string]bool{}
	ma := map[string]bool{}

	for _, x := range a {
		ma[x] = true
	}

	for _, x := range b {
		mb[x] = true
	}

	ab := []string{}

	for _, x := range a {
		if _, ok := mb[x]; !ok {
			ab = append(ab, x)
		}
	}

	for _, x := range b {
		if _, ok := ma[x]; !ok {
			ab = append(ab, x)
		}
	}

	return ab
}


/**
 * 判断两个数组内元素是否一致，不考虑顺序
 */
func ArrayEqual(a, b []string) bool {
	mb := map[string]bool{}
	ma := map[string]bool{}

	for _, x := range a {
		ma[x] = true
	}

	for _, x := range b {
		mb[x] = true
	}

	for _, x := range a {
		if _, ok := mb[x]; !ok {
			return false
		}
	}

	for _, x := range b {
		if _, ok := ma[x]; !ok {
			return false
		}
	}

	return true
}