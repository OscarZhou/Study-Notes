package main

func TrimJSONString(s string) string {
	var newString []byte
	for i := 0; i < len(s); i++ {

		if (s[i] == '\n') || (s[i] == ' ') || (s[i] == '\t') {
			continue
		}
		newString = append(newString, byte(s[i]))
	}
	return string(newString)
}
