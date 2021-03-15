package tmpl

var InterfaceTMPL = `
package interfaces

type Cybertron interface {
    Attribues() ([]string)
    Insert()    (int64, error)
    Destroy()   (bool, error)
    ToJson()    (map[string]interface{}, error)
    Update(map[string]interface{}) (int64, error)
}
`