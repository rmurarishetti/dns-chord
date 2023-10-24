package message

import "fmt"

// Sample message structure. To be replaced with a struct for protobuff
type Message struct{
	Content string
	Type string // PING | SYNC | ACK
}

/*
***************************************
		UTILITY FUNCTIONS
***************************************	
*/
func (msg *Message) PrintContent(){
	fmt.Println("Message content:", msg.Content)
}