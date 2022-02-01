# TreeNode 🌲🌲🌲🌲

Binary Tree Node struct with LeetCode compatible Serialize / Deserialize format

[![Build Status](https://github.com/egregors/TreeNode/actions/workflows/go.yml/badge.svg)](https://github.com/egregors/TreeNode/actions) 


## Install

`go get github.com/egregors/TreeNode`

## Usage

This package provides `TreeNode` data struct and a few function for serialization and deserialization respectively.

```go
package main

import (
	"log"

	tn "github.com/egregors/TreeNode"
)

func main() {
	data := "[1,2,3,null,null,4,5]"
	root, err := tn.NewTreeNode(data)
	if err != nil {
		log.Fatal(err)
	}

	root.Right.Left.Val = 42
	newData := root.String()
	...
}
```

A `NewTreeNode(data string) (*TreeNode, error)` constructor expect LeetCode style formatted string. 
Otherwise, to serialize a tree just call `String()` for `TreeNode` object.

## Benchmarks
```shell
goos: darwin
goarch: amd64
pkg: github.com/egregors/TreeNode
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkEmptyTree-12           13766786                77.96 ns/op
Benchmark1_5Tree-12              2645138               448.3 ns/op
```

# Contributing
Bug reports, bug fixes and new features are always welcome.
Please open issues and submit pull requests for any new code.
