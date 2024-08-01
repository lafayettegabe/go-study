package iteration

func Repeat(ch string, x int) string {
	var res string
	for i := 0; i < x; i++ {
		res += ch
	}
	return res
}
