package main

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
)

type WeConfig struct {
	WorkDir   string // 工作目录,默认就是当前目录
	RepoDir   string // 软件目录,默认为~/.we
	OutFormat string // 输出格式,默认是文本,可以是json
	Fake      bool   // 是否为模拟,默认为false
}

func ReadConfig(path string) (cnf WeConfig, err error) {
	if path == "" {
		var u *user.User
		u, err = user.Current()
		if err != nil {
			return
		}
		path = u.HomeDir + "/.we/setting.json"
	}
	var data []byte
	data, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &cnf)
	return
}
