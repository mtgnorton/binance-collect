package utility

import (
	"fmt"
	"github.com/gogf/gf/v2/os/genv"
	"os"
)

func LoadTestCfg(configFilename ...string) {
	envFile := "config/config.toml"

	if len(configFilename) != 0 {
		envFile = configFilename[0]
	}
	// 读取配置文件, 解决跑测试的时候找不到配置文件的问题，最多往上找5层目录
	for i := 0; i < 10; i++ {
		fmt.Println(envFile)
		if _, err := os.Stat(envFile); err == nil {
			genv.Set("GF_GCFG_FILE", envFile)
			break
		} else {
			envFile = "../" + envFile
		}
	}

}
