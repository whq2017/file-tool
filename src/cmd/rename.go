package cmd

import (
    "file-tools/src/rename"
    "fmt"
    "github.com/fatih/color"
    "github.com/spf13/cobra"
    "os"
    "strconv"
    "strings"
)

var(
    dest string
    src[] string
)

var renameCommand = &cobra.Command{
    Use: "rename [flags]",
    Short: "delete string in files name by flag <src>",
    //Long: `--`,
    Run: func(cmd *cobra.Command, args []string) {
        renameFunc()
    },
}

func init(){
    renameCommand.PersistentFlags().StringVarP(&path, "path", "p",".", "Setting files or directions path")
    renameCommand.PersistentFlags().StringVarP(&dest, "dest", "d", "","Replace string in the files name")
    renameCommand.PersistentFlags().StringArrayVarP(&src, "src", "s", nil,"Strings in files name")
}

func renameFunc(){
    
    color.Set(color.FgCyan)
    fmt.Println("Flags Info: ")
    fmt.Println(" Path: ", path)
    
    if strings.TrimSpace(dest) == "" {
        fmt.Printf(" Dest: %v [deletefile string]\n", strconv.Quote(dest))
    } else{
        fmt.Printf(" Dest: %v\n", dest)
    }
    
    fmt.Println(" String Array: ", src)
    color.Unset() // Don't forget to unset
    
    
    if src == nil ||
        strings.TrimSpace(path) == ""{
        
        helpRenameInfo()
    }
    
    err := rename.DeleteName(src, path, dest)
    if err != nil {
        fmt.Printf("Error: %s\n", err.Error())
        
        os.Exit(1)
    }
    
    magentaColor := color.New(color.FgMagenta).Add(color.Bold)
    _,_ = magentaColor.Println("\n Running complete.")
}

func helpRenameInfo(){
    red := color.New(color.FgRed)
    _, _ = red.Println("Usage: tools rename [--path/-p <path>] [--dest/-d <replace string>] --src/-s <string 1> <string 2> ...")
}