package vault

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	"strconv"
	"github.com/hashicorp/vault/api"
	"github.com/otakup0pe/lachash/helpers"
)

func GetToken() (token string) {
	env_token := os.Getenv("VAULT_TOKEN")
	if len(env_token) > 0 {
		token = env_token
	} else {
		file_token, err := ioutil.ReadFile(fmt.Sprintf("%s/.vault-token", os.Getenv("HOME")))
		if err == nil {
			token = strings.TrimSpace(string(file_token))
		} else {
			helpers.Problems("Unable to find token")
		}
	}
	return
}

func GetClient(token string) (*api.Client) {
	client, c_err := api.NewClient(api.DefaultConfig())
	if c_err != nil {
		helpers.Problems(c_err.Error())
	}
	helpers.Log(fmt.Sprintf("Server %s", client.Address()))
	client.SetToken(token)
	return client
}

func GetStashToken(client *api.Client, ttl int, uses int, policy string,) string {
	policies := []string{policy}
	helpers.Log(fmt.Sprintf("Requesting token for %ds, %d uses", ttl, uses))
	var use_default bool = false
	if policy != "default" {
		use_default = true
	}
	tcr := &api.TokenCreateRequest{
		TTL: strconv.Itoa(ttl),
		Policies: policies,
		DisplayName: fmt.Sprintf("lachash"),
		NoDefaultPolicy: use_default,
		NumUses: uses + 1,
	}
	token_secret, t_err := client.Auth().Token().Create(tcr)
	if t_err != nil {
		helpers.Problems(t_err.Error())
	}
	return token_secret.Auth.ClientToken
}

func ReadStash(client *api.Client) (data string) {
	stash_secret, err := client.Logical().Read(helpers.StashPath())
	if err != nil {
		helpers.Problems(err.Error())
	}
	if stash_secret == nil {
		helpers.Problems("Nothing found!")
	}
	if deets, err := stash_secret.Data["data"]; err {
		data = fmt.Sprintf("%v", deets)
	}
	return
}

func WriteStash(junk string, client *api.Client) {
	var data = make(map[string]interface{})
	data["data"] = junk
	_, err := client.Logical().Write(helpers.StashPath(), data)
	if err != nil {
		helpers.Problems(err.Error())
	}
}
