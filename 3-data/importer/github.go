package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	GITHUB_ENDPOINT = "https://api.github.com/users"
)

type GithubUser struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
}

func GetUser(u string) GithubUser {
	url := fmt.Sprintf("%s/%s", GITHUB_ENDPOINT, u)
	resp, _ := http.Get(url)

	defer resp.Body.Close()

	b, e := ioutil.ReadAll(resp.Body)

	if e != nil {
		fmt.Println("Error2", resp.Body)
	}
	i := &GithubUser{}
	e = json.Unmarshal(b, i)
	return *i

}
