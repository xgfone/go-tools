package compare

type Comparer interface {
	// return a positve integer if > v, 0 if == v, a negative if < v
	Compare(v interface{}) int
}
