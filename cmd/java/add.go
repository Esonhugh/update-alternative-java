package java

import (
	"bufio"
	"github.com/AlecAivazis/survey/v2"
	"github.com/esonhugh/update-alternative-java/lib/define"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os/exec"
	"path"
)

var AddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add alternative java version",
	Long:    `Add alternative java version`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			AskAndAdd()
		}
	},
}

func AskAndAdd() *define.Java {
	var qs = []*survey.Question{
		{
			Name: "version",
			Prompt: &survey.Select{
				Message: "Java 版本 (Please select version of java)",
				Options: (define.GetJavaVersionList()),
			},
		},
		{
			Name: "javahome",
			Prompt: &survey.Input{
				Message: "请输入 java 的安装路径 (Please input java home path)",
			},
		},
	}

	// Generate the new config struct named cred to receive the inputted values.
	Answer := struct {
		Version  string `survey:"version"`
		JavaHome string `survey:"javahome"`
	}{}

	survey.Ask(qs, &Answer)

	promot := &survey.Confirm{
		Message: "以上信息是否正确 (make sure correctness) "}
	sure := true // Break out
	survey.AskOne(promot, &sure)
	if !sure {
		logrus.Info("User cancel add new java version")
		return nil
	}

	key := define.Java{
		Version:  define.JAVA_VERSION_MAP[Answer.Version],
		JavaHome: Answer.JavaHome,
	}
	JavaBin := path.Join(key.JavaHome, "bin", "java")
	cmd := exec.Command(JavaBin, "-version")
	errorpipe, err := cmd.StderrPipe()
	if err != nil {
		logrus.Error(err)
	}
	if err := cmd.Start(); err != nil {
		logrus.Error("Java not found in ", key.JavaHome)
		logrus.Error("Or Added Binary is invalid.")
		logrus.Panicln(err)
	}
	buffer, _, err := bufio.NewReader(errorpipe).ReadLine()
	key.DetailVersion = string(buffer)
	// Make user to check
	JavaVersionDB.Where("version = ?", key.Version).Save(&key)
	return &key
}
