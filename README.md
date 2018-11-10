[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/asmcos/requests/master/LICENSE)

# pathlib

A golang path library, it is easy to use. Similar to Python patblib.

> 该项目目前处于积极发展期😀，欢迎 issues

# Installation

```
go get -u github.com/small-tk/pathlib
```

# Enjoy


```go

package main

import "github.com/small-tk/pathlib"

func main (){
	p := pathlib.New("../")

	fmt.Println(p.Absolute())
	fmt.Println(p.Cwd())
	fmt.Println(p.Exists())
	fmt.Println(p.Parent())
}

```