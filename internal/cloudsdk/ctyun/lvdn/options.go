package lvdn

import (
	common "github.com/certimate-go/plugin-provider-ctcccloud-lvdn-deployer/internal/cloudsdk/ctyun/zz-shared-common"
)

func WithAkSk(ak, sk string) common.OptionsFunc {
	return common.WithAkSk(ak, sk)
}
