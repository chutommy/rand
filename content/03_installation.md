---
weight: 3
title: "Installation"
draft: false
---

## Installation

1. Download the package:

```shell
$ go get -u github.com/chutommy/rand
```

2. Import it in your code:

```
import "github.com/chutommy/rand"
```

3. (Optional) Import `math/rand` and/or `crypto/rand`. Use aliases to
   distinguish the packages:

```
import (
   crand "crypto/rand"
   mrand "math/rand"

   "github.com/chutommy/rand"
)
```
