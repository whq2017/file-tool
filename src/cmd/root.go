package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
)

const(
    VERSION = 0.1
)

var(
    path string
)

var rootCommand = &cobra.Command{
    Use: "tools",
    Short: "tools is a file deal tool array.",
    //Long: `--`,
    Run: func(cmd *cobra.Command, args []string) {
        helpInfo()
    },
}

var versionCommand = &cobra.Command{
    Use: "version",
    Short: "tools is a file deal tool array.",
    Long: `--`,
    Run: func(cmd *cobra.Command, args []string) {
        versionInfo()
    },
}


func init() {
    
    addRootCommandFlags()
    
    rootCommand.AddCommand(deleteCommand)
    rootCommand.AddCommand(renameCommand)
    rootCommand.AddCommand(versionCommand)
}

func helpInfo(){
    fmt.Println("tools -h")
}

func versionInfo(){
    fmt.Println(`Tools version: `,VERSION)
}

func addRootCommandFlags(){
    rootCommand.Flags().BoolP("version", "v", false, "Help message for version")
}



func Execute() {
    if err := rootCommand.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}