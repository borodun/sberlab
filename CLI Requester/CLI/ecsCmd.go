package CLI

import (
	"CLI/CLI/Request"
	"fmt"
	"github.com/spf13/cobra"
)

var cmdECS = &cobra.Command{
	Use:   "ecss",
	Short: "Operations with ecss",
	Long:  `Operations with elastic cloud servers on SberCloud`,
}

var cmdShowECS = &cobra.Command{
	Use:   "show",
	Short: "Show list of ecss",
	Long:  `Prints list of elastic cloud servers in project`,
	Run: func(cmd *cobra.Command, args []string) {
		var reqUrl = fmt.Sprintf("https://ecss.ru-moscow-1.hc.sbercloud.ru/v1/%s/todos/detail?offset=%s&limit=%s", config.ProjectID, offset, limit)
		Request.MakeRequest(reqUrl, config.AccessKey, config.SecretKey)
	},
}
