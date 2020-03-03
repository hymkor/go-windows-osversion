package main

import (
	"fmt"

	"github.com/zetamatta/go-windows-osversion"
)

func main() {
	v := osversion.Query()

	fmt.Printf("Major=%d\n", v.Major)
	fmt.Printf("Minor=%d\n", v.Minor)
	fmt.Printf("BuildNumber=%d\n", v.Build)
	fmt.Printf("PlatformId=%d\n", v.PlatformId)
}
