package main

import (
	"fmt"
	"lib/RelayLib"
	"math/rand"
)

var sessId int = 1
var didPrefix string = "1234567890" + "1234567890"

func main() {
	mysortarr := relaylib.InitSRArray(relaylib.MaxSession)
	GenSessPairs(10, &mysortarr)
	fmt.Println("------- Print Session Array---------")
	PrintSPTree()

	fmt.Println("---------Sorted Arrary----------")
	mysortarr.PrintTree()
	lookfor := (relaylib.GetSP(2)).GetDid()
	fmt.Println("Look for did = ", lookfor)
	sNode, err := mysortarr.FindUnusedSessByDid(lookfor)
	//spPtr, err := mysortarr.FindUnusedSessByDid("2313121")
	if err == nil {
		// found
		fmt.Println("found did=", sNode.SpPtr.GetDid())
	} else {
		fmt.Println("not found")
	}
	delnode := (relaylib.GetSP(5)).GetDid()
	fakesid := 5
	node := genANode(delnode, fakesid, false)
	fmt.Println("delete did = ", delnode, "sid", fakesid)
	delerr := mysortarr.DelRSNode(node)
	if delerr != nil {
		fmt.Println("delest error", delerr)
	}

}

func genANode(did string, sessId int, join bool) *relaylib.SessionPair {
	sp := new(relaylib.SessionPair)
	sp.SetDid(did)
	sp.SessionID = sessId
	sp.ClientIsJoin = join
	return sp
}

func GenSessPairs(count int, arr *relaylib.SRArr) {
	for i := 0; i < count; i++ {
		pos := relaylib.AddSP(genSession())
		sp1 := relaylib.GetSP(pos)
		spcopy := *sp1
		//sp1.ClientIsJoin = true
		spcopy.ClientIsJoin = false
		spcopy.SessionID = sessId
		sessId = sessId + 1
		pos2 := relaylib.AddSP(spcopy)
		ierr := arr.InsertRSNode(relaylib.GetSessPair(pos))
		if ierr != nil {
			fmt.Println("insert failed error = ", ierr)
		}
		ierr = arr.InsertRSNode(relaylib.GetSessPair(pos2))
		if ierr != nil {
			fmt.Println("insert failed error = ", ierr)
		}
	}
}

func genSession() relaylib.SessionPair {
	ret := new(relaylib.SessionPair)
	ret.SetDid(genDid())
	ret.SessionID = sessId
	ret.ClientIsJoin = true
	sessId = sessId + 1
	return *ret
}

func PrintSPTree() {
	for i := 0; i < relaylib.GetCount(); i++ {
		fmt.Println(relaylib.GetSP(i).GetDid())
	}
}

func genDid() string {
	//num := rand.Int63n(281474976710655)
	num := rand.Int31n(1000000000)

	return didPrefix + fmt.Sprintf("%012x ", num)
	//return fmt.Sprintf("%10d", num)
}
