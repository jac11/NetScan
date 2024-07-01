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
    "os/exec"
    "regexp"
)
var file *os.File
var OPenPort int
var Reset   = "\033[0m" 
var Red     = "\033[31m" 
var Blue    = "\033[34m" 
var Cyan    = "\033[36m" 
var White   = "\033[97m"
var Yellow  = "\033[33m" 
var OutFile =""
var IntNum  = 0
var current_time = time.Now().Local()
type Config struct{

    Port       string
    Domain     string
    StartScan  string 
    EndScan    string
    WriteFile  string
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
        fmt.Println("🚨 Scan Port        -----------| > ",Styles.Port)
        fmt.Println("🕰  Strating Time    -----------| > ",current_time.Format("15:04:05"))
        fmt.Println(Red+strings.Repeat("_", 40)+Reset)
        fmt.Println("")
        if Styles.WriteFile !=""{
            OutFile += Banner+"\n"
            OutFile += fmt.Sprintf("ScanDomain       -----------| > "+Styles.Domain)+"\n"
            OutFile += fmt.Sprintf("Scan Port        -----------| > %s",Styles.Port)+"\n"
            OutFile += fmt.Sprintf("Strating Time    -----------| > %s",current_time.Format("15:04:05"))+"\n"
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
            OutFile += fmt.Sprintf("ScanDomain       -----------| > %s",Styles.Domain)+"\n"
            OutFile += fmt.Sprintf("Staring Port     -----------| > %s",Styles.StartScan)+"\n"
            OutFile += fmt.Sprintf("Ending  Port     -----------| > %s",Styles.EndScan)+"\n"
            OutFile += fmt.Sprintf("Strating Time    -----------| > %s",current_time.Format("15:04:05"))+"\n"
            OutFile += fmt.Sprintf("%s","")+"\n"
            OutFile += fmt.Sprintf("%s",strings.Repeat("_", 40))+"\n"
            OutFile += fmt.Sprintf("%s","")+"\n"
           
        }
    }     
}                       
func PingHost(Domain string)(string,string){
    cmd := exec.Command("ping", "-c", "1",Domain)
    output, err := cmd.Output()
    if err != nil {
        fmt.Println("⛔️ Status-1\t\t    -----------| > Executing Internet Error")
        fmt.Println("⛔️ Status-2\t\t    -----------| > Use Domain Without http/https ")
        os.Exit(0)
    }
    re := regexp.MustCompile(`ttl=(\d+)`)
    matches := re.FindStringSubmatch(string(output))
    IPGet := regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`)
    matches2 := IPGet.FindStringSubmatch(string(output))[1]
    return matches[1],matches2
}
func ScanSinglPort(Domain string, Port string){
   
    DomainNet := net.JoinHostPort(Domain,Port)
    Connect, err := net.DialTimeout("tcp",DomainNet,3*time.Second)
    if err != nil{
       fmt.Println("🕵‍  Connection Fail          -----------| > ", Port, Red+" Close"+Reset)
       OutFile += fmt.Sprintf("Connection Fail     -----------| > %s",Port)+" Close\n"
       return
   }
    for PortService , Service  := range (myMap){
        if Port == PortService {
            fmt.Println("🚀️ Connection Succeeded     -----------| > ",Port, Red+" Open "+Reset+ Cyan +  Service + Reset)
            OutFile += fmt.Sprintf("Connection Succeeded     -----------| > %s",Port)+" Open "+Service+"\n"
            Connect.Close()
            
        }
    } 
   
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
                fmt.Printf("🕵‍  Connection Fail          -----------| > %-5d%s", Port, Red+" Close"+Reset)
                time.Sleep(10 * time.Millisecond)
                fmt.Print("\033[G\033[K") 
            }else{
                IntNum ++
                for PortService , Service  := range (myMap){
                   if strconv.Itoa(Port) == PortService {
                       fmt.Printf("🚀️ Connection Succeeded     -----------| > %-4d %s Open %s %s%s %s\n", Port, Red, Reset, Cyan, Service, Reset)
                       OutFile += fmt.Sprintf("Connection Succeeded     -----------| > %-4d",Port)+" Open "+Service+"\n"
                       fmt.Print("\033[G\033[K")
                    }
                }
            }
        }(Port)
    }
    
    WaitGroup.Wait() 
    fmt.Println()
}

func ResaltScan(Conut Config){
    var Value = ""
    TTL,IP := PingHost(Conut.Domain)
    for TTLValue,TTLMessage := range(TTLOS){ 
        if TTL == strconv.Itoa(TTLValue){
            Value = TTLMessage
            break
        }else{
         TTLMessage = "UnKwon OS-Guess Linux"
         Value = TTLMessage
       }
    }
    fmt.Println(Red+strings.Repeat("_", 40)+Reset)
    fmt.Println("")
    TimeEnd :=  time.Now().Local()
    AllTime := TimeEnd.Sub(current_time)
    CountPort1 ,_ := strconv.Atoi(Conut.EndScan)
    CountPort2 ,_ := strconv.Atoi(Conut.StartScan)
    fmt.Println("🖥  Domain IP         -----------| > ",White+IP+Reset)
    fmt.Println("💡 Guess OS          -----------| > ",Yellow+ Value +Reset)
    fmt.Println("🧭 EndTime           -----------| > ",TimeEnd.Format("15:04:05"))
    fmt.Println("⏳ Scan Time         -----------| > ", AllTime )
    fmt.Println("🎯 Port Conut        -----------| > ", CountPort1 - CountPort2+1 )
    fmt.Println("🪲️ Ports Live        -----------| > ", IntNum)
    fmt.Println("🐞 Close Ports       -----------| > ", CountPort1 - IntNum )
    if Conut.WriteFile !=""{
        OutFile += fmt.Sprintf("%s",strings.Repeat("_", 40))+"\n\n"
        OutFile += fmt.Sprintf("Domain IP         -----------| > %s ",IP)+"\n"
        OutFile += fmt.Sprintf("Guess OS          -----------| > %s", Value )+"\n"
        OutFile += fmt.Sprintf("EndTime           -----------| > %s",TimeEnd.Format("15:04:05"))+"\n"
        OutFile += fmt.Sprintf("Scan Time         -----------| > %s", AllTime )+"\n"
        OutFile += fmt.Sprintf("Port Conut        -----------| > %d", CountPort1 - CountPort2+1 )+"\n"
        OutFile += fmt.Sprintf("Ports Live        -----------| > %d", IntNum)+"\n"
        OutFile += fmt.Sprintf("Close Ports       -----------| > %d", CountPort1 - IntNum )+"\n"
    }
    writeFile(OutFile,Conut)
}  

func writeFile(OutPut string , DataInfo Config){
    var file *os.File
    var err error
    if _, err := os.Stat(DataInfo.WriteFile); err == nil {
        os.Remove(DataInfo.WriteFile)
   } 
    file, err = os.OpenFile(DataInfo.WriteFile, os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil{
        return
    }
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
