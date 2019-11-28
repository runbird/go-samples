package container

import "fmt"

func main() {
	testRepeatingSubStr()
}

//寻找最长不含有重复字符的子串
// awwbk --> wbk
// ssssss --> s
func findLengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		//lastOccurred[ch]不一定存在 可能为0  ==>lastI, ok := lastOccurred[ch]
		//	if lastOccurred[ch] >= start {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func testRepeatingSubStr() {
	fmt.Println(findLengthOfNonRepeatingSubStr("aaaaaa"))
	fmt.Println(findLengthOfNonRepeatingSubStr("aaabaaa"))
	fmt.Println(findLengthOfNonRepeatingSubStr("abcdefgaaa"))
}
