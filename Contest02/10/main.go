import "math"
func shift(data []int, steps int) {
temp := make([]int, len(data))
copy(temp, data)
realSteps := int(math.Abs(float64(steps))) % len(data)
if steps < 0 {
realSteps = -realSteps
	}

for i := 0; i < len(data); i++ {
idx := i - realSteps
if idx < 0 {
idx = len(data) + idx
		} else if idx >= len(data) {
idx -= len(data)
		}
		data[i] = temp[idx]
	}
}