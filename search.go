package main

import (
	"fmt"
	"math/rand"
	"relaylib"
)

func main() {
	mysortarr := relaylib.InitArray(relaylib.MaxSession)
	GenSessPair(5, mysortarr)
	fmt.Println("------- Print Session Array---------")
	PrintSPTree()
	fmt.Println("---------Sorted Arrary----------")
	mysortarr.PrintTree()

}

func GenSessPair(count int, arr *relaylib.SRArr) {
	for i := 0; i < count; i++ {
		pos := relaylib.AddSP(genSession())
		arr.InsertRSNode(relaylib.GetSessPair(pos))
	}
}

func genSession() relaylib.SessionPair {
	ret := new(relaylib.SessionPair)
	ret.Did = genDid()
	return *ret
}

func PrintSPTree() {
	for i := int64(0); i < relaylib.GetCount(); i++ {
		fmt.Println(relaylib.GetSP(i).Did)
	}

}

func genDid() string {
	//num := rand.Int63n(281474976710655)
	num := rand.Int63n(1000000000)
	//return didPrefix + fmt.Sprintf("%012x", num)
	return fmt.Sprintf("%10d", num)
}
