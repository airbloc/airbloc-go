package blockchain

import (
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"reflect"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/klaytn/klaytn/accounts/abi"
	"github.com/klaytn/klaytn/common"
	"github.com/pkg/errors"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var (
	// ContractList is filled by automatically generated contract binds in package `adapter`.
	contractList = make(map[string]ContractConstructor)
)

func AddContractConstructor(contractName string, contractConstructor ContractConstructor) {
	contractList[contractName] = contractConstructor
}

type contractManager struct {
	client     TxClient
	addrToName map[common.Address]string
	storage    map[reflect.Type]interface{}
}

func NewContractManager(client TxClient) *contractManager {
	return &contractManager{
		client:     client,
		storage:    make(map[reflect.Type]interface{}),
		addrToName: make(map[common.Address]string),
	}
}

func (cm *contractManager) Load(path string) error {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		resp, err := http.Get(path)
		if err != nil {
			return errors.Wrap(err, "failed to load deployment from url")
		}
		defer resp.Body.Close()
		return cm.load(resp.Body)
	}
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "failed to load deployment from file")
	}
	defer f.Close()
	return cm.load(f)
}

func (cm *contractManager) load(reader io.Reader) error {
	decoder := json.NewDecoder(reader)

	contracts := make(map[string]struct {
		Address   common.Address `json:"address"`
		TxHash    common.Hash    `json:"tx_hash"`
		CreatedAt *big.Int       `json:"created_at"`
		ABI       abi.ABI        `json:"abi"`
	})
	if err := decoder.Decode(&contracts); err != nil {
		return errors.Wrap(err, "contract manager: failed to decode json")
	}

	for name, info := range contracts {
		if _, ok := contractList[name]; !ok {
			continue
		}

		if name == "Consents" {
			methods := []abi.Method{
				info.ABI.Methods["consent"],
				info.ABI.Methods["consentMany"],
				info.ABI.Methods["consentByController"],
				info.ABI.Methods["consentManyByController"],
				info.ABI.Methods["modifyConsentByController"],
				info.ABI.Methods["modifyConsentManyByController"],
			}

			for _, method := range methods {
				for index, input := range method.Inputs {
					if input.Type.String() == "(uint8,string,bool)" {
						typ, err := abi.NewType("tuple", []abi.ArgumentMarshaling{
							{Name: "action", Type: "uint8"},
							{Name: "dataType", Type: "string"},
							{Name: "allowed", Type: "uint8"},
						})
						if err != nil {
							return errors.Wrap(err, "contract manager: failed to setup consent")
						}

						method.Inputs[index] = abi.Argument{
							Name: input.Name,
							Type: typ,
						}
					}
				}
			}

			d, _ := json.MarshalIndent(info.ABI.Methods["consent"].Inputs, "", "    ")
			log.Println(string(d))
		}

		cm.addrToName[info.Address] = name
		cm.SetContract(contractList[name](
			info.Address,
			info.TxHash,
			info.CreatedAt,
			info.ABI,
			cm.client,
		))
	}
	return nil
}

func (cm *contractManager) GetContract(c interface{}) interface{} {
	return cm.storage[reflect.ValueOf(c).Type()]
}

func (cm *contractManager) SetContract(c interface{}) {
	cm.storage[reflect.ValueOf(c).Type()] = c
}
