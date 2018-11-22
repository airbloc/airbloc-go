// SHOULD RUN IN PARENT DIRECTORY
// go run contracts/generate_adapter.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/airbloc/airbloc-go/contracts/utils"
	"github.com/valyala/fastjson"
)

const BuildOutput = "adapter"
const ContractDir = "contracts/build/contracts"

type Contract struct {
	Name string
	ABI  *fastjson.Object
	AST  *fastjson.Object
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
		data, err := ioutil.ReadFile(ContractDir + "/" + fileName)
		if err != nil {
			log.Println(err)
			return
		}

		var p fastjson.Parser
		v, err := p.ParseBytes(data)
		if err != nil {
			log.Println(err)
			return
		}

		contracts[i] = Contract{
			Name: string(v.GetStringBytes("contractName")),
			ABI:  v.GetObject("abi"),
			AST:  v.GetObject("ast"),
		}
	}

	for _, contract := range contracts {
		abi, err := json.Marshal(contract.ABI)
		if err != nil {
			log.Println(err)
			return
		}

		if contract.Name == "Migrations" {
			continue
		}

		outPath := path.Join(BuildOutput, contract.Name+".go")
		fmt.Println(string(abi))

		tmp, err := utils.Bind(
			[]string{contract.Name},
			[]string{string(abi)},
			[]string{""},
			"adapter",
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
