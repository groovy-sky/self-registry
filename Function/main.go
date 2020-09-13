package main
import (
	//"fmt"
	"github.com/docker/distribution/registry"
	dcontext "github.com/docker/distribution/context"
	"github.com/docker/distribution/version"
	"github.com/docker/distribution/configuration"
	_ "github.com/docker/distribution/registry/storage/driver/azure"
	"bytes"
    "os"
)



func main(){
    strgAccName,_ := os.LookupEnv("FUNCTIONS_STORAGE_NAME")
    strgAccKey,_ := os.LookupEnv("FUNCTIONS_STORAGE_KEY")
    strgContainer,_ := os.LookupEnv("FUNCTIONS_STORAGE_CONTAINER")
    var dYAML string = ("version: 0.1\nstorage:\n  azure:\n    accountname: " + strgAccName + "\n    accountkey: " + strgAccKey + "\n    container: " + strgContainer + "\nhttp:\n  addr: :5000")
	config, _ := configuration.Parse(bytes.NewReader([]byte(dYAML)))
	ctx := dcontext.WithVersion(dcontext.Background(), version.Version)
	r, _ := registry.NewRegistry(ctx, config)
    r.ListenAndServe()
}
