package cli
import (
	"os"
	"fmt"
	"flag"
	"io/ioutil"
	"encoding/base64"
	"github.com/otakup0pe/lachash/helpers"
	"github.com/otakup0pe/lachash/vault"
)

type Stash struct {}
var ttl, uses int
var return_hash_token bool
var source_file, policy string

func (c *Stash) Synopsis() string {
	return "Stash something in a Vault cubbyhole"
}

func (c *Stash) Help() string {
	f := stash_flags()
	help_msg := `%s

With the stash command you can safely exchange information
with anyone else who has network access to the same Vault
server as you. Note that since lachash generates unique
vault tokens for every stash, the receiving party does not
neccesarily need an account with the Vault server in
question. You do still need an account to stash. The normal
mechanisms apply for looking up a Vault server and Vault
Token prior to making connections.
`
	fmt.Fprintln(os.Stderr, fmt.Sprintf(help_msg, c.Synopsis()))
	f.PrintDefaults()
	return fmt.Sprintf("")
}

func (c *Stash) Run(args []string) int {
	var junk []byte
	var err error
	if err != nil {
		helpers.Problems(err.Error())
	}
	f := stash_flags()
	if err := f.Parse(args); err != nil {
		helpers.Problems(err.Error())
	}
	if source_file != "" {
		junk, err = ioutil.ReadFile(source_file)
	} else if helpers.IsStdin() {
		junk, err = ioutil.ReadAll(os.Stdin)
	} else {
		helpers.Problems("Nothing to stash")
	}
	
	var encoded = base64.StdEncoding.EncodeToString(junk)
	client := vault.GetClient(vault.GetToken())
	stash_token := vault.GetStashToken(client, ttl, uses, policy)
	client.SetToken(stash_token)
	vault.WriteStash(encoded, client)
	if return_hash_token {
		helpers.Log(fmt.Sprintf("Stash Token %s", stash_token))
		helpers.Output(helpers.EncodeUUID(stash_token))
	} else {
		helpers.Output(stash_token)
	}
	return 0
}

func stash_flags() (fs *flag.FlagSet) {
	fs = helpers.Flags("stash")
	fs.IntVar(&ttl, "ttl", 1800, "Time To Live for the stashed data, in seconds")
	fs.IntVar(&uses, "uses", 1, "How many times the stashed data may be retrieved")
	fs.BoolVar(&return_hash_token, "hash-token", false, "Produce a hashed token")
	fs.StringVar(&source_file, "input", "", "File to stash")
	fs.StringVar(&policy, "policy", "default", "The Vault policy to use")
	return
}
