package controllers

import (
	"context"
	"log"
	"park/models"
	"park/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	//"github.com/ethereum/go-ethereum/log"
)

type SmartContractController struct {
	BaseController
}

// GetTest GetTest
// @Title Get
// @Summary 测试
// @Description 测试
// @Success 200 {object} controllers.ResponseData
// @router /test [get]
func (c *SmartContractController) GetTest() {

	//https://linea-sepolia.infura.io/v3/<YOUR-API-KEY>
	//wss://linea-sepolia.infura.io/ws/v3/<YOUR-API-KEY>
	//url:= "https://linea-sepolia.infura.io/v3/dccc894f6609403fa5e005bd1b228f88"
	url2 := "https://ethereum.matrix-net.tech/"
	conn, err := ethclient.Dial(url2)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// 交易哈希（合约部署交易的 TxHash）
	txHash := common.HexToHash("0x74f8b86e26b60f5883e4eb47158cd674863569fd38a2349bc4d942fdffd9c761")

	// 获取交易收据
	receipt, err := conn.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatalf("Failed to get transaction receipt: %v", err)
	}

	// 获取合约地址
	if receipt.ContractAddress.Hex() != "0x0000000000000000000000000000000000000000" {
		log.Println("合约地址:", receipt.ContractAddress.Hex())
	} else {
		log.Println("该交易未创建合约")
	}

	address := "0x05C0a833D158E97484F6887D42f92eC3807c4A49"
	parking_lot, err := models.NewParkingLot(common.HexToAddress(string(address)), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}
	//获取所有车位
	spots, err := parking_lot.GetAllParkingSpots(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to get status: %v", err)
	}
	utils.Println(spots)

	/**
	// eventCh := make(chan *StorageStatusUpdated)
	// opts := &bind.WatchOpts{}
	// go func() {
	// 	for event := range eventCh {
	// 		log.Printf("Status updated: %s", event.NewStatus)
	// 	}
	// }()
	// _, err = store.WatchStatusUpdated(opts, eventCh, nil)
	// if err != nil {
	// 	log.Fatalf("Failed to watch status updated: %v", err)
	// }
	// log.Println("Listening for events...")
	*/
	c.DoReturn(spots, err)
}
