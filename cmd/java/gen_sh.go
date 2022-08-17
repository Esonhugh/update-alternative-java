package java

import (
	"fmt"
	"github.com/esonhugh/update-alternative-java/lib/define"
	"github.com/esonhugh/update-alternative-java/lib/misc"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var GenShCmd = &cobra.Command{
	Use:     "gen_sh",
	Aliases: []string{"gen"},
	Short:   "Generate shell script",
	Long:    `Generate shell script`,
	Run: func(cmd *cobra.Command, args []string) {
		var keys []define.Java
		row := JavaVersionDB.Find(&keys).RowsAffected
		if row == 0 {
			log.Error("Affect Rows is 0")
			return
		}
		shellFileContent := GenShCmds(keys)
		err := os.WriteFile(misc.Configlocate(), []byte(shellFileContent), 0644)
		if err != nil {
			log.Error(err)
		}
		log.Info("Generate shell script success")
		log.Info("ShellScript location at ", misc.Configlocate())
		log.Info("Please run following command to enable in your shell")
		log.Infof("echo 'source %v' >> ~/.zshrc", misc.Configlocate())
		log.Infof("echo 'source %v' >> ~/.bashrc", misc.Configlocate())
	},
}

const Header = `#!/bin/sh`
const Tails = `
default_java () {
	# use_java_8
}
default_java

##############################
# SHELL DESCRIPTION: 
# THIS SHELL IS MAINLY USED FOR SWITCHING YOUR CURRENT JAVA BINARY.
# USAGE: use_java_[version] to switch to it.
# EXAMPLE: use_java_8
# AUTO_GEN: By update-alternative-java
# also you can edit the default_java () function to switch to another java version when the script load manually.
##############################
`

func GenShCmds(keys []define.Java) string {
	Context := ""
	for _, key := range keys {
		shell_export := fmt.Sprintf(
			`export JAVA%v_HOME=%v`,
			key.Version, key.JavaHome)
		shell_using := fmt.Sprintf(
			`use_java_%v () {
  export JAVA_HOME=$JAVA%v_HOME
  export PATH="$JAVA%v_HOME/bin:$PATH"
}`,
			key.Version, key.Version, key.Version)
		Context = fmt.Sprintf("%v\n%v\n%v\n", Context, shell_export, shell_using)
	}
	return strings.Join([]string{Header, Context, Tails}, "\n")
}
