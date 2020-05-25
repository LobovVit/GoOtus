package itoa

import (
	"strings"
)

var lastRune rune
var preLastRune rune

func unpack(packetString string) string {
	var retVal string
	var repeatCnt int
	var repeatableRune rune
	for _, currentRune := range packetString + "X" {

		//это число
		if isNum(currentRune) && string(repeatableRune) != "" {
			repeatCnt = repeatCnt*10 + int(currentRune-'0')
		} //else {repeatCnt = 0}

		// это не число а предыдущее число
		if !isNum(currentRune) && isNum(lastRune) {
			//fmt.Println(currentRune)
			//fmt.Print(repeatCnt)
			//fmt.Print(string(currentRune))
			retVal = retVal + repeatRune(repeatCnt, repeatableRune) + string(currentRune)
			//fmt.Println(repeatCnt)
			repeatCnt = 0
		}

		// это не число и предыдущее не число
		if !isNum(currentRune) && !isNum(lastRune) {
			//fmt.Print(string(currentRune))
			if string(currentRune) != `\` || string(lastRune) == `\` {
				retVal = retVal + string(currentRune)
			}
		}

		// это не число - записываем что повторять
		if !isNum(currentRune) {
			repeatableRune = currentRune
		}

		preLastRune = lastRune
		lastRune = currentRune

	}
	retVal = retVal[:len(retVal)-1]
	return retVal
}

func isNum(s rune) bool {

	if strings.Compare(string(s), "0") >= 0 && strings.Compare(string(s), "9") < 0 && (string(lastRune) != `\` || string(preLastRune) == `\`) {
		return true
	} else {
		return false
	}

}

func repeatRune(cnt int, r rune) string {
	var retString string

	for i := 1; i < cnt && string(r) != "\x00"; i++ {
		retString = retString + string(r)
	}

	return retString
}
