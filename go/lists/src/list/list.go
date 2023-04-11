package list

// Methods of a List
type IList interface {
	Init()
	Add(value int)
	AddOnIndex(value int, index int)
	Remove() error
	RemoveOnIndex(index int) error
	Get(index int) int
	Set(value int, index int)
	Length() int
	PrintList()
}
