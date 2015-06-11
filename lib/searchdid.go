package searchdid

import (
	"fmt"
	"sort"
)

//SRNode is element for arrary used for search did and its SessionPair
type SRNode struct {
	keyDid  string
	valSess *SessionPair //A session Pair pointer, point to actual array
}

//SRArr is defined to implement sort interface
type SRArr []SRNode

func (arr SRArr) Len() int {
	return len(arr)
}

func (arr SRArr) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr SRArr) Less(i, j int) bool {
	return arr[i].keyDid < arr[j].keyDid
}

//InsertRSNode insert a new SRNode according to given SessionPair to an sorted array
func InsertRSNode(arr *SRArr, idx int, sp *SessionPair) {

	did := sp.did

	findPos := sort.Search(len(arr), func(i int) bool {
		return arr[i].keydid >= sDid
	})

	if findPos < len(tree) && tree[findPos].keydid == sDid {
		fmt.Println(sDid, " Given session has exist in sorted tree at pos =  ", findPos)
	} else {
		fmt.Println(sDid, " is not Found . will be add at index=", findPos)
		node := make([]SRNode, 1, 1)
		//	node[0].keyDid = *sp.
		//	node[0].valSess = sess
		insertNode(&tree, findPos, node)
	}
}

func insertNode(tr *SRArr, idx int, node []SRNode) {
	rightLen := len(*tr) - idx + 1
	iNode := make([]SRNode, 1, rightLen)
	copy(iNode, node)
	//iNode = append(iNode, (*tr)[idx:]...)
	*tr = append((*tr)[:idx], append(iNode, (*tr)[idx:]...)...)
	//PrintTree(*tr)
}

func delNode(tr *SRArr, idx int) {
	copy((*tr)[idx:], (*tr)[idx+1:])
	(*tr)[len(*tr)-1] = RSTNode{} // prevent memory leak, this shall be verified in the future
	*tr = (*tr)[:len(*tr)-1]
	PrintTree(*tr)
}
