package configs

import (
	"github.com/spf13/viper"
	"llmTranslator/logHelper"
	"os"
)

func init() {
	//使用viper从config.yaml中读取配置信息
	//获取文件夹路径
	workDir, _ := os.Getwd()

	//设置配置文件名和路径
	viper.SetConfigName("setting")
	viper.AddConfigPath(workDir + "/configs")
	//设置配置文件类型
	viper.SetConfigType("json")

	_, err := os.Stat("configs")
	if os.IsNotExist(err) {
		err = os.Mkdir("configs", os.ModePerm)
		if err != nil {
			logHelper.Error("创建configs文件夹错误: %v", err)
		}
	}
	_, err = os.Stat("configs/setting.json")
	if os.IsNotExist(err) {
		createDefaultConfig()
	}

	//读取配置信息
	LoadSettingByFile()

	//如果无错误，显示配置文件读取成功
	logHelper.Info("config load success")

	_, err = os.Stat("tmp_img")
	if os.IsNotExist(err) {
		err = os.Mkdir("tmp_img", os.ModePerm)
		if err != nil {
			logHelper.Error("创建tmp_img文件夹错误: %v", err)
			return
		}
	}

}
