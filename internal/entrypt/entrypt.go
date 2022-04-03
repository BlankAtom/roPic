package entrypt

var hashCode = "ztVQhXSFUeCis0yZGKYlgR4mbWjkfHn7rduLP1OoMqI2T8DcA6Bxpaw3E5JNv9"

func Encrypt(s string) string {
	// Get the length of param-s.
	l := len(s)
	lh := len(hashCode)
	// Use length to make slice.
	ss := make([]byte, l+1)
	ssLen := 0

	// Rand an index of hashCode to be the encrypt key.
	//rn := randr.Int() % (len(hashCode) / 2)

	// Build the new string.
	for i := 0; i < l; i++ {
		for j := 0; j < lh; j++ {
			if s[i] == hashCode[j] {
				ss[ssLen] = byte(j) + byte('0')
				ssLen++
			}
		}
	}

	// Return the new string.
	return string(ss)
}

func Decrypt(s string) string {
	slen := len(s)
	ss := make([]byte, slen)

	for i := 0; i < slen; i++ {
		ss[i] = hashCode[int(s[i])-'0']
	}
	return string(ss)
}

//
//func Decrypt(s string) string {
//	// Get the length.
//	l := len(s)
//
//	// Check the length.
//	if l < 1 || l > len(hashCode) {
//		return "error length!"
//	}
//
//	// Get the key.
//	key := int(s[0])
//	ss := make([]byte, l-1)
//
//	// Build old string.
//	for i := 1; i < l; i++ {
//		ss[i-1] = s[i] - hashCode[key+i]
//	}
//
//	// Return the old string.
//	return string(ss)
//}
