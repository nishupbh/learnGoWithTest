package iterators

func iterator(str string, n int) string {
	strValue := ""
	for i := 0; i < n; i++ {
		strValue += str
	}
	return strValue
}
