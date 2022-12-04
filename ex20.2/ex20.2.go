package main 

import "github.com/heejinshin/mini-fedex/fedex"


func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	sender := &fedex.FedexSender{}
	SendBook("어린왕자", sender)
	SendBook("오만과 편견", sender)
}