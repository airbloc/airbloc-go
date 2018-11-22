// SHOULD RUN IN PARENT DIRECTORY
// go run contracts/generate_adapter.go
package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/airbloc/airbloc-go/contracts/utils"
	"github.com/valyala/fastjson"
)

const BuildOutput = "adapter"
const ContractDir = "contracts/build/contracts"

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

	contracts := make([]utils.Contract, len(fileNames))
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

		contracts[i] = utils.Contract{
			Name: string(v.GetStringBytes("contractName")),
			ABI:  v.Get("abi"),
			AST:  v.Get("ast"),
		}
	}

	for _, contract := range contracts {
		if contract.Name == "Migrations" {
			continue
		}

		tmp, err := utils.Bind(contract, "adapter")
		if err != nil {
			log.Printf("%+v", err)
			return
		}

		if err = ioutil.WriteFile(
			path.Join(BuildOutput, contract.Name+".go"),
			tmp,
			os.ModePerm,
		); err != nil {
			log.Println(err)
			return
		}
	}
}
