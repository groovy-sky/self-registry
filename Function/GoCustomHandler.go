package main

import (
    //"log"
    "os"
    "github.com/docker/distribution/registry"
    dcontext "github.com/docker/distribution/context"                                                         
    "github.com/docker/distribution/version"                                                                  
    "github.com/docker/distribution/configuration"                                                            
    _ "github.com/docker/distribution/registry/storage/driver/azure"                                          
    "bytes"
)

//func httpTriggerHandler(w http.ResponseWriter, r *http.Request) {
//    w.Write([]byte("Hello There"))
//}

func main() {
	httpInvokerPort,_ := os.LookupEnv("FUNCTIONS_HTTPWORKER_PORT")
    strgAccName,_ := os.LookupEnv("FUNCTIONS_STORAGE_NAME")                                                   
    strgAccKey,_ := os.LookupEnv("FUNCTIONS_STORAGE_KEY")                                                     
    strgContainer,_ := os.LookupEnv("FUNCTIONS_STORAGE_CONTAINER")                                            
    var dYAML string = ("version: 0.1\nstorage:\n  azure:\n    accountname: " + strgAccName + "\n    accountkey: " + strgAccKey + "\n    container: " + strgContainer + "\nhttp:\n  addr: :" + httpInvokerPort + "\n  prefix: /api/httptrigger/\n  relativeurls: true\n  host: http://localhost:7071/api/httptrigger") 
    config, _ := configuration.Parse(bytes.NewReader([]byte(dYAML)))                                          
    ctx := dcontext.WithVersion(dcontext.Background(), version.Version)                                       
    r, _ := registry.NewRegistry(ctx, config)                                                                 
    r.ListenAndServe() 
    //log.Println("Go server Listening...on httpInvokerPort:", httpInvokerPort)
	//log.Fatal(http.ListenAndServe(":"+httpInvokerPort, mux))
}
