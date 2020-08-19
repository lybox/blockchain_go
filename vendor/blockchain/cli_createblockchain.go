package blockchain

import (
	"fmt"
	"log"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	// 校验钱包地址是否存在，如果没有，可以先创建钱包以便获取地址领取创建区块的奖励
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := CreateBlockchain(address, nodeID)
	defer bc.db.Close()

	UTXOSet := UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
