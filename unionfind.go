package unionfind

type UnionFind struct {
	parent []int
}

func New(sz int) *UnionFind {
	uf := new(UnionFind)
	uf.parent = make([]int, sz)

	for i := 0; i < sz; i++ {
		uf.parent[i] = -1
	}

	return uf
}

func (uf *UnionFind) Root(x int) int {
	if uf.parent[x] < 0 {
		return x
	}

	uf.parent[x] = uf.Root(uf.parent[x])
	return uf.parent[x]
}

func (uf *UnionFind) Same(x, y int) bool {
	return uf.Root(x) == uf.Root(y)
}

func (uf *UnionFind) Size(x int) int {
	return -uf.parent[uf.Root(x)]
}

func (uf *UnionFind) Union(x, y int) bool {
	rx, ry := uf.Root(x), uf.Root(y)

	if rx == ry {
		return false
	}

	if uf.parent[rx] > uf.parent[ry] {
		rx, ry = ry, rx
	}

	uf.parent[rx] += uf.parent[ry]
	uf.parent[ry] = rx

	return true
}
