package utils

import (
	"fmt"
	"testing"
)

func TestUUID(t *testing.T) {
	uuid := CreateUUID()
	fmt.Println(uuid)
}
