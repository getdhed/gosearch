package index

type Index struct {
	Key   string
	Value []int
}

func NewIndex(s string, id int) *Index {

	return &Index{
		Key:   s,
		Value: []int{},
	}
}
