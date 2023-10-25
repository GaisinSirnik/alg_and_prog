import ("strconv")
func isLucky(a string) bool {
first, second := 0, 0
for _, b := range a[:3] {
d, _ := strconv.Atoi(string(b))
first += d
	}
for _, c := range a[3:] {
d, _ := strconv.Atoi(string(c))
second += d
	}
return first == second
}