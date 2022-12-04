package main

import (
	"github.com/heejinshin/mini-fedex/fedex"
	"github.com/heejinshin/mini-fedex/koreaPost"
)

func SendBook(name string, sender *fedex.FedexSender) {  
	sender.Send(name)
}

func main() {
	sender := &koreaPost.PostSender{}
	SendBook("노인과 바다", sender)
	SendBook("보물섬", sender)

}