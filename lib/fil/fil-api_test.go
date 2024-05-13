package fil

import (
	"log"
	"testing"
)

func TestGetBlock(t *testing.T) {
	cid := "bafy2bzacebkfx4wbywkskjfperofr7mvjudsxoisvipgsiqx6kpf6lrntwuog"
	got, err := GetBlock(cid)
	if err != nil {
		t.Errorf("GetBlock() error = %v", err)
		return
	}

	log.Print(got)
}

func TestGetChainNotify(t *testing.T) {
	got, err := GetChainHead()
	if err != nil {
		t.Errorf("GetBlock() error = %v", err)
		return
	}

	log.Print(got)
}

func TestNetPeers(t *testing.T) {
	got, err := NetPeers()
	if err != nil {
		t.Errorf("GetBlock() error = %v", err)
		return
	}

	log.Print(got)
}

func TestGetTipSetByHeight(t *testing.T) {
	got, err := GetTipSetByHeight(uint64(3905052))
	if err != nil {
		t.Errorf("GetBlock() error = %v", err)
		return
	}

	log.Print(got)
}
