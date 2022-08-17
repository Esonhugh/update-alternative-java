package java

import (
	"github.com/esonhugh/update-alternative-java/cmd"
	"github.com/esonhugh/update-alternative-java/lib/db"
	"github.com/esonhugh/update-alternative-java/lib/define"
	"gorm.io/gorm"
)

var JavaVersionDB *gorm.DB

func init() {
	cmd.RootCmd.AddCommand(AddCmd)
	cmd.RootCmd.AddCommand(GenShCmd)
	db.DB.Main.AutoMigrate(&define.Java{})
	JavaVersionDB = db.DB.Main
}
