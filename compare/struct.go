package compare

type Comparer interface {
	// return 1 if > v, 0 if == v, -1 if < v
	Compare(v interface{}) int
}
