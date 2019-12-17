package coinbase

import (
	"encoding/binary"
	"github.com/hacash/core/fields"
	"github.com/hacash/core/interfaces"
)

//
func ParseMinerPoolCoinbaseMessage(msgwords string, minernum uint32) [16]byte {
	var msg [16]byte
	copy(msg[0:11], []byte(msgwords)[0:11]) // minerpoolcn
	binary.BigEndian.PutUint32(msg[12:16], minernum)
	return msg
}

//
func UpdateCoinbaseMessageForMiner(tx interfaces.Transaction, minernum uint32) {
	newmsg := ParseMinerPoolCoinbaseMessage(string(tx.GetMessage()), minernum)
	tx.SetMessage(fields.TrimString16(string(newmsg[:])))
}

//
func UpdateBlockCoinbaseMessageForMiner(block interfaces.Block, minernum uint32) {
	UpdateCoinbaseMessageForMiner(block.GetTransactions()[0], minernum)
}
