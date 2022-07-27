package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Cfg 全局变量，用来保存程序的所有配置信息
var Cfg = new(Config)

func Init(filepath string) error {
	//// 基本上搭配远程配置中心使用的，告诉 viper 当前的数据使用什么格式来解析
	////viper.SetConfigType("yaml")
	//
	//viper.SetConfigName("config") // 配置文件名称(无扩展名)
	//viper.AddConfigPath(".")      // 查找配置文件所在的路径（注意：相对于当前执行命令是所在目录的相对路径）
	//viper.AddConfigPath("./conf") // 多次调用以添加多个搜索路径
	//viper.AddConfigPath("$HOME/.appname") // 还可以在工作目录中查找配置

	viper.SetConfigFile(filepath) // 指定配置文件

	// 查找并读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		// 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig() failed,err:%v \n", err)
		return err
	}

	if err := viper.Unmarshal(Cfg); err != nil {
		fmt.Printf("viper.Unmarshal failed,err:%v \n", err)
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("配置文件修改了", e.Name)
		if err := viper.Unmarshal(Cfg); err != nil {
			fmt.Printf("viper.Unmarshal failed,err:%v \n", err)
		}
	})
	return nil
}
