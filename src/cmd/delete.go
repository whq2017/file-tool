package cmd

import (
    "file-tools/src/deletefile"
    "fmt"
    "github.com/fatih/color"
    "github.com/spf13/cobra"
    "os"
    "strings"
)

var deleteCommand = &cobra.Command{
    Use: "delete [flags]",
    Short: "tools is a file deal tool array.",
    //Long: `--`,
    Run: func(cmd *cobra.Command, args []string) {
        deleteInfo()
    },
}

var(
    fileNames []string
    toAsk bool
)

func init(){
    deleteCommand.PersistentFlags().StringArrayVarP(&fileNames,"names", "n", nil, "File name Array to delete")
    deleteCommand.PersistentFlags().StringVarP(&path,"path", "p", ".","Setting files or directions path")
    deleteCommand.PersistentFlags().BoolVarP(&toAsk,"ask", "a", false,"Ask user to confirm when file name is similar.")
}

func deleteInfo(){
    
    color.Set(color.FgCyan)
    fmt.Println("Flags Info: ")
    fmt.Println(" Path: ", path)
    fmt.Println(" File Names: ", fileNames)
    color.Unset() // Don't forget to unset
    
    if strings.TrimSpace(path) == "" ||
        fileNames == nil {
        
        helpDeleteInfo()
    }
    
    deletefile.SetAsk(toAsk)
    
    deleteCnt, err := deletefile.DeleteFiles(path, fileNames)
    if err != nil {
        _,_ = color.New(color.FgRed).Printf("Error: %s\n", err.Error())
    
        os.Exit(1)
    }
    
    magentaColor := color.New(color.FgMagenta).Add(color.Bold)
    _,_ = magentaColor.Println("\n Run complete!")
    _,_ = magentaColor.Println(" Delete file count: ", deleteCnt)
}

func helpDeleteInfo(){
    red := color.New(color.FgRed)
    _, _ = red.Println("Usage: tools delete [--path/-p <path>] [--names/-n <string 1>] [--names/-n <string 2>] [...]")
}