package core

type BTCTx struct {
	Signature          string             `json:"signature"`
	BtcData            BTCData            `json:"btcData"`
	BtcCreateNamespace BTCCreateNamespace `json:"btcCreateNamespace"`
	BtcUpdateNamespace BTCUpdateNamespace `json:"btcUpdateNamespace"`
	EthUserAddress     string             `json:"ethUserAddress"`
	MethodName         string             `json:"methodName"`
	IsAgentTx          bool               `json:"isAgentTx"`
}

type BTCData struct {
	DataID         string `json:"dataID"`
	BitcoinAddress string `json:"bitcoinAddress"`
	MethodName     string `json:"methodName"`
	Nonce          uint64 `json:"nonce"`

	Payload DataPayload `json:"payload"`

	PublicKey  string `json:"publicKey"`
	StorageFee uint64 `json:"storageFee"`
}

func (bd *BTCData) ComputeHash() string {
	return ""
}

type DataPayload struct {
	RawData     string            `json:"data"`
	Labels      map[string]string `json:"labels"`
	NamespaceID string            `json:"namespaceID"`
}

type BTCCreateNamespace struct {
	BitcoinAddress string `json:"bitcoinAddress"`
	MethodName     string `json:"methodName"`
	Nonce          uint64 `json:"nonce"`

	Payload CreateNamespacePayload `json:"payload"`

	PublicKey  string `json:"publicKey"`
	StorageFee uint64 `json:"storageFee"`
}

func (bd *BTCCreateNamespace) ComputeHash() string {
	return ""
}

type BTCUpdateNamespace struct {
	BitcoinAddress string `json:"bitcoinAddress"`
	MethodName     string `json:"methodName"`
	Nonce          uint64 `json:"nonce"`

	Payload UpdateNamespacePayload `json:"payload"`

	PublicKey  string `json:"publicKey"`
	StorageFee uint64 `json:"storageFee"`
}

func (bd *BTCUpdateNamespace) ComputeHash() string {
	return ""
}

type CreateNamespacePayload struct {
	Admins     []string `json:"admins"`
	Name       string   `json:"name"`
	Owner      string   `json:"owner"` //btc address
	Permission string   `json:"permission"`
}

type UpdateNamespacePayload struct {
	Admins      []string `json:"admins"`
	Name        string   `json:"name"`
	NamespaceID string   `json:"namespaceID"`
	Permission  string   `json:"permission"`
}

type NamespaceChainModel struct {
	Admins      []string `json:"admins"`
	Name        string   `json:"name"`
	Owner       string   `json:"owner"` //btc address
	NamespaceID string   `json:"namespaceID"`
	Permission  string   `json:"permission"`
}

func NewNamespaceChainByCns(cns CreateNamespacePayload, nid string) NamespaceChainModel {
	return NamespaceChainModel{
		Admins:      cns.Admins,
		Name:        cns.Name,
		Owner:       cns.Owner,
		NamespaceID: nid,
		Permission:  cns.Permission,
	}
}

func NewNamespaceChainByUns(uns UpdateNamespacePayload, owner string) NamespaceChainModel {
	return NamespaceChainModel{
		Admins:      uns.Admins,
		Name:        uns.Name,
		Owner:       owner,
		NamespaceID: uns.NamespaceID,
		Permission:  uns.Permission,
	}
}
