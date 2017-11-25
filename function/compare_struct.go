package function

// Comparer is an interface to compare the value whose type implements it.
type Comparer interface {
	// return a positve integer if > v, 0 if == v, a negative if < v
	Compare(v interface{}) int
}
