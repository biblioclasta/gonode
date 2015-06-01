package main

import (
	"time"

	json "github.com/biblioclasta/go-simplejson"
	gonode "github.com/biblioclasta/gonodepkg"
)

func main() {
	gonode.Start(process)
}

func process(cmd *json.Json) (response *json.Json) {
	// Each command property can easily be accessed as shown below
	text := cmd.Get("text").MustString()
	if text == "delay me" {
		time.Sleep(time.Second * 2) // Delay this command beyond the timout
	} else if text == "crash me" {
		time.Sleep(time.Second) // Wait a while to let command #3 execute successfully before panicing
		panic(1)                // Cause a panic!
	}
	return cmd
}
