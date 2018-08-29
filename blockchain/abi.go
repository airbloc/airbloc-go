package blockchain

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"log"

	"github.com/ethereum/go-ethereum/common/compiler"
	"github.com/frostornge/ethornge/solc"
)

func getContractsFromABI(p string) (map[string]*compiler.Contract, error) {
	cs := make(map[string]*compiler.Contract)
	files, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		abi, err := os.Open(filepath.Join(p, f.Name()))
		if err != nil {
			return nil, err
		}

		decoder := json.NewDecoder(abi)

		c := new(compiler.Contract)
		if err = decoder.Decode(c); err != nil {
			return nil, err
		}
		filename := f.Name()
		cs["a:"+filename[0:len(filename)-len(filepath.Ext(filename))]+":a"] = c
	}

	return cs, nil
}

func GenerateBind(abiPath, bindPath string) error {
	cs, err := getContractsFromABI(abiPath)
	if err != nil {
		return err
	}

	for n := range cs {
		log.Println(n)
	}

	return solc.ExportBind(cs, bindPath)
}
