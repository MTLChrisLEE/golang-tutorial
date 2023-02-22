#Packaging and Modules

 * If you import a module, it has to be used.
 * Prior to go 1.11, everything had to be in source directory (usually ~/go)
 * Now, we have modules (3rd party library)
 * Even if we have a module in ~/go/pkg/mod, we cannot import/find it. 

 * To fix it, we run 
```
go mod init example,com/packages
```
 This generates a mod file (go.mod)


 * To add module requirements and sum, we run

```
go mod tidy
```
This generates go.sum that represents chain of dependency

## Creating own package

 * create a folder
   * folder name = package name
 * We have to use the entire url
   * "example.com/packages/util"
   * "example.com/packages" part is coming from the command "go mod init example.com/packages" or "module example.com/packages" in go.mod file 

 
  
  









