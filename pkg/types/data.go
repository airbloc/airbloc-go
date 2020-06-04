package types

type Data struct {
	UserId      ID     `json:"userId"`
	RowId       RowId  `json:"rowId"`
	CollectedAt Time   `json:"collecedAt"`
	Payload     string `json:"payload"`
}

type EncryptedData struct {
	UserId      ID     `json:"userId"`
	RowId       RowId  `json:"rowId"`
	Capsule     []byte `json:"capsule"`
	CollectedAt Time   `json:"collectedAt"`
	Payload     []byte `json:"payload"`
}
