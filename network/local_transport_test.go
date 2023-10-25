package network

import "testing"
import "github.com/stretch/testify/assert/"

func TestConnect(t *testing.T){
	tra := NewLocaltransport("A")
	trb := NewLocaltransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	assert.Equal(t,tra.peers[trb.addr],trb)
	assert.Equal(t,trb.peers[tra.addr],tra)

}

func TestSendMessage(t *testing.T){
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("hello wold")
	tra.SendMessage(trb.addr,msg)
	assert.Nil(t,tra.SendMessage(trb.addr,msg))

	rpc:= <- trb.Consume()
	assert.Euqal(t,rpc.Payload,msg)
	assert.Equal(t,rpc.From,tra.addr)
	
}