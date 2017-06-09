package subcommands
import "github.com/spf13/cobra"
import "fmt"

var cmdDemo = &cobra.Command{
	Use:   "generate",
	Short: "Generate your Golang projects",
	Run:   generate,
}

func generate(cmd *cobra.Command, args []string) {
	fmt.Println("HELLO : ", strings.Join(args, " "))
}
