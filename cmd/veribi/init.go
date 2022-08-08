package veribi

import (
	"errors"
	"fmt"
	"net/mail"

	"github.com/manifoldco/promptui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Veribi CLI",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("cmd init")

		validateEmail := func(input string) error {
			_, err := mail.ParseAddress(input)
			if err != nil {
				return errors.New("invalid email")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:    "Veribi account email",
			Validate: validateEmail,
		}
		email, err := prompt.Run()
		if err != nil {
			log.Errorf("Prompt failed %v\n", err)
			return
		}
		viper.Set("email", email)

		prompt = promptui.Prompt{
			Label: "Password",
		}
		pass, err := prompt.Run()
		if err != nil {
			log.Printf("Prompt failed %v\n", err)
			return
		}
		viper.Set("password", pass)
		viper.Set("sessionId", "")

		if err := viper.WriteConfig(); err != nil {
			log.Info(err)
		}

		fmt.Println("You're all set abnd can login now.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
