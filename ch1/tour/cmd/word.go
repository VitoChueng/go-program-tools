package cmd

import (
	"github.com/VitoChueng/go-program-tools/ch1/tour/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var (
	str string
	mode int8
)


var wordCmd = &cobra.Command{
	Use:"word",
	Short: "word transfer",
	Long: desc,
	Run: func(cmd *cobra.Command,args []string){
		var content string
		switch mode {
		case Mode_Upper:
			content = word.ToUpper(str)
		case Mode_Lower:
			content = word.ToLower(str)
		case Mode_Underscore_To_Lower_CamelCase:
			content = word.UnderScoreToLoweCamelCase(str)
		case Mode_Underscore_To_Upper_CamelCase:
			content = word.UnderScoreToUpperCamelCase(str)
		case Mode_CamelCase_To_Underscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatal("sorry no such method")
		}
		log.Printf("result:%s",content)
	},
}

const (
	Mode_Upper = iota+1
	Mode_Lower
	Mode_Underscore_To_Upper_CamelCase
	Mode_Underscore_To_Lower_CamelCase
	Mode_CamelCase_To_Underscore
)

var desc = strings.Join([]string{
	"1","2","3","4","5",
},"\n")


func init(){
	wordCmd.Flags().StringVarP(&str,"str","s","","word content")
	wordCmd.Flags().Int8VarP(&mode,"mode","m",0,"word trans mode")
}

func Run(){

}