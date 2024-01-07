package main

import benchmark "benchmark/storage"

func main() {

	benchmark.Execute("fio",
		"--randrepeat=1",
		"--ioengine=libaio",
		"--direct=1 ",
		"--gtod_reduce=1",
		"--name=test",
		"--bs=4k",
		"--iodepth=64",
		"--readwrite=randrw",
		"--rwmixread=75",
		"--size=1G ",
		"--filename=/home/datahawk/fiotests/testfile",
	)

}
