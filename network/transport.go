package network

type RPC struct{
	From string
	Payload []byte
}

type Transport interface{
	Consume() <- chain RPC
	Connect(Transport) error
	SendMessage(NetAddr,[]btyte) errorAddr( NetAddr)
}