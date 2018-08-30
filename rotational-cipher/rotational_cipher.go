package rotationalcipher

const az = "abcdefghijklmnopqrstuvwxyz"

// RotationalCipher takes a string and a cipher key, encodes it using a rotational cipher
// and returns the encoded string.
func RotationalCipher(s string, k int) string {

	b := []byte(s)

	for i := range b {
		if b[i] >= 'A' && b[i] <= 'Z' {
			c := int(b[i])
			rc := ((c - 'A' + (k % 26)) % 26) + 'A'
			b[i] = byte(rc)
		}

		if b[i] >= 'a' && b[i] <= 'z' {
			c := int(b[i])
			rc := ((c - 'a' + (k % 26)) % 26) + 'a'
			b[i] = byte(rc)
		}
	}

	return string(b)
}
