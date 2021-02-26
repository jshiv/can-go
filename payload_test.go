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

	dataLE := payload.PackLittleEndian()
	// for _, d := range payload.Data {
	// 	//packed |= uint8(packed >> (i * 8))
	// 	fmt.Println(fmt.Sprintf("%08b", d))
	// }
	fmt.Println(dataLE)
	fmt.Println(fmt.Sprintf("%b", dataLE))
	//data := Data{0x40, 0x23, 0x01, 0x12}

	// dataLE := data.PackLittleEndian()
	// fmt.Println(dataLE)
	// fmt.Println(strconv.FormatUint(dataLE, 2))
}

func TestPackBigEndian(t *testing.T) {
	// 4621538819433299968
	// 100000000100011000000010001001000000000000000000000000000000000
	data := []byte{0x40, 0x23, 0x01, 0x12}
	payload := Payload{Data: data}

	// dataBE := payload.PackBigEndian()
	// fmt.Println(dataBE)
	// fmt.Println(fmt.Sprintf("%b", dataBE))

	for _, d := range payload.Data {
		//packed |= uint8(packed >> (i * 8))
		fmt.Println(fmt.Sprintf("%08b", d))
	}
	// data := Data{0x40, 0x23, 0x01, 0x12}

	// dataBE := data.PackBigEndian()
	// fmt.Println(dataBE)
	// fmt.Println(strconv.FormatUint(dataBE, 2))
}
