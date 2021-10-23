package main

import (
	"fmt"
	"os"
)

func flagTest() {
	var flagtest FlagSetter
	flagtest.FlagSetting()
	fmt.Println(*flagtest.Setup)
	fmt.Println(*flagtest.Reset)
	fmt.Println(*flagtest.Decrypt)
	fmt.Println(*flagtest.Encrypt)
	fmt.Println(*flagtest.FileName)
}

func ExampleFlag() {
	// set flag value of test
	old := os.Args
	testArgs := []string{"serial", "-setup", "-reset", "-decrypt", "-encrypt", "-filename", "hoge.txt"}
	os.Args = testArgs
	flagTest()
	os.Args = old
	// Output:
	// true
	// true
	// true
	// true
	// hoge.txt
}
