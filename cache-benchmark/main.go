package main

import "yly/ycache/cache-benchmark/cacheClient"

func main() {
	cmd1 := cacheClient.Cmd{
		Name: "set",
		Key: "yuliuyuan",
		Value: "guoying",
	}
	cmd2 := cacheClient.Cmd{
		Name: "get",
		Key: "yuliuyuan",
	}
	httpClient := cacheClient.New("http", "localhost:8080/")
	httpClient.Run(&cmd1)
	httpClient.Run(&cmd2)
}