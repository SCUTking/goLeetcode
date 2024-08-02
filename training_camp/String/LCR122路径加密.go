package String

func pathEncryption(path string) string {

	runes := []rune(path)
	for i, r := range runes {
		if r == '.' {
			runes[i] = ' '
		}
	}
	return string(runes)
}
