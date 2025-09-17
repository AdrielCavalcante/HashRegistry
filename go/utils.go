package hashregistry

import (
	"encoding/json"
	"io/ioutil"
)

// DeployInfo contém informações do contrato deployado
type DeployInfo struct {
	Address    string `json:"address"`
	Network    string `json:"network"`
	DeployedAt string `json:"deployedAt"`
}

// LoadContractAddress carrega o endereço do contrato do arquivo JSON
func LoadContractAddress(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	var deployInfo DeployInfo
	err = json.Unmarshal(data, &deployInfo)
	if err != nil {
		return "", err
	}

	return deployInfo.Address, nil
}
