# cybertron
> 通过[cobra](https://github.com/spf13/cobra)实现的CLI，用来生成模板go struct文件
> CLI with [cobra](https://github.com/spf13/cobra), to generate go struct file

#### 1. 安装/install
```
go get -u github.com/grayxiaoxiao/cybertron@v1.1.1
```
#### 2. 使用/usage
##### 2.1 创建Customer/Create Customer
> `cybertron cube new customer`，生成的文件如下/The struct file is:
```
type Customer struct {
}
func (obj Customer) Attributes() []string {
  // code here......
}
func (obj Customer) Insert() (obj_id int64, err error) {
  // code here......
}
// code here.......
```
##### 2.2 创建带字段的Product/Create Product with fields
> `cybertron cube new product id:int64 name:string serial_number:string description`, 生成的文件如下/The struct file is:
```
type Product struct {
  Id int64 `column_name:"id" json:"id"`
  Name string `column_name:"name" json:"name"`
  SerialNumber string `column_name:"serial_number" json:"serial_number"`
  // description without data type, so generate faild
}
```
##### 2.3 创建带路劲的Order/Create Order with path
> `cybertron cube new models/structs/order`，生成CurrentPath/models/structs/order.go
