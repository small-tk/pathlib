[![Go Report Card](https://goreportcard.com/badge/github.com/small-tk/pathlib)](https://goreportcard.com/report/github.com/small-tk/pathlib) [![license](https://img.shields.io/github/license/small-tk/pathlib.svg)](https://github.com/small-tk/pathlib/blob/master/LICENSE)

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
	p := New("test.txt")

	fmt.Println(p.Absolute())
	fmt.Println(p.Cwd())
	fmt.Println(p.Parent())
	fmt.Println(p.Touch())

	fmt.Println(p.Unlink())
	fmt.Println(p.MkDir(os.ModePerm, true))
	fmt.Println(p.RmDir())
	fmt.Println(p.Open())
	fmt.Println(p.Chmod(os.ModePerm))
	fmt.Println(p.Chmod(os.ModePerm))

	fmt.Println(p.Exists())
	fmt.Println(p.IsDir())
	fmt.Println(p.IsFile())
	fmt.Println(p.IsAbs())
}

```