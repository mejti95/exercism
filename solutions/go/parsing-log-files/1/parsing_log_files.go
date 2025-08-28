package parsinglogfiles
import "regexp"

func IsValidLine(text string) bool {
    checkRegex := regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
    if checkRegex.MatchString(text){
        return true
    }
    return false
}

func SplitLogLine(text string) []string {
	checkRegex := regexp.MustCompile(`<(~|\*|=|-)*>`)
    return checkRegex.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	checkRegex := regexp.MustCompile(`\".*(?i:password).*\"`)
    matchesFound := 0
    for _, v := range lines {
        if checkRegex.MatchString(v) {
            matchesFound += 1
        }
    }
    return matchesFound
}

func RemoveEndOfLineText(text string) string {
	checkRegex := regexp.MustCompile(`end-of-line\d+`)
    return checkRegex.ReplaceAllString(text,"")
}

func TagWithUserName(lines []string) []string {
	checkRegex := regexp.MustCompile(`User\s+\w+`)
    replaceRegex := regexp.MustCompile(`User\s+`)
	for i, v := range lines {
        if checkRegex.MatchString(v) {
            newString := replaceRegex.ReplaceAllString(checkRegex.FindString(v),"[USR] ")
            lines[i] = newString + " " + v
        }
    }
    return lines
}
