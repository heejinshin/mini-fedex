package koreapost 

import "fmt"

type PostSender struct {

} // 구조체 이렇게 구현없이 비워둬도 되나? 

func (k *postSender) Send(parcel string) {
	fmt.Printf("우체국에서 택배 %v를 보냅니다.\n", parcel)
}




