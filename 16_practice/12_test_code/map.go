package main

import "fmt"

func main() {
	m := map[string]map[int8][]string{
		"android": {
			1: []string{"去投稿", "bilibili://uper/user_center/add_archive/"},
			2: []string{"去分享", "bilibili://uper/user_center/manuscript-list/"},
		},
	}
	res, ok := m["ios"][1]
	fmt.Println(res)
	fmt.Println(ok)
}
