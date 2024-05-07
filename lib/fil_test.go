package lib

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
