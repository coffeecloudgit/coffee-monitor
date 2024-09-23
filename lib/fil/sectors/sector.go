package sectors

import (
	"coffee-monitor/lib/client"
	config2 "coffee-monitor/lib/config"
	"coffee-monitor/lib/fil"
	"coffee-monitor/lib/log"
	"coffee-monitor/lib/shell"
	"coffee-monitor/lib/util"
	"errors"
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/tidwall/gjson"
	"io"
	"strconv"
	"strings"
)

type Sector struct {
	Id                    uint64
	ExpirationBlockHeight uint64
}

func SendSectorsExpireInfo() error {
	minerSectors := make(map[string]interface{}, 5)
	minerSectors["miner"] = config2.CONF.Fil.Account
	err, sectorsInfo, total := GetSectorsExpireInfo()
	if err != nil {
		return err
	}
	minerSectors["total"] = total
	minerSectors["sectors"] = sectorsInfo

	msg := client.Message{Type: client.SectorsExpireInfo, Data: minerSectors}
	//msgBytes, err := json.Marshal(msg)
	//if err != nil {
	//	return errors.New("message marshal fail")
	//}
	//fmt.Println(string(msgBytes))
	//time.Sleep(2000 * time.Millisecond)
	err2 := client.SendMessage(msg)
	if err2 != nil {
		log.Logger.Info(err2.Error())
		return err2
	}

	return nil
}

func GetSectorsExpireInfo() (error, []*client.ExpireSameDaySectors, uint64) {
	err, _ := shell.GenerateLotusMinerSectorsFile()
	if err != nil {
		return err, nil, 0
	}
	return SectorFileProcessor()
}

func SectorFileProcessor() (error, []*client.ExpireSameDaySectors, uint64) {
	//go client.ConnectServer()
	config := config2.CONF
	if len(config.Fil.Sectors) == 0 {
		return errors.New("sectors file is empty"), nil, 0
	}
	t, err := tail.TailFile(config.Fil.Sectors, tail.Config{
		Follow:   false,                                //是否跟随
		Location: &tail.SeekInfo{Whence: io.SeekStart}, //从文件的什么地方开始读
	})

	if err != nil {
		log.Logger.Info(err.Error())
		return err, nil, 0
	}
	sectors := make([]Sector, 0)
	for line := range t.Lines {
		//log.Logger.Info(line.Text)
		sectorInfo := strings.Fields(line.Text)
		if len(sectorInfo) < 5 {
			log.Logger.Error("sector array length is error,", "sector", line.Text)
			continue
		}
		id, err := strconv.ParseUint(sectorInfo[0], 10, 64)
		if err != nil {
			log.Logger.Error("sector ID is error,", "ID", sectorInfo[0], "sector", line.Text)
			continue
		}
		block, err := strconv.ParseUint(sectorInfo[4], 10, 64)
		if err != nil {
			log.Logger.Error("sector expire block is error,", "block", sectorInfo[4], "sector", line.Text)
			continue
		}
		sector := Sector{Id: id, ExpirationBlockHeight: block}
		sectors = append(sectors, sector)
	}

	chainHead, err := fil.GetChainHead()
	if err != nil {
		return fmt.Errorf("lotus info error:%v", err), nil, 0
	}
	height := gjson.Get(chainHead.Raw, "Height").Uint()

	dayAndSectors := make(map[string]*client.ExpireSameDaySectors)
	dayAndSectorsArray := make([]*client.ExpireSameDaySectors, 0)

	//根据当前高度估算扇区到期时间
	for _, sector := range sectors {
		day := util.GetTimeOfDay(sector.ExpirationBlockHeight, height)
		var daySector *client.ExpireSameDaySectors
		var b bool
		if daySector, b = dayAndSectors[day]; !b {
			daySector = &client.ExpireSameDaySectors{Day: day, From: sector.Id, To: sector.Id, SectorNum: 0}
			dayAndSectors[day] = daySector
			dayAndSectorsArray = append(dayAndSectorsArray, daySector)
		}

		daySector.To = sector.Id
		daySector.SectorNum++
		//dayAndSectors[day] = daySector
		//fmt.Println("sector:", sector.Id, ":", day)
	}
	totalNum := uint64(0)
	for _, item := range dayAndSectorsArray {
		totalNum += item.SectorNum
	}
	//fmt.Println(dayAndSectors)
	return nil, dayAndSectorsArray, totalNum
}
