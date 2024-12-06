package internal

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitViper() {
	var config string = ""
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	workDir, _ := os.Getwd() //工作目录
	if config != "" {
		viper.SetConfigName(config) //配置文件文件名
		fmt.Println("您正在使用命令行的-c参数传递的值,config的文件名为" + config)
	} else {
		viper.SetConfigName("config.yaml") //配置文件文件名
	}
	viper.SetConfigType("yaml")            //配置文件的类型
	viper.AddConfigPath(workDir + "/conf") //工作目录
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("ReadInConfig err:" + err.Error())
	}
	fmt.Println("Viper 初始化成功")
}
