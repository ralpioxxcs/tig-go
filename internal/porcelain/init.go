package porcelain

import "tig/internal/config"

type InitParam struct {
	WorkingCopy string
	Config      config.Config
}

func Init(param InitParam) error {

	return nil

}
