# vdf2json
Written in Go, change Valve Data Format string into standard json format.
# Usage
Use this in project's directory
```
go get -u github.com/Purple-CSGO/vdf2json-go
```
Then use this method in Go
```
jsonData := vdf2json_go.ToJson(vdfData)
```
Notice that vdfData & jsonData are both `string` type