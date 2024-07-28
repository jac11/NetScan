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
var OSGuess = ""
var current_time = time.Now().Local()
var distroLinux = ""

var Banner string = `
███████  ██████  █████  ███    ██ ███    ██ ███████ ████████ 
██      ██      ██   ██ ████   ██ ████   ██ ██         ██    
███████ ██      ███████ ██ ██  ██ ██ ██  ██ █████      ██    
     ██ ██      ██   ██ ██  ██ ██ ██  ██ ██ ██         ██    
███████  ██████ ██   ██ ██   ████ ██   ████ ███████    ██.Go    
                        @jacstory`+"\n"
type Config struct{

    Port       string
    Domain     string
    StartScan  string 
    EndScan    string
    WriteFile  string

}
func Style (Styles Config) {
    
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
func CheckNet(Domain string){
    if  strings.HasPrefix(Domain, "http://") || strings.HasPrefix(Domain, "https://"){
        fmt.Println(Blue+Banner+Reset)
        fmt.Println("⛔ Error-Status-2\t\t    -----------| > "+Cyan+"Use Domain Without "+Red+"{ "+Cyan+"http://-or-https://"+Red+" }"+Reset)
        os.Exit(0)
    }
    cmd := exec.Command("ping", "-c", "1",Domain)
    _, err := cmd.Output()
    if  err != nil {
        Command := exec.Command("dig" ,Domain ,"+short")
        _ ,err := Command.Output()
        if err != nil{
            fmt.Println(Blue+Banner+Reset) 
            fmt.Println("⛔️Error-Status-9\t\t    -----------| > "+Cyan +"Executing Internet Error"+Reset)
            os.Exit(0)
        } 
    } 
}                            
func PingHost(Domain string)(string,string){
    var matches,matches2 string 
    cmd := exec.Command("ping", "-c", "1",Domain)
    output, err := cmd.Output()
    if  err != nil {
        Command := exec.Command("dig" ,Domain ,"+short")
        OutPut ,err := Command.Output()
        if err != nil{
            fmt.Println("⛔️ Status-1\t\t    -----------| > Executing Internet Error")
            fmt.Println("⛔️ Status-2\t\t    -----------| > Use Domain Without http/https ")
            os.Exit(0)
        }else{
            matches  ,  matches2 :=  "OS UnKwon ", strings.Replace(string(OutPut),"\n","",1)
            return matches,matches2
        }     
    }else{
        re := regexp.MustCompile(`ttl=(\d+)`)
        matches := re.FindStringSubmatch(string(output))[1]
        IPGet := regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`)
        matches2 := IPGet.FindStringSubmatch(string(output))[1]
        return matches,matches2
    }        
    return matches,matches2
}
func ScanSinglPort(Domain string, Port string){
    IntNum = 0
    DomainNet := net.JoinHostPort(Domain,Port)
    Connect, err := net.DialTimeout("tcp",DomainNet,3*time.Second)
    if err != nil{
       fmt.Printf("🕵‍  Connection Fail          -----------| > %s %s Close %s\n ", Port, Red , Reset)
       OutFile += fmt.Sprintf("Connection Fail     -----------| > %s",Port)+" Close\n"
       return
    }
    buffer := make([]byte, 512)
    Connect.SetReadDeadline(time.Now().Add(1* time.Second)) 
    ServicePort, _ := Connect.Read(buffer)
    if ServicePort== 0 {
        for PortService , Service  := range (myMap){
            if Port == PortService {
                fmt.Printf("🚀️ Connection Succeeded     -----------| > %s %s Open %s %s %s\n", Port, Red, Reset, Cyan, Service)
                OutFile += fmt.Sprintf("Connection Succeeded     -----------| > %s",Port)+" Open "+Service+"\n"
                Connect.Close()
                IntNum = 1
            }
        }
    }else{
        ServiceName := string(buffer)
        fmt.Printf("🚀️ Connection Succeeded     -----------| > %s %s Open %s %s %s", Port, Red, Reset, Cyan, ServiceName)
        OutFile += fmt.Sprintf("Connection Succeeded     -----------| > %s",Port)+" Open "+ ServiceName +"\n"
        Connect.Close()
        IntNum = 1
        if strings.Contains(ServiceName,"Windows") || strings.Contains(ServiceName,"Microsoft"){
            OSGuess  = "128"
        }
    }
}
func ScanRangePort(Domain string,Start string , End string){
    IntNum ++
    var WaitGroup sync.WaitGroup
    var mutex sync.Mutex
    StartInt,_ := strconv.Atoi(Start)
    EndInt,_ := strconv.Atoi(End)
    for Port := StartInt ; Port <= EndInt ;Port++{
        WaitGroup.Add(1)
        go func (Port int){
            defer  WaitGroup.Done()
            DomainNet := net.JoinHostPort(Domain,strconv.Itoa(Port))
            Connect , err := net.DialTimeout("tcp",DomainNet,3*time.Second)
            mutex.Lock()
            defer mutex.Unlock()
            if err != nil {
                fmt.Printf("🕵‍  Connection Fail          -----------| > %-5d%s", Port, Red+" Close"+Reset)
                time.Sleep(10 * time.Millisecond)
                fmt.Print("\033[G\033[K") 
            }else{
                IntNum ++
                buffer := make([]byte, 512)
                timeoutDuration := 100 * time.Millisecond
                Connect.SetReadDeadline(time.Now().Add(timeoutDuration)) 
                ServicePort, _ := Connect.Read(buffer)
                if ServicePort== 0 {
                    for PortService , Service  := range (myMap){
                        if strconv.Itoa(Port) == PortService {
                            fmt.Printf("🚀️ Connection Succeeded     -----------| > %-4d %s Open %s %s %s%s\n",Port, Red, Reset, Cyan, Service,Reset)
                            OutFile += fmt.Sprintf("Connection Succeeded     -----------| > %s",Port)+" Open "+Service+"\n"
                            Connect.Close()
                            fmt.Print("\033[G\033[K")
                        }
                    }         
                }else{
                    ServiceName := string(buffer)
                    fmt.Printf("🚀️ Connection Succeeded     -----------| > %-4d %s Open %s %s %s%s",Port, Red, Reset, Cyan, ServiceName,Reset)
                    OutFile += fmt.Sprintf("Connection Succeeded     -----------| > %s",Port)+" Open "+ ServiceName +"\n"
                    Connect.Close()
                    if strings.Contains(ServiceName,"Windows") || strings.Contains(ServiceName,"Microsoft"){
                        OSGuess  = "128"

                        fmt.Print("\033[G\033[K")
                    }else{
                        for _, distro := range linuxDistributions{
                            if strings.Contains(ServiceName,distro){
                                OSGuess  = "65"
                                distroLinux = distro
                                fmt.Print("\033[G\033[K")
                                break
                            }
                        }
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
        if TTL == strconv.Itoa(TTLValue) && OSGuess == ""{
            Value = TTLMessage
            fmt.Println(OSGuess)
            break
        }else if OSGuess == "128" {
            Value = "128"
            Value = "Windows NT/2000/XP/Vista/7/10"
        }else if OSGuess == "65"{
            Value = "Linux "+distroLinux
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
    if IntNum ==1{
        fmt.Println("🪲️ Ports Live        -----------| > ", IntNum)
    }else if IntNum >= 2{
        fmt.Println("🪲️ Ports Live        -----------| > ", IntNum -1 )
    }
    if IntNum ==1 {
        fmt.Println("🐞 Close Ports       -----------| > ", IntNum-1 )
    }else if IntNum == 0 {
        fmt.Println("🐞 Close Ports       -----------| > ", IntNum+1)
    }else{
        fmt.Println("🐞 Close Ports       -----------| > ", CountPort1 - IntNum+1 )
    }
    if Conut.WriteFile !=""{
        OutFile += fmt.Sprintf("%s",strings.Repeat("_", 40))+"\n\n"
        OutFile += fmt.Sprintf("Domain IP         -----------| > %v ",IP)+"\n"
        OutFile += fmt.Sprintf("Guess OS          -----------| > %s", Value )+"\n"
        OutFile += fmt.Sprintf("EndTime           -----------| > %s",TimeEnd.Format("15:04:05"))+"\n"
        OutFile += fmt.Sprintf("Scan Time         -----------| > %s", AllTime )+"\n"
        OutFile += fmt.Sprintf("Port Conut        -----------| > %d", CountPort1 - CountPort2+1 )+"\n"
        OutFile += fmt.Sprintf("Ports Live        -----------| > %d", IntNum)+"\n"
        if IntNum ==1{
            OutFile += fmt.Sprintf("Close Ports       -----------| > %d", IntNum-1 )+"\n"
        }else if IntNum == 0{
            OutFile += fmt.Sprintf("Close Ports       -----------| > %d", IntNum+1 )+"\n"
        }else{
           OutFile += fmt.Sprintf("Close Ports       -----------| > %d", CountPort1 - IntNum+1 )+"\n" 
        }
        
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
    CheckNet(DataInfo.Domain)
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
