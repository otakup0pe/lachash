package helpers
import (
	"os"
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"encoding/base64"
	"archive/zip"	
	"flag"
	cli_mod "github.com/mitchellh/cli"	
)

var _ui cli_mod.Ui
var verbose bool
var stash_path string

func StashPath() string {
	return stash_path
}

func IsStdin() bool {
	stat, _ := os.Stdin.Stat()
	return ((stat.Mode() & os.ModeCharDevice) == 0)
}

func Log(msg string) {
	if verbose {
		_ui.Info(msg)
	}
}

func Output(msg string) {
	fmt.Println(msg)
}

func Problems(msg string) {
	_ui.Error(msg)
	os.Exit(1)
}

func EncodeUUID(p_uuid string) string {
	t_token, _ := uuid.Parse(p_uuid)
	b_token, _ := t_token.MarshalBinary()
	s_token := base64.StdEncoding.EncodeToString(b_token)
	r_token := fmt.Sprintf("%s", s_token[0:len(s_token) - 2])
	return r_token
}

func DecodeUUID(data string) string {
	e_token := fmt.Sprintf("%s==", data)
	b_token, _ := base64.StdEncoding.DecodeString(e_token)
	t_token := uuid.New()
	t_token.UnmarshalBinary(b_token)
	s_token := t_token.String()
	return s_token
}

func Flags(command string) (f *flag.FlagSet) {
	f = flag.NewFlagSet(command, 0)
	f.BoolVar(&verbose, "verbose", false, "For some detailed debugging information")
	f.StringVar(&stash_path, "path", "cubbyhole/lachash", "Specify a Vault path to use")
	return
}

func Compress(junk []byte) *bytes.Buffer {
	buf := new(bytes.Buffer)
	zipfile := zip.NewWriter(buf)
	handle, err := zipfile.Create("file.dat")
	if err != nil {
		Problems(err.Error())
	}
	_, err = handle.Write(junk)
	if err != nil {
		Problems(err.Error())
	}
	err = zipfile.Close()
	if err != nil {
		Problems(err.Error())
	}
	return buf
}

func Init() {
	_ui = &cli_mod.ColoredUi{
		OutputColor: cli_mod.UiColorNone,
		InfoColor: cli_mod.UiColorNone,
		ErrorColor: cli_mod.UiColorRed,
		WarnColor: cli_mod.UiColorYellow,
		Ui: &cli_mod.BasicUi{
			Writer: os.Stderr,
			ErrorWriter: os.Stderr,
		},
	}
}
