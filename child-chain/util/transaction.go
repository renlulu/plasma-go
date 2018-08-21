package util

const BLKNUM_OFFSET  = 1000000000
const TXINDEX_OFFSET = 10000


func EncodeUTXOId(blkNum uint64, txIndex uint64, oIndex uint64) uint64 {
	return blkNum * BLKNUM_OFFSET + txIndex * TXINDEX_OFFSET + oIndex
}

func DecodeUTXOId(id uint64) (uint64,uint64,uint64) {
	blkNum := id / BLKNUM_OFFSET
	txIndex := (id % BLKNUM_OFFSET) / TXINDEX_OFFSET
	oIndex := id - blkNum * BLKNUM_OFFSET - txIndex * TXINDEX_OFFSET
	return blkNum,txIndex,oIndex
}

func GetDepositTx() {

}