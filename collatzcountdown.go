package piscine

func CollatzCountdown(start int) int {
	var b byte
	count := 0
	if start > 0 {
		for b = 250; b <= 255; b++ {
			if start == 1 {
				break
			}
			if start%2 == 0 {
				start = start / 2
				count++
			} else if start%2 != 0 {
				start = (start * 3) + 1
				count++
			}
		}
	} else {
		return -1
	}
	return count
}
