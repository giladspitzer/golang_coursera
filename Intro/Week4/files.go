package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	dat, e := ioutil.ReadFile("test.txt")
	fmt.Println(dat, e)

	dat2 := []byte("Hello World")
	ioutil.WriteFile("output.txt", dat2, 0777)

	f, err := os.Open("dt.txt")
	barr := make([]byte, 10)
	nb, err := f.Read(barr)
	f.Close()
	fmt.Println(err, nb)

	os.Create("outfile.txt")
	barr = []byte{1, 2, 3}

	f.Write(barr)
	f.WriteString("Hiii")

}
