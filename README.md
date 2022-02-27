<p align="center">
  <a href="https://gominima.studio">
  <img alt="go-requests" src="./assets/logo.png" />
</a>
</p>

<p align="center">
<a href="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"> <img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat" /></a>
<a href="https://discord.gg/gRyCr5APmg"> <img src="https://img.shields.io/discord/916969864512548904" /></a>
<img src="https://img.shields.io/tokei/lines/github/gominima/go-requests" />
<img src="https://img.shields.io/github/languages/code-size/gominima/go-requests" />
</p>

## 🦄 Example

```go
package main

import (
	"fmt"

	"github.com/gominima/go-requests"
)

func main() {
	resp, err := goquests.Get(goquests.Request{
		URL:     "https://random-data-api.com/api/users/random_user",
		Data:    make(map[string]interface{}),
		Headers: make(map[string]string),
	})
	if err != nil {
		panic(err)
	}
	fmt.Print(resp.Body)
 }
}

```

## ❓Why Go-requests

Go requests is made for making data fetch wat easier and fun to use, it uses all standard go packages without any other dependency, also is pretty fast and reliable .

<br/>

## ⭐ Contributing

**If you wanna help grow this project or say a thank you!**

1. Give minima a [GitHub star](https://github.com/gominima/go-requests/stargazers)
2. Fork requests and Contribute
4. Join our [Discord](https://discord.gg/gRyCr5APmg) community
