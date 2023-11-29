package types

import ("fmt"
"encoding/hex"
)

type Address [20]uint8


func(a Address) ToSlice() []byte{
	b := make([]byte,20)
	for i := 0; i < 20; i++{
		b[i] = a[i]
	}
	return b
}

func (a Address) String() string{
	return hex.EncodeToString(a.ToSlice())
}

func NewAdressFromBytes(b []byte) Address{
	if len(b) != 32{
		msg := fmt.Sprintf("given bytes with length %d should be 32",len(b))
		panic(msg)
	}

	var value [20]uint8
	for i := 0; i < 20; i++{
		value[i] = b[i]
	}

	return Address(value)
}