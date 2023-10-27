package core

import {
	"github.com/ShubhzDev/golang-blockchain/types"
	"io"
	"encoding/binary"
}

type Header struct{
	Version uint32
	PrevBlock types.Hash
	Timestamp uint64
	Height uint64
	Nonce uint64
}

func (h +Header) EncodeBinary(w io.writer) error{
	if err := binary.Write(w,binary.LittleEndian,&h.Version);
	err != null{
		return err
	}

	if err := binary.Write(w,binary.LittleEndian,&h.PrevBlock);
	err != null{
		return err
	}

	if err := binary.Write(w,binary.LittleEndian,&h.Timestamp);
	err != null{
		return err
	}

	if err := binary.Write(w,binary.LittleEndian,&h.Height);
	err != null{
		return err
	}

	if err := binary.Write(w,binary.LittleEndian,&h.Nonce);
	err != null{
		return err
	}
	
}

type Block struct{
	Header
	Transactions []Transaction
}