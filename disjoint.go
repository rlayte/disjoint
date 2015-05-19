package disjoint

type DisjointSet struct {
	data   interface{}
	parent *DisjointSet
	rank   int
}

func MakeSet(item interface{}) *DisjointSet {
	return &DisjointSet{data: item}
}

func (s *DisjointSet) Union(s2 *DisjointSet) {
	r1 := s.Find()
	r2 := s2.Find()

	if r1 == r2 {
		return
	}

	if r1.rank > r2.rank {
		r2.parent = r1
	} else if r2.rank > r1.rank {
		r1.parent = r2
	} else {
		r2.parent = r1
		r1.rank += 1
	}
}

func (s *DisjointSet) Find() *DisjointSet {
	if s.parent != nil {
		s.parent = s.parent.Find()
		return s.parent
	} else {
		return s
	}
}
