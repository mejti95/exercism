package logs
import (s "strings"
        "fmt"
       "unicode/utf8")
// Application identifies the application emitting the given log.
func Application(log string) string {
    applications := map[string]int32{"recommendation":'❗',"search":'🔍',"weather":'☀'}
	for _, v := range log {
        for k, x := range applications {
            if v == x {
                return k
            }
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return s.Replace(log,fmt.Sprintf("%c",oldRune),fmt.Sprintf("%c",newRune), -1)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) <= limit {
        return true
    }
    return false
}
