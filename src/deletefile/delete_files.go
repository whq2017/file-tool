package deletefile

import (
    "errors"
    "fmt"
    "github.com/fatih/color"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
)

var(
    nameMap map[string]bool
    toAsk bool
    
    deleteCount = 0
)

const(
    INC  =  5
)


func DeleteFiles(path string, names []string) (int,error) {
    
    if err :=dealNames(names); err != nil{
        return deleteCount,err
    }
    
    return deleteCount,deleteFiles(path)
}

func SetAsk(ask bool){
    toAsk = ask
}

func deleteFiles(path string) error {
    fileInfos, err := ioutil.ReadDir(path)
    if err != nil {
        return err
    }
    for _, v := range fileInfos {
        
        newPath := filepath.Join(path, v.Name())
        if v.IsDir() {
            _ = deleteFiles(newPath)
        } else if isExistName(v.Name()) {
            _ = os.Remove(newPath)
            _,_ = fmt.Fprintf(color.Output,"Delete file: %v success!\n", color.RedString(newPath))
            deleteCount++
        }
    }
    return nil
}

func isExistName(name string) bool {
    var newName string
    
    if nameMap[name] {
        return true
    }else{
        nameSplit := strings.Split(name, ".")
        if len(nameSplit) == 1 {
            newName = nameSplit[0]
        } else{
            newName = strings.Join(nameSplit[:len(nameSplit)-1] , ".")
        }
        
        if nameMap[newName] {
            _,_ = fmt.Fprintf(color.Output,"Do you want to delete: %v ?\n", color.BlueString(name))
            if toAsk {
                return userControlOp(1, name)
            }
        }
    }
    return false
}

func dealNames(names[]string) error {
    n := len(names)
    if n == 0 {
        return errors.New("需要指定删除的文件名")
    }
    nameMap = make(map[string]bool, n + INC)
    
    for _, v := range names {
        nameMap[v] = true
    }
    
    return nil
}

func userControlOp(count int, name string) bool {
    
    if count > 3 {
        return false
    }
    
    var userInput string
    fmt.Printf("Enter y/yes, n/no or a/all.(default y): ")
    _,err := fmt.Scanln(&userInput)
    
    if err != nil{
        if err.Error() == "unexpected newline" {
            userInput = "y"
        }else {
            color.Red("Enter error data: ", err.Error())
            return userControlOp(count + 1,name)
        }
    }
    
    if userInput == "y" || userInput == "yes" {
         return true
    } else if userInput == "n" || userInput == "no" {
        return false
    }else if userInput == "a" || userInput == "all" {
        nameMap[name] = true
        return true
    }else {
        return userControlOp(count + 1,name)
    }
}
