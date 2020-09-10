package main
import (
	"fmt"
	"github.com/docker/distribution/registry"
	dcontext "github.com/docker/distribution/context"
	"github.com/docker/distribution/version"
	"github.com/docker/distribution/configuration"
	_ "github.com/docker/distribution/registry/storage/driver/azure"
	"bytes"
)

const strgAccName string = ""
const strgAccKey string = ""
const strgContainer string = ""

var dYAML string = ("version: 0.1\nstorage:\n  azure:\n    accountname: " + strgAccName + "\n    accountkey: " + strgAccKey + "\n    container: " + strgContainer + "\nhttp:\n  addr: :5000")

func main(){
	fmt.Println(dYAML)
	config, _ := configuration.Parse(bytes.NewReader([]byte(dYAML)))
	ctx := dcontext.WithVersion(dcontext.Background(), version.Version)
	r, _ := registry.NewRegistry(ctx, config)
	fmt.Println(dYAML,r)
}
