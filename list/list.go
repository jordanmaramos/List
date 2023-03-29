package list

// Methods of a List
type IList interface {
	init()
	add(value int)
	addOnIndex(value int, index int)
	remove()
	removeOnIndex(index int)
	get(index int) int
	set(value int, index int)
	length() int
}

func Hello() string{
	return "Hello World"
}