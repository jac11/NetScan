package main

import (

    "fmt"
    "net"
    "flag"
    "time"
    "strconv"
)


type Config struct{

    Port       string
    Domain     string
    StartScan  string 
    EndScan    string
}

func ScanSinglPort(Domain string,Port string){
    DomainNet := net.JoinHostPort(Domain,Port)
    Connect, err := net.DialTimeout("tcp",DomainNet,3*time.Second)
    if err != nil{
       fmt.Println("Connection Fail in Port ", Port)
       return
   }
   fmt.Println("Connection Suessed Working On Port" ,Port)
   Connect.Close()
}
func ScanRangePort(Domain string,Start string , End string){
	StartInt,_ := strconv.Atoi(Start)
	EndInt,_ := strconv.Atoi(End)
    for Port := StartInt ; Port <= EndInt ;Port++{
    	DomainNet := net.JoinHostPort(Domain,strconv.Itoa(Port))
        _, err := net.DialTimeout("tcp",DomainNet,3*time.Second)
        if err != nil{
           fmt.Println("Connection Fail in Port ", Port)

        }else{
        	 fmt.Println("Connection Suessed Working On Port" ,Port)
        }
    }    
}
func main(){
    var DataInfo Config
    flag.StringVar(&DataInfo.Port,"Port","80","default Port Scan")
	flag.StringVar(&DataInfo.Domain,"Domain","","IP/Domain To Scan")
	flag.StringVar(&DataInfo.StartScan,"StartScan","","Start Range Of Port Scan")
	flag.StringVar(&DataInfo.EndScan,"EndScan","","End Of Port Sacn")
	flag.Parse()
	if DataInfo.Domain == ""{
		fmt.Println("Domain name or ip Not Valed")
		return
	}
	if DataInfo.StartScan != "" && DataInfo.EndScan !=""{
	    ScanRangePort(DataInfo.Domain , DataInfo.StartScan , DataInfo.EndScan)
    }else{
    	ScanSinglPort(DataInfo.Domain,DataInfo.Port)
    }
}