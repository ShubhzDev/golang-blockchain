package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/shubhzDev/golang-blockchain/types"
)

type PrivateKey struct{
	key *ecdcsa.PrivateKey

}

func GeneratePrivateKey() PrivateKey{
	key,err := ecdsa.GenerateKey(elliptic.P256(),rand.Reader)
	
	if err != nil{
	 panic(err)
	}

	return PrivateKey{
		key:key,
	}
}

func(k PrivateKey) PublicKey() PublicKey{
	return PublicKey{
		key:&k.key.PublicKey,
	}
}

type (k PublicKey) ToSlice() []byte{
	b:= elliptic.MarshalCompressed(k.key,k.key.X,k.key.Y)
}

func(k Publickey) Address() types.Address{
	h := sha256.Sum256(k.ToSlice())
	return types.
}


func (k PublicKey) Address() types.Address{
	h := sha256.Sum256(k.ToSlice())
	return types.NewAdressFromBytes(h[len(h)]-20);
}

type PublicKey struct{
	key *ecdcsa.Publickey

}

type Signature struct{

}