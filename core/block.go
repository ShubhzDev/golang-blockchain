package core

import (
	"github.com/ShubhzDev/golang-blockchain/types"
	"io"
	"encoding/binary"
)

type Header struct{
	Version uint32
	PrevBlock types.Hash
	Timestamp uint64
	Height uint64
	Nonce uint64
}

func (h *Header) EncodeBinary(w io.Writer) error{
	if err := binary.Write(w,binary.LittleEndian,&h.Version);
	err != nil{
		return err
	}

	if err := binary.Write(w,binary.LittleEndian,&h.PrevBlock);
	err != nil{
		return err
	}

	if err := binary.Write(w,binary.LittleEndian,&h.Timestamp);
	err != nil{
		return err
	}

	if err := binary.Write(w,binary.LittleEndian,&h.Height);
	err != nil{
		return err
	}

	return binary.Write(w,binary.LittleEndian,&h.Nonce);
}

type Block struct{
	Header
	Transactions []Transaction
}