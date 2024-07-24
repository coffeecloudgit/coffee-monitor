package shell

import (
	config2 "coffee-monitor/lib/config"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func LotusSyncCheck() error {
	//cmd := exec.Command("ls")
	//log.Println("LotusSyncCheck start...")
	out, err := exec.Command("bash", "-c", "timeout 36s lotus sync wait").Output()
	if err != nil && err.Error() != "exit status 124" {
		return err
	}

	outString := strings.TrimSpace(string(out))
	if strings.HasSuffix(outString, "Done!") {
		log.Println("节点检测正常")
		return nil
	}
	fmt.Println("out string is:")
	fmt.Println(outString)
	fmt.Println("节点同步异常，需要添加内部节点！")
	return LotusNetAddPeer()
}

func LotusNetAddPeer() error {
	config := config2.CONF

	if len(config.Lotus.Nodes) <= 0 {
		return errors.New("config lotus nodes is null")
	}

	for _, node := range config.Lotus.Nodes {
		if len(node) <= 0 {
			continue
		}
		cmd := fmt.Sprintf("timeout 36s lotus net connect %s", node)
		log.Println("add node：", cmd)
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			return err
		}
		log.Println(out)
		time.Sleep(5000 * time.Millisecond)
	}
	return nil
}

func LotusMinerInfo() (error, string) {
	out, err := exec.Command("bash", "-c", "timeout 36s lotus-miner info").Output()
	if err != nil {
		return err, ""
	}
	return nil, string(out)
}

//try:
//out = sp.getoutput("timeout 36s lotus sync wait")
//print("chain_check:")
//print(out)
//if out.endswith("Done!"):
//print("true")
//return True
//server_post(machine_name, "节点同步出错，请及时排查！")
//return False
//except Exception as e:
//print("Fail to send message: " + e)
