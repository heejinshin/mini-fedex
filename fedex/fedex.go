package fedex 

import "fmt"

type FedexSender struct {

}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("소포를 보냅니다. %v", parcel)
}



