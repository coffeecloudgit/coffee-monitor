package fil

import (
	"coffee-monitor/lib/log"
	"github.com/tidwall/gjson"
	"testing"
)

/*
2024/08/09 17:21:17
{"BLSAggregate":{"Data":"r6B6oJJq88woe5Ss9rg/YpUXhFWtEbEuzB0l/cfiA4Kf8in/PM8ZcWjFBbfT5MyhFmuEDQ9RS4Otu4zuU8oGw7KiHVJ+NFOmzhKJfMvHBCAMC+fHBuvVteocg/MPmKF5","Type":2},
"BeaconEntries":[{"Data":"ha9xhPzATliZXhLkzPx07b+M1RJyktmXhoehbehK7N3CwPpNDkBLVaVDDjQnR+d5","Round":7420652}],
"BlockSig":{"Data":"seJi3jBu2huZoDJhQqlD1vHPaLDI80VmNyJFjKOFWOxteeOFotTBjdpwFrqK3tjrGfzO6evgow6T66HuG98Bvocvc46OUtEhAgojv2xgh8vU5uR7eea4vr0UH+nfHRyG","Type":2},
"ElectionProof":{"VRFProof":"pgg1JSHrCdTzpBzzD/B4wMf4fQczD6PtQ3ys6WtL4hkDvDQZjjQU83xa3X219zNRFu8RP9/DgGEYQLiUUm/ri3MIVSuPSH3h4l7bFBpCN444yT2QbO6jDxxRHf31sjO/","WinCount":1},
"ForkSignaling":0,
"Height":3891965,
"Messages":{"/":"bafy2bzacebqzbkk6ck3xv6qnsoennguva5tilbwzfuwctz4ucvsyqn3krolj4"},"Miner":"f02824157","ParentBaseFee":"100",
"ParentMessageReceipts":{"/":"bafy2bzaceb3usskcg7akb2nilwjporrqgfujzoyymnx3m5bzvet6wtunpnnhg"},"ParentStateRoot":{"/":"bafy2bzacebb7uc2wyqu23flmwcqzw6m6ed45jrhzggqmh4ux2nd76pyjiw4xs"},
"ParentWeight":"92408498040",
"Parents":[{"/":"bafy2bzacec4z7qcvbjoapnyfpctcgmg3nzfltc3so3cpp6ghjq7ic2klopuo6"},{"/":"bafy2bzacea7m35gxcueod6l6xfcxygsmvjqxjwear5oliwjawddzcwpffl44c"},{"/":"bafy2bzaceceaphuph4rauvpmiroahx76htgpueomdamt7pesbaz2iphxseov2"},{"/":"bafy2bzacecqtrsdil5ipcn5kle6jbnjrslvsvpenzqktrvuybbbpkme2bwzsu"},{"/":"bafy2bzacedmhyuybtvr3z35m464skzr6qidvlmcaozzyfabg3cjtnqtguwijy"},{"/":"bafy2bzaceapdm2ltxee7b4vftd3gwxtrz54zxstbdewbkn3xqnhrlfetjm5f2"},{"/":"bafy2bzacea3zbwuhvxt27byb3eaw2ln32pl66ujh6msabxtdmk33ukmr6m2ms"},{"/":"bafy2bzacea74ptjxgukkhy3bb5izg24t2wnzkzovl5jnuxizal65zxeandfm4"}],
"Ticket":{"VRFProof":"gGbO6NLDV+TxwVDrY3NO/C0YJSYFA2qlpovegJpgSDtpFlovHnMjeydRuItCkd1kEveRJ9EmjTR8WJQIvdf8abIzHJyu500poS1t42+JV+qGySYe9JHh2lmsoSIqJNHe"},
"Timestamp":1715065350,
"WinPoStProof":[{"PoStProof":3,"ProofBytes":"kpk2yIhzZr6399lcd4MlpjTFFm+Csm3wpppkoN/7bjqhg358GgJzjTyZfYXB2Ss1krvvKIpLJfAYAOhZO0BVXkf9HQkaOrntf7sAKjK6cjUcpQWmKqhQ7k66I3oHODyhFaieJLnGQEljEdBvzfffF/IjIEXIpIQEejZYZ2zhsfLQv/eqL0qJHguXeCiFbMG5rKDa1xsirZW88Aw0A7B95lpKirxvE+RzLcSdJpE9RSSvnyYtGTNDAzRsqXlUfOLG"}]}
*/
func TestGetBlock(t *testing.T) {
	cid := "bafy2bzacebkfx4wbywkskjfperofr7mvjudsxoisvipgsiqx6kpf6lrntwuog"
	got, err := GetBlock(cid)
	if err != nil {
		t.Errorf("GetBlock() error = %v", err)
		return
	}

	log.Log.Println(got)
}

func TestGetChainNotify(t *testing.T) {
	got, err := GetChainHead()
	if err != nil {
		t.Errorf("GetBlock() error = %v", err)
		return
	}

	log.Log.Println(got)
}

func TestNetPeers(t *testing.T) {
	got, err := NetPeers()
	if err != nil {
		t.Errorf("GetBlock() error = %v", err)
		return
	}

	log.Log.Println(got)
}

func TestGetTipSetByHeight(t *testing.T) {
	got, err := GetTipSetByHeight(uint64(4191604))
	if err != nil {
		t.Errorf("GetTipSetByHeight() error = %v", err)
		return
	}

	cids := got.Get("Cids").Array()

	for _, cidJSON := range cids {
		//log.Logger.Info("cid", "index", cidIndex, "val", gjson.Get(cidJSON.Raw, "/").String())
		cid := gjson.Get(cidJSON.Raw, "/").String()

		log.Logger.Info(cid)
	}

	log.Log.Println(cids)
}

func TestSyncState(t *testing.T) {
	got, err := SyncState()
	if err != nil {
		t.Errorf("SyncState() error = %v", err)
		return
	}

	log.Log.Println(got)
}
