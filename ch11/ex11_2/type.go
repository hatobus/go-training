package intset

type IntSet interface {
	Has(x int) bool
	Add(x int)
	UnionWith(t IntSet)
	Len() int
	Remove(x int)
	Clear()
	Copy() IntSet
	String() string
}
