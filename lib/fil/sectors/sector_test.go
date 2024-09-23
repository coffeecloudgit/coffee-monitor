package sectors

import (
	"fmt"
	"testing"
)

func TestSectorsProcessor(t *testing.T) {
	err, array, total := SectorFileProcessor()

	fmt.Println(err)

	fmt.Println(array)
	fmt.Println(total)
}

func TestSendSectors(t *testing.T) {
	err := SendSectorsExpireInfo()

	fmt.Println(err)

}
