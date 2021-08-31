package connection

var MEMORY = make(map[string]int)

func Setup() {
	MEMORY["apple"] = 1
}
