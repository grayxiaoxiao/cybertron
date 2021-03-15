package transform

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/grayxiaoxiao/cybertron/cmd/cube/tmpl"

	"github.com/spf13/cobra"
)

var CmdNew = &cobra.Command{
    Use:   "new",
    Short: "create a struct/构建一个struct",
    Long:  "create a struct/构建一个struct, Example: cybertron new customer org_id:int64",
    Run:   run,
}

func run(cmd *cobra.Command, args []string) {
    if len(args) == 0 {
        logsPrint("Error", "struct name is required.\033[m Example: cybertron new customer")
        return
    }
    currentDir, err := os.Getwd()
    if err != nil {
        logsPrint("Error", "get current directory failed with error " + err.Error())
        panic(err)
    }
    logsPrint("Info", "work in directory " + currentDir)

    interfaceDir, fullDir, packageName, cybertronName := autobots(args[0], currentDir)
    logsPrint("Info", "create struct " + cybertronName + " of package " + packageName + " in directory " + fullDir)
    fields := getStructFields(args)
    writeStructFile(fullDir, packageName, cybertronName, fields)
    writeInterfaceFile(interfaceDir)
    logsPrint("Info", "create struct" + cybertronName + " success........")
}

func writeStructFile(fullDir, packageName, cybertronName string, fields []tmpl.StructField) error {
    _, err := os.Stat(fullDir)
    if os.IsNotExist(err) {
        os.MkdirAll(fullDir, os.ModePerm)
    }
    tableName := strings.ToLower(cybertronName)
    fileName  := tableName + ".go"
    file, err := os.Create(fullDir + "/" + fileName)
    if err != nil {
        logsPrint("Error", "create struct file failed with error " + err.Error())
    }
    defer file.Close()
    tmplStr, err := template.New("struct").Parse(tmpl.StructTMPL)
    if err != nil {
        logsPrint("Error", "parse struct template file failed with error " + err.Error())
    }

    structMap  := map[string]interface{}{
        "packageName":   packageName,
        "cybertronName": cybertronName,
        "structFields":  fields,
        "tableName":     tableName,
    }
    structFile := new(bytes.Buffer)
    err = tmplStr.Execute(structFile, structMap)
    if err != nil {
        logsPrint("Error", "execute struct template file failed with error " + err.Error())
    }
    _, err = file.WriteString(strings.Replace(structFile.String(), "\n\n", "\n", -1))
    if err != nil {
        logsPrint("Error", "write struct file failed with error " + err.Error())
    }
    return err
}

func getStructFields(args []string) []tmpl.StructField {
    num := len(args)
    if num == 1 {
        return []tmpl.StructField{}
    } else {
        pgs := args[1:]
        fields := make([]tmpl.StructField, len(pgs))
        logsPrint("Info", "creating struct fields................")
        for _, item := range pgs {
            sps := strings.Split(item, ":")
            if len(sps) != 2 {
                logsPrint("Warning", item + " is not an effective parameter")
                continue
            }
            field := tmpl.StructField{
                FieldType: sps[1],
                FieldTags: "`column_name:\"" + sps[0] + "\" json:\"" + sps[0] + "\"`",
            }
            _sps := strings.Split(sps[0], "_")
            fieldName := ""
            for _, es := range _sps {
                fieldName += strings.Title(es)
            }
            field.FieldName = fieldName
            fields = append(fields, field)
        }
        return fields
    }
}

func writeInterfaceFile(interfaceDir string) error {
    _, err := os.Stat(interfaceDir + "/interfaces")
    if os.IsNotExist(err) {
        os.MkdirAll(interfaceDir + "/interfaces", os.ModePerm)
    }
    file, err := os.Create(interfaceDir + "/interfaces/cybertron.go")
    if err != nil {
        logsPrint("Error", "create interface file failed with error " + err.Error())
    }
    defer file.Close()
    _, err = file.WriteString(tmpl.InterfaceTMPL)
    if err != nil {
        logsPrint("Error", "write interface file failed with error " + err.Error())
    }
    return err
}

func autobots(pms string, currentDir string) (interfaceDir, fullDir, pkgName, cbtName string) {
    pgs := strings.Split(pms, "/")
    num := len(pgs)
    if num == 1 {
        return currentDir + "/models", currentDir + "/models", "models", strings.Title(pgs[0])
    } else {
        return currentDir + "/" + pgs[0], currentDir + "/" + strings.Join(pgs[0:num-1], "/"), pgs[num - 2], strings.Title(pgs[num - 1])
    }
}

func logsPrint(mode, log string) {
    switch mode {
    case "Error":
        fmt.Fprintf(os.Stdout, "\033[31mError: %s\033[m\n", log)
    case "Info":
        fmt.Fprintf(os.Stdout, "\033[32mInfo: %s\033[m\n", log)
    case "Warning":
        fmt.Fprintf(os.Stdout, "\033[33mWarning: %s\033[m\n", log)
    default:
        fmt.Fprintf(os.Stdout, log)
    }
}