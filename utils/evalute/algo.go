package evalute

func CheckContainsString(a, b string) bool {
	if a != "" && b != "" {
		aLen, bLen := len(a), len(b)
		for i := 0; i <= aLen-bLen; i++ {
			if a[i:i+bLen] == b {
				return true
			}
		}
	}
	return false
}

func CheckStartWithString(a, b string) bool {
	if a != "" && b != "" && len(a) >= len(b) {
		if b == a[0:len(b)] {
			return true
		}
	}
	return false
}

func CheckEndWithString(a, b string) bool {
	if a != "" && b != "" && len(a) >= len(b) {
		if b == a[len(a)-len(b):] {
			return true
		}
	}
	return false
}
