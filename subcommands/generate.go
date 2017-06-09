package subcommands

import "github.com/spf13/cobra"
import "fmt"
import "strings"



type Subs struct {

}
func (s *Subs)generate(cmd *cobra.Command, args []string) {
	fmt.Println("HELLO : ", strings.Join(args, " "))
}

func (s *Subs)AddGenerate(root *cobra.Command)  {

  var cmdDemo = &cobra.Command{
  	Use:   "generate",
  	Short: "Generate your Golang projects",
  	Run: s.generate,
  }

  root.AddCommand(cmdDemo)
}
