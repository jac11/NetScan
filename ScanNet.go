package main

import (

    "fmt"
    "net"
    "flag"
    "time"
    "strconv"
    "sync"
    "strings"
    "os"
)
var file *os.File
var OutFile =""
var IntNum = 0
var current_time = time.Now().Local()
var OPenPort int
var Reset = "\033[0m" 
var Red = "\033[31m" 
var Blue = "\033[34m" 

type Config struct{

    Port       string
    Domain     string
    StartScan  string 
    EndScan    string
    WriteFile  string
}
func ScanSinglPort(Domain string, Port string){
    DomainNet := net.JoinHostPort(Domain,Port)
    Connect, err := net.DialTimeout("tcp",DomainNet,3*time.Second)
    if err != nil{
       fmt.Println("🕵‍  Connection Fail          -----------| > ", Port, Red+" Close"+Reset)
       OutFile += fmt.Sprintf("🕵‍ Connection Fail     -----------| > %s",Port)+" Close\n"
       return
   }
   fmt.Println("🚀️ Connection Succeeded     -----------| > ",Port, Red+" Open\n "+Reset)
   OutFile += fmt.Sprintf("🚀️ Connection Succeeded     -----------| > %s",Port)+" Open\n"
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
            if err != nil {
                fmt.Printf("🕵‍  Connection Fail          -----------| > %d%s", Port, Red+" Close"+Reset)
                time.Sleep(10 * time.Millisecond)
                fmt.Print("\033[G\033[K") 
            }else{
                IntNum ++
                fmt.Printf("🚀️ Connection Succeeded     -----------| > %d%s",Port, Red+" Open\n "+Reset)
                OutFile += fmt.Sprintf("🚀️ Connection Succeeded     -----------| > %d",Port)+" Open\n"
                fmt.Print("\033[G\033[K")
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
███████  ██████ ██   ██ ██   ████ ██   ████ ███████    ██.Go    
                        @jacstory`+"\n"
                        
    fmt.Println(Blue+Banner+Reset)  
    if Styles.Port !="" && Styles.Domain !="" &&Styles.EndScan== "" && Styles.StartScan==""{
        fmt.Println("🌏 ScanDomain       -----------| > ",Styles.Domain)
        fmt.Println("🚨 Staring Port     -----------| > ",Styles.StartScan)
        fmt.Println("🕰  Strating Time    -----------| > ",current_time.Format("15:04:05"))
        fmt.Println(Red+strings.Repeat("_", 40)+Reset)
        fmt.Println("")
        if Styles.WriteFile !=""{
            OutFile += Banner+"\n"
            OutFile += fmt.Sprintf("🌏 ScanDomain       -----------| > "+Styles.Domain)+"\n"
            OutFile += fmt.Sprintf("🚨 Staring Port     -----------| > %s",Styles.StartScan)+"\n"
            OutFile += fmt.Sprintf("🕰  Strating Time    -----------| > %s",current_time.Format("15:04:05"))+"\n"
            OutFile += fmt.Sprintf("%s","")+"\n"
            OutFile += fmt.Sprintf("%s",strings.Repeat("_", 40))+"\n"
            OutFile += fmt.Sprintf("%s","")+"\n"
        }

    }else if Styles.StartScan !="" && Styles.EndScan !=""{
        fmt.Println("🌏 ScanDomain       -----------| > ",Styles.Domain)
        fmt.Println("🚨 Staring Port     -----------| > ",Styles.StartScan)
        fmt.Println("🎰️ Ending  Port     -----------| > ",Styles.EndScan)
        fmt.Println("🕰  Strating Time    -----------| > ",current_time.Format("15:04:05"))
        fmt.Println("")
        fmt.Println(Red+strings.Repeat("_", 40)+Reset)
        fmt.Println("")
   
        if Styles.WriteFile !=""{
            OutFile += Banner+"\n"
            OutFile += fmt.Sprintf("🌏 ScanDomain       -----------| > %s",Styles.Domain)+"\n"
            OutFile += fmt.Sprintf("🚨 Staring Port     -----------| > %s",Styles.StartScan)+"\n"
            OutFile += fmt.Sprintf("🎰️ Ending  Port     -----------| > %s",Styles.EndScan)+"\n"
            OutFile += fmt.Sprintf("🕰  Strating Time    -----------| > %s",current_time.Format("15:04:05"))+"\n"
            OutFile += fmt.Sprintf("%s","")+"\n"
            OutFile += fmt.Sprintf("%s",strings.Repeat("_", 40))+"\n"
            OutFile += fmt.Sprintf("%s","")+"\n"
           
        }
    }     
}                    
func ResaltScan(Conut Config){
    fmt.Println(Red+strings.Repeat("_", 40)+Reset)
    fmt.Println("")
    TimeEnd :=  time.Now().Local()
    AllTime := TimeEnd.Sub(current_time)
    CountPort1 ,_ := strconv.Atoi(Conut.EndScan)
    CountPort2 ,_ := strconv.Atoi(Conut.StartScan)
    fmt.Println("🧭 EndTime           -----------| > ",TimeEnd.Format("15:04:05"))
    fmt.Println("⏳ Scan Time         -----------| > ", AllTime )
    fmt.Println("🎯 Port Conut        -----------| > ", CountPort1 - CountPort2+1 )
    fmt.Println("🪲️ Ports Live        -----------| > ", IntNum)
    fmt.Println("🐞 Close Ports       -----------| > ", CountPort1 - IntNum )
    if Conut.WriteFile !=""{
        OutFile += fmt.Sprintf("%s",strings.Repeat("_", 40))+"\n\n"
        OutFile += fmt.Sprintf("🧭 EndTime           -----------| > %s",TimeEnd.Format("15:04:05"))+"\n"
        OutFile += fmt.Sprintf("⏳ Scan Time         -----------| > %s", AllTime )+"\n"
        OutFile += fmt.Sprintf("🎯 Port Conut        -----------| > %d", CountPort1 - CountPort2+1 )+"\n"
        OutFile += fmt.Sprintf("🪲️ Ports Live        -----------| > %d", IntNum)+"\n"
        OutFile += fmt.Sprintf("🐞 Close Ports       -----------| > %d", CountPort1 - IntNum )+"\n"
    }
    writeFile(OutFile,Conut)
}  

func writeFile(OutPut string , DataInfo Config){
    var file *os.File
    var err error

    file, err = os.OpenFile(DataInfo.WriteFile, os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Error opening file:", err)
            return
    }
    defer file.Close()
    if OutFile != ""{
        file.WriteString( OutFile  + "\n")
    }
} 

func main(){
    var DataInfo Config
    flag.StringVar(&DataInfo.Port,"Port","80","default Port Scan")
    flag.StringVar(&DataInfo.Domain,"Domain","","IP/Domain To Scan")
    flag.StringVar(&DataInfo.StartScan,"StartScan","","Start Range Of Port Scan")
    flag.StringVar(&DataInfo.EndScan,"EndScan","","End Of Port Sacn")
    flag.StringVar(&DataInfo.WriteFile , "WriteFile" ,"","Write STdout To File")
    flag.Parse()
    if DataInfo.WriteFile !=""{
        writeFile(OutFile,DataInfo)
    }
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
    ResaltScan(DataInfo)
}
