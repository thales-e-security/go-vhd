package main

import (
	"./vhd"
)

func main() {
	vhd.CreateSparseVHD(1024*1024, "test.vhd")
}
