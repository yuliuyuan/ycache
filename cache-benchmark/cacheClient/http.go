package cacheClient

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type HttpClient struct {
	client *http.Client
	server string
}

func (c *HttpClient) Run(cmd *Cmd) {
	if cmd.Name == "get" {
		resp, err := c.client.Get(c.server + "get/" + cmd.Key)
		if err != nil {
			log.Println(cmd.Key)
			panic(err)
		}
		cmd.Value = resp.Header.Get("value")
		return
	}

	if cmd.Name == "set" {
		reqBytes := strings.NewReader(fmt.Sprintf("key:%v,value:%v", cmd.Key, cmd.Value))
		resp, err := c.client.Post(c.server+"set", "application/json", reqBytes)
		if err != nil {
			panic("err")
		}
		log.Println("http set resp: %v", resp)
		return
	}
	panic("unknown cmd name " + cmd.Name)
}

func (c *HttpClient) PipelinedRun(cmds []*Cmd) {

}

func newHTTPClient(server string) *HttpClient {
	client := &http.Client{Transport: &http.Transport{MaxIdleConnsPerHost: 1}}
	return &HttpClient{client, "http://" + server}
}
