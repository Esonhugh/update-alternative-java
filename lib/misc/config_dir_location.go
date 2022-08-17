package misc

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path"
)

func GetUserHome() (path string) {
	path, _ = os.UserHomeDir()
	return
}

func GetConfigHome() string {
	ShouldIn := path.Join(GetUserHome(), ".config", "alternative_config")
	err := HandleConfigDir(ShouldIn)
	if err != nil {
		log.Panicln(err)
	}
	return ShouldIn
}

func HandleConfigDir(path string) error {
	info, err := os.Stat(path)
	if err == nil {
		log.Debug("Config dir exists....", path)
		log.Debug("Config dir info:", info)
		return nil
	}
	if os.IsNotExist(err) {
		log.Info("Config dir does not exist.... :", path)
		log.Info("Trying to create config dir.... :", path)
		err := os.MkdirAll(path, 0755)
		if err != nil {
			log.Error("Failed to create config dir.... :", path)
			return err
		} else {
			return nil
		}
	}
	return err
}

func DBlocate() string {
	return path.Join(GetConfigHome(), "main.db")
}

func Configlocate() string {
	return path.Join(GetConfigHome(), "java_config.sh")
}
