package cmd

import (
	"github.com/VitoChueng/go-program-tools/ch1/tour/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var (
	dur string
	calTime string
)

var timeCmd = &cobra.Command{
	Use: "",
	Short: "",
	Long: "",
	Run: func(cmd *cobra.Command,args []string){},
}

var nowTimeCmd = &cobra.Command{
	Use: "now",
	Short: "",
	Long: "",
	Run: func(cmd *cobra.Command,args []string){
		nowTime := timer.GetNowTime()
		log.Printf("%s %d",nowTime.Format(time.RFC822Z),nowTime.Unix())
	},
}