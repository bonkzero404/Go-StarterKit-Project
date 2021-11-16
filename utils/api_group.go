package utils

import (
	"fmt"
	"go-boilerplate-clean-arch/config"
)

func SetupApiGroup() string {
	str := fmt.Sprintf("/%s/%s", config.Config("API_WRAP"), config.Config("API_VERSION"))
	return str
}
