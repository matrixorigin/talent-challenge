package metadata

type Attribute struct {
	Type uint32 // type of attribute
	Name string // name of attribute
}

type Metadata struct {
	Num   int
	Attrs []Attribute
}
