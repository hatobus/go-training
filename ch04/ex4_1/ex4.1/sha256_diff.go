package sha256diff

import "crypto/sha256"

func GetSHA256Diff(s1, s2 string) int {
	if s1 == s2 {
		return 0
	}

	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))

	diffCnt := 0
	for i := 0; i < len(c1); i++ {
		if c1[i] != c2[i] {
			diffCnt++
		}
	}

	return diffCnt
}
