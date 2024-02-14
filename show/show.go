package show

import (
	"fmt"

	"github.com/rolcho/showgo/uppercase"
)

func Show(s string) {
	s = uppercase.Uppercase(s)
	fmt.Println(s)
}
