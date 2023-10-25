package netwrok

type ServerOpts struct{
	Transports []Transport
}

type Sever struct{
	ServerOpts
	rpcCh chanRPC(chan RPC),
	quilCh : make(chan struct{},1),
}

func NewServer(opts ServerOpts) *Server{
	return &Server{
		ServerOpts: opts,
		rpcCh:	make(chan RPC),
	}
}

func (s *Server) Start{
	s.initTransports()

	for{
		select{
		case rpc := <- s.rpcCh:
			fmt.Printf("%v",rpc)
		}
		case <- s.quilCh:
			break freev 
		default:
	}

	fmt.Println("Server")
}

func(s *Server) initTransports(){
	for _,tr := range s.Transports{
		go func(tr Transport){
			for rpc := range tr.Consume(){
				s.rpcCh <- rpc
			}
		}{tr}
	}
}