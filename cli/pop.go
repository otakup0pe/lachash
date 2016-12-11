package cli
import (
	"os"
	"fmt"
	"flag"
	"io/ioutil"
	"encoding/base64"
	"github.com/hashicorp/vault/api"	
	"github.com/otakup0pe/lachash/helpers"
	"github.com/otakup0pe/lachash/vault"
)

var token, hash_token, dest_file string

type Pop struct {}
func (c *Pop) Synopsis() string {
	return "Pop a stash from a Vault cubbyhole"
}
func (c *Pop) Help() string {
	f := pop_flags()
	help_msg := `%s

With the pop command you can receive information that
someone else has previously stashed in a Vault server.
You do not neccesarily need an account on the Vault server,
but you do need network access. If you do not specify
either a token or a hashed token, then the normal Vault
environmental defaults will be used. This is almost
certainly not actually what you want.
`
	fmt.Fprintln(os.Stderr, fmt.Sprintf(help_msg, c.Synopsis()))
	f.PrintDefaults()
	return fmt.Sprintf("")
}

func (c *Pop) Run(args []string) int {
	var junk []byte
	var client *api.Client
	f := pop_flags()
	if err := f.Parse(args); err != nil {
		helpers.Problems(err.Error())
	}
	if token != "" && hash_token != "" {
		helpers.Problems("Can not specify both token and hash-token")
	}
	if token == "" && hash_token == "" {
		helpers.Log("Using environmental token")
		client = vault.GetClient(vault.GetToken())
	} else if token != "" {
		helpers.Log("Using specified token")
		client = vault.GetClient(token)
	} else if hash_token != "" {
		helpers.Log("Using specified hashed token")
		client = vault.GetClient(helpers.DecodeUUID(hash_token))
	}
	data := vault.ReadStash(client)
	var d_err error
	junk, d_err = base64.StdEncoding.DecodeString(data)
	if d_err != nil {
		helpers.Problems(d_err.Error())
	}
	if dest_file != "" {
		ioutil.WriteFile(dest_file, junk, 0640)
	} else {
		helpers.Output(fmt.Sprintf("%s", junk))
	}
	return 0
}

func pop_flags() (fs *flag.FlagSet) {
	fs = helpers.Flags("pop")
	fs.StringVar(&token, "token", "", "Specify token")
	fs.StringVar(&hash_token, "hash-token", "", "Specify a hashed token")
	fs.StringVar(&dest_file, "output", "", "Write to file, instead of stdout")
	return
}
