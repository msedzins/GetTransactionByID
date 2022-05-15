package main

import (
	"bufio"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric-protos-go/peer"
	"github.com/hyperledger/fabric/protoutil"

	"github.com/hyperledger/fabric-protos-go/ledger/rwset"
	"github.com/hyperledger/fabric-protos-go/ledger/rwset/kvrwset"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	respBytes, err := hex.DecodeString(text[:len(text)-1])
	failOnError(err)

	//Response from the chaincode
	l1 := peer.ProcessedTransaction{}
	failOnError(proto.Unmarshal(respBytes, &l1))
	prettyPrint("Response from the chaincode", l1)

	l2 := common.Payload{}
	failOnError(proto.Unmarshal(l1.TransactionEnvelope.Payload, &l2))
	prettyPrint("Envelope payload", l2)

	l3 := peer.Transaction{}
	failOnError(proto.Unmarshal(l2.Data, &l3))
	prettyPrint("Data", l3)

	//We take first Action
	l4 := peer.ChaincodeActionPayload{}
	failOnError(proto.Unmarshal(l3.Actions[0].Payload, &l4))
	prettyPrint("Action", l4)

	//Here are input parameters to the chaincode - part 1
	l5, err := protoutil.UnmarshalChaincodeProposalPayload(l4.ChaincodeProposalPayload)
	failOnError(err)
	prettyPrint("Input parameters to the chaincode", l5)

	//Input parameters to the chaincode - part 2
	l6, err := protoutil.UnmarshalChaincodeInvocationSpec(l5.Input)
	failOnError(err)
	prettyPrint("Input parameters to the chaincode", l6)
	fmt.Printf("First argument extracted:\n%+v\n", string(l6.ChaincodeSpec.Input.Args[0]))

	//R/W set - part 1
	l7 := peer.ProposalResponsePayload{}
	failOnError(proto.Unmarshal(l4.Action.ProposalResponsePayload, &l7))
	prettyPrint("Read/Write set", l7)

	//R/W set - part 2
	l8 := peer.ChaincodeAction{}
	failOnError(proto.Unmarshal(l7.Extension, &l8))
	prettyPrint("Read/Write set", l8)

	//R/W set - part 3
	l9 := rwset.TxReadWriteSet{}
	failOnError(proto.Unmarshal(l8.Results, &l9))
	prettyPrint("Read/Write set", l9)

	//R/W set - part 4
	//Extract two elements from an array
	l10 := kvrwset.KVRWSet{}
	failOnError(proto.Unmarshal(l9.NsRwset[0].Rwset, &l10))
	prettyPrint("Read/Write set (0)", l10)

	failOnError(proto.Unmarshal(l9.NsRwset[1].Rwset, &l10))
	prettyPrint("Read/Write set (1)", l10)
	fmt.Printf("Value of Write extracted:\n%+v\n", string(l10.Writes[0].Value))
}

func failOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func prettyPrint(desc string, in any) {
	x, _ := json.MarshalIndent(in, "", "    ")
	fmt.Printf("%v:\n%v\n", desc, string(x))
	fmt.Println("*****")
}
