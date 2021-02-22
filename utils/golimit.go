package utils

type limit struct {
	Num int
	C   chan struct{}
}

func NewG(num int) *limit {
	return &limit{
		Num: num,
		C:   make(chan struct{}, num),
	}
}

func (g *limit) Run(f func()) {
	g.C <- struct{}{}
	go func() {
		f()
		<-g.C
	}()
}
