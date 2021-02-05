package concrete

type Shoes struct {
	Name string
	Size int32
}

func (s Shoes) SetName(name string) {
	s.Name = name
}

func (s Shoes) SetSize(size int32) {
	s.Size = size
}
