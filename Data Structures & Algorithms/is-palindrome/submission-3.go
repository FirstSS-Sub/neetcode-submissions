func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < len(s) || j >= 0 {
		if i < len(s) {
			isSAlnum := (s[i] >= 'a' && s[i] <= 'z') || 
						(s[i] >= 'A' && s[i] <= 'Z') || 
						(s[i] >= '0' && s[i] <= '9')
			if !isSAlnum {
				i++
				continue
			}
		}
		if j >= 0 {
			isTAlnum := (s[j] >= 'a' && s[j] <= 'z') || 
						(s[j] >= 'A' && s[j] <= 'Z') || 
						(s[j] >= '0' && s[j] <= '9')
			if !isTAlnum {
				j--
				continue
			}
		}

		// アルファベットの「大文字」と「小文字」の背番号（ASCIIコード）を2進数で並べると、右から6番目のビットが 0 か 1 という違いしかないのでこれで判定できる
		if s[i] | 0x20 != s[j] | 0x20 {
			return false
		}

		i++
		j--
	}

	return j == -1 && i == len(s)
}
