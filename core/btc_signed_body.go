package core

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	verifier "github.com/bitonicnl/verify-signed-message/pkg"
)

type BTCSignedData struct {
	BitcoinAddress string `json:"bitcoinAddress"`
	MethodName     string `json:"methodName"`
	Nonce          uint64 `json:"nonce"`

	Payload DataPayload `json:"payload"`

	PublicKey  string `json:"publicKey"`
	StorageFee uint64 `json:"storageFee"`
}

func (bsd *BTCSignedData) verifySignature(signature string) (bool, error) {
	bz, _ := json.Marshal(bsd)
	hash := Sha256hash(string(bz))
	validSig, err := verifier.Verify(verifier.SignedMessage{
		Address:   bsd.BitcoinAddress,
		Message:   hash,
		Signature: signature,
	})
	return validSig, err
}

func NewBTCSignedData(data BTCData) BTCSignedData {
	return BTCSignedData{
		BitcoinAddress: data.BitcoinAddress,
		MethodName:     data.MethodName,
		Nonce:          data.Nonce,
		Payload:        data.Payload,
		PublicKey:      data.PublicKey,
		StorageFee:     data.StorageFee,
	}
}

type BTCSignedCreateNamespace struct {
	BitcoinAddress string `json:"bitcoinAddress"`
	MethodName     string `json:"methodName"`
	Nonce          uint64 `json:"nonce"`

	Payload CreateNamespacePayload `json:"payload"`

	PublicKey  string `json:"publicKey"`
	StorageFee uint64 `json:"storageFee"`
}

func NewBTCSignedCreateNamespace(cns BTCCreateNamespace) BTCSignedCreateNamespace {
	return BTCSignedCreateNamespace{
		BitcoinAddress: cns.BitcoinAddress,
		MethodName:     cns.MethodName,
		Nonce:          cns.Nonce,
		Payload:        cns.Payload,
		PublicKey:      cns.PublicKey,
		StorageFee:     cns.StorageFee,
	}
}

func (bscns *BTCSignedCreateNamespace) verifySignature(signature string) (bool, error) {
	bz, _ := json.Marshal(bscns)
	hash := Sha256hash(string(bz))
	validSig, err := verifier.Verify(verifier.SignedMessage{
		Address:   bscns.BitcoinAddress,
		Message:   hash,
		Signature: signature,
	})
	return validSig, err
}

type BTCSignedUpdateNamespace struct {
	BitcoinAddress string `json:"bitcoinAddress"`
	MethodName     string `json:"methodName"`
	Nonce          uint64 `json:"nonce"`

	Payload UpdateNamespacePayload `json:"payload"`

	PublicKey  string `json:"publicKey"`
	StorageFee uint64 `json:"storageFee"`
}

func (bsuns *BTCSignedUpdateNamespace) verifySignature(signature string) (bool, error) {
	bz, _ := json.Marshal(bsuns)
	hash := Sha256hash(string(bz))
	validSig, err := verifier.Verify(verifier.SignedMessage{
		Address:   bsuns.BitcoinAddress,
		Message:   hash,
		Signature: signature,
	})
	return validSig, err
}

func NewBTCSignedUpdateNamespace(uns BTCUpdateNamespace) BTCSignedUpdateNamespace {
	return BTCSignedUpdateNamespace{
		BitcoinAddress: uns.BitcoinAddress,
		MethodName:     uns.MethodName,
		Nonce:          uns.Nonce,
		Payload:        uns.Payload,
		PublicKey:      uns.PublicKey,
		StorageFee:     uns.StorageFee,
	}
}

func Sha256hash(rawdata string) string {
	data := []byte(rawdata)
	// Create a new SHA-256 hash
	hash := sha256.New()
	// Write data to the hash
	hash.Write(data)
	// Get the final hash sum
	hashSum := hash.Sum(nil)
	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hashSum)
	return hashString
}
