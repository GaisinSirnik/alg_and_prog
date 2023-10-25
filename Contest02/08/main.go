import (
"strings"
)

func isPalindrome(a string) bool {
var b string
for _, char := range a {
if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') { 
b += strings.ToLower(string(char))
		}
	}
for i := 0; i < len(b)/2; i++ {
if b[i] != b[len(b)-i-1] { 
return false
		}
	}
return true
}