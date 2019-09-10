package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/AlekSi/zabbix"
)

func findokokok(name, key_, value, group, ip string) (error, string) {
	var (
		username string
		password string
	)
	username = "diaodebao"
	password = "ftp.http.2"
	zbxapi := zabbix.NewAPI(`http://zb.xnsudai8.com/zabbix/api_jsonrpc.php`)
	_, err := zbxapi.Login(username, password)
	if err != nil {
		log.Print(err)
		return err, ""
	}
	var (
		hostid   string
		hostname string
		//name     string
		//key_     string
		//ip       string
		//group    string
	)
	//运维设定参数
	hostid = "10084"
	hostname = "127.0.0.1"
	//应用传递参数
	//name = "userCenter dubbo"
	//key_ = "maxprocess"
	//group = "xiaoniusudai"
	//ip = "192.168.111.111"

	newkey := fmt.Sprintf("%s-%s-%s", group, ip, key_)
	newname := fmt.Sprintf("%s-%s-%s", group, ip, name)
	_, ok := checkKey[newkey]
	var out string
	if !ok {
		res, err := zbxapi.Call("item.create", zabbix.Params{
			"delay":       30,
			"hostid":      hostid,
			"key_":        newkey,
			"name":        newname,
			"type":        2,
			"value_type":  4,
			"description": "",
			"history":     7,
		})
		if err == nil {
			checkKey[newkey] = "1"
		}
		log.Println(err, res)
		out += fmt.Sprint(res)
		res, err = zbxapi.Call("trigger.create", zabbix.Params{
			"description": newkey,
			"expression":  fmt.Sprintf(`{%s:%s.str(okokok,#1)}=0`, hostname, newkey),
			"priority":    4,
		})
		log.Println(err, res)
		out += fmt.Sprint(res)

	}
	cmd := exec.Command("zabbix_sender", "-z", "172.16.0.180", "-p", "10051", "-s", "127.0.0.1", "-k", newkey, "-o", value)
	out2, err := cmd.CombinedOutput()
	return err, string(out2) + out
}
