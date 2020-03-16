/*
@Time : 2020/3/16 3:18 下午
@Author : sunmoon
@File : boot.go
@Software: GoLand
*/
package boot

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcmd"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
)

const (
	currentConfigPath = "../../config"
	configFileName    = "config.toml"
)

func getOption(parser *gcmd.Parser, name string, defValue ...string) (result string) {
	result = parser.GetOpt(name)
	if result == "" && len(defValue) > 0 {
		result = defValue[0]
	}
	return
}

func init() {
	//设置config的目录
	parser, err := gcmd.Parse(g.MapStrBool{
		"c,cd": true,
	})
	if err != nil {
		glog.Line().Println(err.Error())
	}
	path := getOption(parser, "c", currentConfigPath)
	configFile := path + gfile.Separator + configFileName
	if !gfile.Exists(configFile) {
		glog.Line().Println("无效的配置文件：" + configFile)
	}
	fmt.Println("正在以 " + configFile + " 文件启动")
	err = g.Cfg().SetPath(currentConfigPath)
	if err != nil {
		fmt.Println("set cfg path error:"+err.Error())
	}
	//开启数据库调试模式
	db := g.DB()
	db.SetDebug(true)
	//日志配置
	logPath := g.Cfg().GetString("app.logPath")
	if len(logPath) == 0 {
		panic("请指定日志目录")
	}
	_ = glog.SetPath(logPath)
	glog.SetStdoutPrint(true)
	s := g.Server()
	s.SetLogPath(logPath)
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(false)
}
