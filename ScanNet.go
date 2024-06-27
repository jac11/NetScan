package main

import (

    "fmt"
    "net"
    "flag"
    "time"
    "strconv"
    "sync"
    "strings"
)

var Reset = "\033[0m" 
var Red = "\033[31m" 
//var Green = "\033[32m" 
//var Yellow = "\033[33m" 
var Blue = "\033[34m" 
//var Magenta = "\033[35m" 
//var Cyan = "\033[36m" 
//var Gray = "\033[37m" 
//var White = "\033[97m"


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
   fmt.Println("Connection Succeeded Working On Port ", Port)
   Connect.Close()
}
func ScanRangePort(Domain string,Start string , End string){
	var WaitGroup sync.WaitGroup
	var mutex sync.Mutex
	StartInt,_ := strconv.Atoi(Start)
	EndInt,_ := strconv.Atoi(End)
    for Port := StartInt ; Port <= EndInt ;Port++{
        WaitGroup.Add(1)
        go func (Port int){
        	defer  WaitGroup.Done()
	    	DomainNet := net.JoinHostPort(Domain,strconv.Itoa(Port))
	        _, err := net.DialTimeout("tcp",DomainNet,3*time.Second)
	        mutex.Lock()
			defer mutex.Unlock()
	        if err != nil{
	           fmt.Printf("\rConnection Fail in Port %d", Port)
	           time.Sleep(10 * time.Millisecond)
	        }else{
	        	fmt.Printf("\rConnection Succeeded Working On Port %d\n", Port)
	        }
        }(Port)
    }
    WaitGroup.Wait() 
    fmt.Println()  
}
func Style (Styles Config) {
    var Banner string = `
███████  ██████  █████  ███    ██ ███    ██ ███████ ████████ 
██      ██      ██   ██ ████   ██ ████   ██ ██         ██    
███████ ██      ███████ ██ ██  ██ ██ ██  ██ █████      ██    
     ██ ██      ██   ██ ██  ██ ██ ██  ██ ██ ██         ██    
███████  ██████ ██   ██ ██   ████ ██   ████ ███████    ██    
                        @jacstory`+"\n"
                        
    fmt.Println(Blue+Banner+Reset)  
    if Styles.Port !="" && Styles.Domain !="" &&Styles.EndScan== "" && Styles.StartScan==""{
       fmt.Println("Staring Port     -----------| > ",Styles.Port)
       fmt.Println("Staring Domain   -----------| > ",Styles.Domain)
       fmt.Println(strings.Repeat("_", 40))
    }else if Styles.StartScan !="" && Styles.EndScan !=""{
       fmt.Println("Staring Port     -----------| > ",Styles.Port)
       fmt.Println("Staring Domain   -----------| > ",Styles.Domain)
       fmt.Println("Staring Port     -----------| > ",Styles.StartScan)
       fmt.Println("Ending  Port     -----------| > ",Styles.EndScan)
       fmt.Println("")
       fmt.Println(Red+strings.Repeat("_", 40)+Reset)
}
    }                 
    
func main(){
    var DataInfo Config
    flag.StringVar(&DataInfo.Port,"Port","80","default Port Scan")
    flag.StringVar(&DataInfo.Domain,"Domain","","IP/Domain To Scan")
    flag.StringVar(&DataInfo.StartScan,"StartScan","","Start Range Of Port Scan")
    flag.StringVar(&DataInfo.EndScan,"EndScan","","End Of Port Sacn")
    flag.Parse()
    Style(DataInfo)
    if DataInfo.Domain == ""{
	fmt.Println("Domain name or IP Not Valid")
	return
	}
     if DataInfo.StartScan != "" && DataInfo.EndScan !=""{
	ScanRangePort(DataInfo.Domain , DataInfo.StartScan , DataInfo.EndScan)
     }else{
    	ScanSinglPort(DataInfo.Domain,DataInfo.Port)
    }
}


