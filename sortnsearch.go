package main

import (
	"fmt"
	"lib"
	"math/rand"
)

const (
	maxSession = 100000
)

var (
	sID int
	//didPrefix = "1234567890" + "1234567890"
	didPrefix = ""
)

var sesArr [maxSession]SessionPair //A SessionPair Array

func main() {
	fmt.Println("start sorting ...")
	tree := make(RSTree, 0, maxSession)
	GenArray(&sesArr, &tree, 0)
	fmt.Println("==== UnSorted ==== ")
	PrintTree(tree)

	tree = tree[0:0]
	fmt.Println("=== Sorted ==== ")
	GenNSort(&sesArr, &tree, 0)
	PrintTree(tree)

	pic := len(tree) - 10

	if pic >= 0 {
		sDid := tree[5].keydid
		findPos := sort.Search(len(tree), func(i int) bool {
			return tree[i].keydid >= sDid
		})

		if findPos < len(tree) && tree[findPos].keydid == sDid {
			fmt.Println(sDid, " is Found at ", findPos)
		} else {
			fmt.Println(findPos, " is not Found . Shall add at", findPos)
		}
	} else {
		fmt.Println("pick element is out of boudary")
	}

	sDid := genDid()
	sess := new(SessionPair)
	sess.did = sDid
	sess.sessID = 10101

	findPos := sort.Search(len(tree), func(i int) bool {
		return tree[i].keydid >= sDid
	})

	if findPos < len(tree) && tree[findPos].keydid == sDid {
		fmt.Println(sDid, " is Found at ", findPos)
	} else {
		fmt.Println(sDid, " is not Found . Shall add at", findPos)
		node := make([]RSTNode, 1, 1)
		node[0].keydid = sDid
		node[0].valPrt = sess
		insertNode(&tree, findPos, node)
	}

	fmt.Println("find if insert ====== ")
	PrintTree(tree)
	findInsert := sort.Search(len(tree), func(i int) bool {
		return tree[findPos].keydid >= sDid
	})

	if findInsert < len(tree) && tree[findInsert].keydid == sDid {
		fmt.Println("insert success")
	} else {
		fmt.Println("insert failed")
	}

	delNode(&tree, findInsert)

	findDel := sort.Search(len(tree), func(i int) bool {
		return tree[findPos].keydid >= sDid
	})

	if findDel < len(tree) && tree[findDel].keydid == sDid {
		fmt.Println("delete failed")
	} else {
		fmt.Println("delete sucess")
	}
}

//GenArray gen count element of a SeesionPair address array
func GenArray(arr *[maxSession]SessionPair, tr *RSTree, count int) {
	for i := 0; i < count; i++ {
		arr[i] = GenSession()
		rsTNodePtr := GenRSTNode(arr[i].did, &arr[i])
		*tr = append(*tr, *rsTNodePtr)
	}
}

//GenNSort gen count element of a SeesionPair address array
func GenNSort(arr *[maxSession]SessionPair, tr *RSTree, count int) {
	for i := 0; i < count; i++ {
		//fmt.Println("GenArray")
		arr[i] = GenSession()
		rsTNodePtr := GenRSTNode(arr[i].did, &arr[i])
		*tr = append(*tr, *rsTNodePtr)
		sort.Sort(tr)
	}
}

//GenRSTNode generate a sorted tree node
func GenRSTNode(did string, sessPtr *SessionPair) *RSTNode {
	node := new(RSTNode)
	node.keydid = did
	node.valPrt = sessPtr
	return node
}

// GenSession generate a new SessionPair object and return its address
func GenSession() SessionPair {
	//fmt.Println("GenSession")
	sID = sID + 1
	ret := new(SessionPair)
	ret.sessID = sID
	ret.did = genDid()
	//fmt.Println("new Session", ret, "new Session content", *ret)
	return *ret
}

func genDid() string {
	//num := rand.Int63n(281474976710655)
	num := rand.Int63n(1000000000)
	//return didPrefix + fmt.Sprintf("%012x", num)
	return fmt.Sprintf("%10d", num)
}

func PrintTree(tr RSTree) {
	for i, c := range tr {
		fmt.Println(i, c.keydid)
	}
}
