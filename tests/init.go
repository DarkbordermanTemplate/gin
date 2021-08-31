package tests

import (
	"practice/connection"
)

func Setup() {
	connection.MEMORY["apple"] = 1
}
