package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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
		abi, err := json.MarshalIndent(contract.Abi, "", "    ")
		if err != nil {
			log.Println(err)
			return
		}

		if contract.Name == "Migrations" {
			continue
		}

		outPath := "bind/" + contract.Name + ".go"
		abiPath := BuildOutput + "/" + contract.Name + ".abi"
		binPath := BuildOutput + "/" + contract.Name + ".bin"

		err = ioutil.WriteFile(abiPath, abi, os.ModePerm)
		if err != nil {
			log.Println(err)
			return
		}

		err = ioutil.WriteFile(binPath, []byte(contract.Bin), os.ModePerm)
		if err != nil {
			log.Println(err)
			return
		}

		if err = exec.Command(
			"abigen",
			"--abi", abiPath,
			"--bin", binPath,
			"--pkg", "contract",
			"--type", contract.Name,
			"--out", outPath,
		).Run(); err != nil {
			log.Println(err)
			return
		}
	}
}
