package structure

type Notify interface{
	SendMessage(msg string,priority string)(*Resp,error)
	// SendImage(img []byte)(*Resp,error)
}

// features
// multiple mailing system can be attached
// can attach their own label to each message optional
// automatic retry
// logging 