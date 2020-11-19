package rename

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    
    "github.com/fatih/color"
    
)

func DeleteName(deleteStr []string, path, destStr string) error{
    infos, e := getFiles(path)
    if e != nil {
        return e
    }
    for _, v := range infos {
        if v.IsDir() {
            _ = DeleteName(deleteStr, filepath.Join(path,v.Name()), destStr)
        }
        
        // 没有放在else里面，是为了处理目录名字中的字符串
        dealFileName(deleteStr, v, path, destStr)
    }
    
    return nil
}

func dealFileName(deleteStr []string, file os.FileInfo, path, destStr string){
    for _, v := range deleteStr {
        if strings.Contains(file.Name(), v) {
            oldName := filepath.Join(path, file.Name())
            newName := filepath.Join(path, strings.ReplaceAll(file.Name(), v, destStr))
            
            _,_ = fmt.Fprintf(color.Output,"oldName: %v \n --> newName: %v\n", color.BlueString(oldName), color.RedString(newName))
            
            _ = os.Rename(oldName,newName)
             
             break
        }
    }
}

func getFiles(dirPath string) ([]os.FileInfo, error){
    
    fileInfoList, e := ioutil.ReadDir(dirPath)
    if e != nil {
        return nil, e
    }
    
    return fileInfoList,nil
}
