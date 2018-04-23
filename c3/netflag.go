package main

import(
	"fmt"
)

type Flags uint

const(
	FlagUp Flags = 1<<iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func Isup(v Flags) bool{
	return v&FlagUp == FlagUp
}

func TurnDown(v *Flags) {
	*v &^= 
}

func SetBroadcast(v *Flags){
	*v |= FlagBroadcast
}

func IsCast(v Flags) bool{
	return v&(FlagBroadcast|FlagMulticast) != 0
}

func main(){
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n",v ,Isup(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n",v,Isup(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n",v,Isup(v))
	fmt.Printf("%b %t\n",v,IsCast(v))
}