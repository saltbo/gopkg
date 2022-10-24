package sliceutil

// In returns a bool that's true if the slice contains the value
func In[T comparable](a T, list []T) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}
