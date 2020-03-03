go-windows-osversion
====================

Query Windows-OS version.

[Sample](./cmd/ver/main.go)

```go
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
```

On Windows 8.1

```
$ cmd\ver\ver.exe
Major=6
Minor=3
BuildNumber=9600
PlatformId=2
```

On Windows 10

```
$ cmd\ver\ver.exe
Major=10
Minor=0
BuildNumber=18363
PlatformId=2
```

This package uses RtlGetVersion API which does not depend on manifest files.

Thanks to [YAMAMOTO's documents (written in Japanese)](http://yamatyuu.net/computer/program/sdk/base/RtlGetVersion/index.html)
