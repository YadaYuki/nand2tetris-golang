package util

// Fill is function get filled string by custom char
func Fill(s string, fillBy string, length int) string {
	fillWord := ""
	for i := 0; i < length-len(s); i++ {
		fillWord += fillBy
	}
	return fillWord + s
}
