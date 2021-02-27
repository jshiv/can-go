package can

import (
	"fmt"
	"testing"
)

func TestPackLittleEndian(t *testing.T) {
	// 302064448
	// 10010000000010010001101000000
	data := []byte{0x40, 0x23, 0x01, 0x12}
	payload := Payload{Data: data}

	dataLittleEndian := payload.PackLittleEndian()
	fmt.Println(dataLittleEndian)
	fmt.Println(fmt.Sprintf("%b", dataLittleEndian))

	// if fmt.Sprintf("%08b", dataLittleEndian) != "10010000000010010001101000000" {
	// 	t.Errorf("fmt.Sprintf('%08b', dataLittleEndian)  = %s); want 10010000000010010001101000000", fmt.Sprintf("%08b", dataLittleEndian))
	// }
}

func TestPackBigEndian(t *testing.T) {
	// 4621538819433299968
	// 100000000100011000000010001001000000000000000000000000000000000
	data := []byte{0x40, 0x23, 0x01, 0x12}
	payload := Payload{Data: data}

	dataBigEndian := payload.PackBigEndian()
	fmt.Println(dataBigEndian)
	fmt.Println(fmt.Sprintf("%b", dataBigEndian))

	// if fmt.Sprintf("%08b", dataBigEndian) != "10000000010001100000001000100100" {
	// 	t.Errorf("fmt.Sprintf('%08b', dataBigEndian)  = %s); want 10000000010001100000001000100100", fmt.Sprintf("%08b", dataBigEndian))
	// }
}
