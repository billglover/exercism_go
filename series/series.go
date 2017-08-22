package series

const testVersion = 2

func All(n int, s string) (series []string) {
	for i := 0; i < len(s) && i+n <= len(s); i++ {
		series = append(series, s[i:i+n])
	}

	return series
}

func UnsafeFirst(n int, s string) (f string) {
	return s[:n]
}

func First(n int, s string) (f string, ok bool) {
	if n > len(s) {
		return f, false
	}
	return s[:n], true
}
