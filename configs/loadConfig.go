package configs

import (
	"encoding/json"
	"github.com/spf13/viper"
	"llmTranslator/logHelper"
	"os"
)

func init() {
	//使用viper从config.yaml中读取配置信息
	//获取文件夹路径
	workDir, _ := os.Getwd()

	_, err := os.Stat("configs")
	if os.IsNotExist(err) {
		err = os.Mkdir("configs", os.ModePerm)
		if err != nil {
			logHelper.Error("创建configs文件夹错误: %v", err)
			logHelper.WriteLog("创建configs文件夹错误: %v", err)
		}
	}
	_, err = os.Stat("configs/setting.json")
	if os.IsNotExist(err) {
		createDefaultConfig()
	}

	//设置配置文件名和路径
	viper.SetConfigName("setting")
	viper.AddConfigPath(workDir + "/configs")
	//设置配置文件类型
	viper.SetConfigType("json")
	//读取配置信息
	err = viper.ReadInConfig()
	//处理错误
	if err != nil {
		logHelper.Debug("config load error: %v", err)
		logHelper.WriteLog("config load error: %v", err)
	}

	//如果无错误，显示配置文件读取成功
	logHelper.Info("config load success")

	_, err = os.Stat("tmp_img")
	if os.IsNotExist(err) {
		err = os.Mkdir("tmp_img", os.ModePerm)
		if err != nil {
			logHelper.Error("创建tmp_img文件夹错误: %v", err)
			logHelper.WriteLog("创建tmp_img文件夹错误: %v", err)
			return
		}
	}

}

func createDefaultConfig() {
	file, err := os.OpenFile("configs/setting.json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logHelper.Error("创建配置文件错误: %v", err)
		logHelper.WriteLog("创建配置文件错误: %v", err)
		return
	}
	defer file.Close()

	config := getDefaultConfig()

	defaultConfig, err := json.Marshal(config)
	if err != nil {
		logHelper.Error("创建默认配置时JSON序列化失败: %v", err)
		logHelper.WriteLog("创建默认配置时JSON序列化失败: %v", err)
		return
	}

	_, err = file.Write(defaultConfig)
	if err != nil {
		logHelper.Error("创建默认配置时写入配置文件错误: %v", err)
		logHelper.WriteLog("创建默认配置时写入配置文件错误: %v", err)
	}
}
