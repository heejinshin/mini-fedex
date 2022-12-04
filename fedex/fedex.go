package fedex // 메인 패키지가 아니어서 못돌린다? 

import "fmt"

type FedexSender struct {

}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("소포를 보냅니다. %v\n", parcel)
}



