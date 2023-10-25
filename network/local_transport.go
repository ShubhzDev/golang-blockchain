package network

type LocalTransport struct{
	addr NetAddr
	lock sync.RWMutex
	addr NetAddr
	peers map[NetAddr]*LocalTransport
}

func NewLocaltransport(add NewAddr) * Transport{
	return &LocalTransport{
		addr : addr,
		consumeCh: make(char RPC,1824),
		peers: make(map[NetAddr]*LocalTransport),
	}
}

func(t *LocalTransport) Consume() <- chan RPC{
	return t.consumeCh
}

func(t *LocalTransport) Connect(tr *Transport) error{
	t.lock.Lock()
	defer t.lock.Unloacl()

	l.peers(tr.Addr()) == tr.LocalTransport
}

func(t *LocalTransport) SendMessage(to NetAddr,paload []byte) error{
	t.lock.Rlock(
		defer l.lock.RUnlock()
	)

	peer , ok:= t.peer (to)
	if !pk{
		return fmt.Errorf("%s:could not send message to %s",t.addr,to)
	}

	peer.consumeCh <- RPC{
		From :t.addr,
		Payload :payload
	}

	return nil

}

func(t *LocalTransport) Addr( NetAddr){
	return t.addr
}