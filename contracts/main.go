package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

const BuildOutput = "contracts/build/out"
const ContractDir = "contracts/build/contracts"

type Contract struct {
	Name string      `json:"contractName"`
	Abi  interface{} `json:"abi"`
	Bin  string      `json:"deployedBytecode"`
}

func main() {
	fileInfos, err := ioutil.ReadDir(ContractDir)
	if err != nil {
		log.Println(err)
		return
	}

	fileNames, i := make([]string, len(fileInfos)), 0
	for _, fileInfo := range fileInfos {
		if fileInfo != nil && !fileInfo.IsDir() {
			fileNames[i] = fileInfo.Name()
			i++
		}
	}

	contracts := make([]Contract, len(fileNames))
	for i, fileName := range fileNames {
		file, err := os.Open(ContractDir + "/" + fileName)
		if err != nil {
			log.Println(err)
			return
		}

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&contracts[i])
		file.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}

	for _, contract := range contracts {
		abi, err := json.Marshal(contract.Abi)
		if err != nil {
			log.Println(err)
			return
		}

		if contract.Name == "Migrations" {
			continue
		}

		outPath := "adapter/" + contract.Name + ".go"

		tmp, err := bind.Bind(
			[]string{contract.Name},
			[]string{string(abi)},
			[]string{contract.Bin},
			"adapter",
			bind.LangGo,
		)
		if err != nil {
			log.Println(err)
			return
		}

		if err = ioutil.WriteFile(
			outPath,
			[]byte(tmp),
			os.ModePerm,
		); err != nil {
			log.Println(err)
			return
		}
	}
}
