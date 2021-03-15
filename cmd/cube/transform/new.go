package transform

import (
	"bytes"
	"cybertron/cmd/cube/tmpl"
	"fmt"
	"os"
	"strings"
	"text/template"

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
        fmt.Fprintf(os.Stderr, "\033[31mError: Struct name is required.\033[m Example: cybertron new customer\n")
        return
    }
    currentDir, err := os.Getwd()
    if err != nil {
        fmt.Fprintf(os.Stderr, "\033[31mGet current directory failed with error '%s'.\033[m\n", err.Error())
        panic(err)
    }
    fmt.Fprintf(os.Stdout, "\033[32mWorking in directory is '%s'\033[m\n", currentDir)

    interfaceDir, fullDir, packageName, cybertronName := autobots(args[0], currentDir)
    fmt.Fprintf(os.Stdout, "\033[32mCreating struct '%s' of package '%s' in directory '%s'..........\033[m\n", cybertronName, packageName, fullDir)
    fields := getStructFields(args)
    writeStructFile(fullDir, packageName, cybertronName, fields)
    writeInterfaceFile(interfaceDir)
    fmt.Fprintf(os.Stdout, "\033[32mCreate struct success...........\033[m\n")
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
        fmt.Fprintf(os.Stderr, "\033[31mCreating struct file failed with error '%s'\033[m\n", err.Error())
    }
    defer file.Close()
    tmplStr, err := template.New("struct").Parse(tmpl.StructTMPL)
    if err != nil {
        fmt.Fprintf(os.Stderr, "\033[31mParse struct template file failed with error '%s'\033[m\n", err.Error())
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
        fmt.Fprintf(os.Stderr, "\033[31mExecute struct template file failed with error '%s'\033[m\n", err.Error())
    }
    _, err = file.WriteString(strings.Replace(structFile.String(), "\n\n", "\n", -1))
    if err != nil {
        fmt.Fprintf(os.Stderr, "\033[31mWrite struct file failed with error '%s'\033[m\n", err.Error())
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
        fmt.Fprintf(os.Stdout, "\033[32mCreating struct fields ...........\033[m\n")
        for _, item := range pgs {
            sps := strings.Split(item, ":")
            if len(sps) != 2 {
                fmt.Fprintf(os.Stdout, "\033[33mWarning: '%s' is not an effective parameter for struct field\033[m\n", item)
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
        fmt.Fprintf(os.Stderr, "\033[31mCreating interface file failed with error '%s'\033[m\n", err.Error())
    }
    defer file.Close()
    _, err = file.WriteString(tmpl.InterfaceTMPL)
    if err != nil {
        fmt.Fprintf(os.Stderr, "\033[31mWrite interface file failed with error '%s'\033[m\n", err.Error())
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