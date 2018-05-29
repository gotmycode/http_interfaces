package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//custom write fn of type struct with 0 fields
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//fmt.Println(resp)

	//declare byteslice - purpose of bs is to hold all the data that gets written into it by that read function
	//make is a built in function that takes a type of a slice, 2nd arg is no of elements/empty spaces we want slice to be initialized with
	//give me a bs of 99999 spaces
	//bs := make([]byte, 99999)
	//resp.Body.Read(bs)
	//fmt.Println(string(bs)) //turn into type string

	lw := logWriter{}

	//goal - not to write byte slice everytime
	//condence above code by using writer interface
	//we have some values in std lib that implement writer interface and take info and send it ouside of our program -- to screen, harddrive, etc
	//io.Copy - take some info from outside of our application and write it or copy it all out to some outside channel
	//os.Stdout -  os.Stdout is a value of type File > File has func called write > implements writer interface
	//resp.Body - 2nd arg is body property of the response struct -- implements reader interface
	//io.Copy(os.Stdout, resp.Body)

	io.Copy(lw, resp.Body)
}

//interfaces are to help guide you down the right path but they dont necessarily make sure that you write exactly correct code
//custom write func
func (logWriter) Write(bs []byte) (int, error) {
	//return 1, nil - just to show incorrect implementation

	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
