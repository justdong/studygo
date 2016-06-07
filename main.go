package main

import (
	"flag"
	"net/http"
	"io/ioutil"
	"fmt"
	"os"
	"os/exec"
	"github.com/guotie/gogb2312"
)
const HOSTPATH = "C:\\Windows\\System32\\drivers\\etc\\hosts"

func main() {
	host_url := flag.String("U","https://coding.net/u/scaffrey/p/hosts/git/raw/master/hosts","google host url")


	flag.Parse()

	hostText,err := getHostsFromWebUrl(*host_url)
	if err != nil {
		fmt.Println(err)
		return
	}

	file ,err := os.OpenFile(HOSTPATH,os.O_CREATE|os.O_RDWR,0660)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	_ , err1 := file.WriteString(hostText)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	cmd := exec.Command("cmd.exe","/c","ipconfig/flushdns")
	bytes,err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(gogb2312.ConvertGB2312String(string(bytes)))


	pingcmd := exec.Command("cmd.exe","/c","ping www.google.com")
	pingBytes,err1 := pingcmd.Output()
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(gogb2312.ConvertGB2312String(string(pingBytes)))

}
//从url获取
func getHostsFromWebUrl(url string) (hostText string,err error) {
	rep,err := http.Get(url)
	if err != nil {
		return "",err
	}
	bytes ,err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return "",err
	}
	return string(bytes),nil
}
