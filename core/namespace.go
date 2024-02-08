package core

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
	"log"
	"strconv"
)

var (
	DataUpload        = "DataUpload"
	CreateNamespace   = "CreateNamespace"
	UpdateNamespace   = "UpdateNamespace"
	TransferNamespace = "TransferNamespace"

	NamespaceAddr = common.HexToAddress("0x5200000000000000000000000000000000000001")

	NamespaceInfoAddr = common.HexToHash("0x00001")

	PublicNID = "0x00000000"
)

type NamespaceInfo struct {
	LastIndex int64 `json:"lastIndex"`
}

// use bitcoin address to generate daid
func generateDAID(fromAddress string, dataHex string, nonce uint64, nid string) string {
	nonceStr := strconv.Itoa(int(nonce))
	combinedData := fromAddress + dataHex + nonceStr + nid
	hash := sha3.Sum224([]byte(combinedData))
	hexStr := hex.EncodeToString(hash[:])
	log.Println("hexstr", hexStr)
	return "0x" + hexStr + nid[2:]
}

func nsIndexInt2Hex(nid int64) string {
	var hex string
	hexNum := strconv.FormatInt(nid, 16)
	hexPrefixLen := 8 - len(hexNum)
	for i := 0; i < hexPrefixLen; i++ {
		hex = hex + "0"
	}
	return "0x" + hex + hexNum
}

func Hex2Int(hexNum string) int64 {
	//hex format is 0x0000000d
	num, _ := strconv.ParseInt(hexNum[2:], 16, 10)
	return num
}
