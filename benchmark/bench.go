package main

import (
	json "github.com/biblioclasta/go-simplejson"
	gonode "github.com/biblioclasta/gonodepkg"
)

func main() {
	gonode.Start(process)
}

func process(cmd *json.Json) (resp *json.Json) {
	return cmd
}
