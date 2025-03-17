package shell

import (
	config2 "coffee-monitor/lib/config"
	"coffee-monitor/lib/log"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

var execTimes = 0

func LotusSyncCheck() error {
	//cmd := exec.Command("ls")
	//log.Logger.Info("LotusSyncCheck start...")
	out, err := exec.Command("bash", "-c", "timeout 20s lotus sync wait").Output()
	if err != nil && !strings.Contains(err.Error(), "exit status 124") {
		log.Logger.Info(err.Error())
		return err
	}
	///

	if err != nil {
		log.Logger.Info("sync 超時......")
	}

	outString := strings.TrimSpace(string(out))
	if strings.HasSuffix(outString, "Done!") {
		if execTimes%8 == 0 {
			log.Logger.Info("节点检测正常")
		}
		execTimes++
		return nil
	}
	//fmt.Println("out string is:")
	log.Logger.Error("out string", "string:", outString)
	log.Logger.Error("节点同步异常，需要添加内部节点！")
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
		cmd := fmt.Sprintf("timeout 12s lotus net connect %s", node)
		log.Logger.Info("add node：", "cmd:", cmd)
		out, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			return err
		}
		log.Logger.Info(string(out))
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

// 獲取對應區塊獎勵
func LotusMinerInfoGetRewardForBlock(blockId string) (error, string) {
	cmd := fmt.Sprintf("timeout 36s lotus-miner info --blocks 2 |grep %s", blockId)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return err, "0 FIL"
	}
	result := string(out)
	resultArray := strings.Split(result, "|")

	if len(resultArray) == 3 {
		return nil, strings.TrimSpace(resultArray[2])
	}

	return fmt.Errorf("error out: %s", result), "0 FIL"
}

/*
*
timeout 36s lotus-miner info --blocks 2 |grep "|"

	Epoch   | Block ID                                                       | Reward
	4795574 | bafy2bzacedcbpl3euk443e3kewu6nd6e4xituhm7m5xfskj4dz3qnegl5txoo | 6.072002358920002932 FIL
	4795438 | bafy2bzaceb3yzpiu7pzoswhw3tcbwu5r7f2mfqm4xy23htnpexvuq6fvfx5xg | 6.072055146770326789 FIL

*
*/
func LotusMinerInfoGetRewardForLastBlocks() (error, map[string]string) {
	rewardMap := make(map[string]string)
	cmd := fmt.Sprintf("timeout 36s lotus-miner info --blocks 6 |grep %s", "|")
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return err, rewardMap
	}
	result := string(out)
	log.Logger.Info("result:", "lines:", result)
	lines := strings.Split(result, "\n")
	for _, line := range lines {
		if strings.Contains(line, "|") {
			resultArray := strings.Split(line, "|")
			if len(resultArray) == 3 {
				rewardMap[strings.TrimSpace(resultArray[1])] = strings.TrimSpace(resultArray[2])
			}
		}
	}
	log.Logger.Info("rewardMap:", "map:", rewardMap)
	return nil, rewardMap
}

func GenerateLotusMinerSectorsFile() (error, string) {
	config := config2.CONF
	if len(config.Fil.Sectors) == 0 {
		return errors.New("sectors file is empty"), ""
	}
	command := fmt.Sprintf("timeout 180s lotus-miner sectors list >%s", config.Fil.Sectors)
	out, err := exec.Command("bash", "-c", command).Output()
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
