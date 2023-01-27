package playfair_cipher

import "strings"

// Table 加密解密使用的密码表
type Table [][]rune

// MakeTable 根据key制作密码表
func MakeTable(key string, options *Options) Table {

	// 先做一个A-Z的集合
	characterSet := make(map[rune]struct{})
	for i := 0; i < 26; i++ {
		character := rune('A' + i)
		// 要拉黑一个字符
		if character == options.BlackCharacter {
			continue
		}
		characterSet[character] = struct{}{}
	}

	runeSlice := make([]rune, 0)

	// 然后把秘钥给用了
	for _, character := range []rune(key) {
		if character >= 'a' && character <= 'z' {
			character -= 32
		}
		if _, exists := characterSet[character]; !exists {
			continue
		}
		delete(characterSet, character)
		runeSlice = append(runeSlice, character)
	}

	// 然后把剩余没用的给用了
	for i := 0; i < 26; i++ {
		character := rune('A' + i)
		if _, exists := characterSet[character]; !exists {
			continue
		}
		delete(characterSet, character)
		runeSlice = append(runeSlice, character)
	}

	// 然后把行转为矩阵
	table := make([][]rune, 0)
	row := make([]rune, 0)
	for _, character := range runeSlice {
		row = append(row, character)
		if len(row) == 5 {
			table = append(table, row)
			row = make([]rune, 0)
		}
	}

	return table
}

// 把加密使用的表格转为字符串返回，用于观察表格长啥样
// 返回数据样例：
//
//	 [
//		[ I, C, L, O, M ]
//		[ P, H, D, R, Z ]
//		[ U, V, F, Y, B ]
//		[ G, X, T, Q, E ]
//		[ S, N, K, W, A ]
//	]
func (x Table) String() string {
	sb := strings.Builder{}
	sb.WriteString("[\n")
	for _, line := range x {
		sb.WriteString("\t[ ")
		for index, column := range line {
			sb.WriteRune(column)
			if index+1 != len(line) {
				sb.WriteString(",")
			}
			sb.WriteString(" ")
		}
		sb.WriteString("]\n")
	}
	sb.WriteString("]")
	return sb.String()
}
