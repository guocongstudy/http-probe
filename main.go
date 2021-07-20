package main

import (
	"flag"
	"http-probe/config"
	"log"

	"http-probe/http"
)

var (
	configFile string
)
func main(){
//传入配置文件路径
 flag.StringVar(&configFile,"c","simple_http_probe.yaml","config file path")
//解析yaml
 conf,err :=config.LoadFile(configFile)
 if err !=nil{
 	log.Printf("[load.yaml.error][err:%v]",err)
	 return
 }
	log.Printf("配置是：%v",conf)
 //启动gin
 go http.StartGin(conf)
 select {}

}
