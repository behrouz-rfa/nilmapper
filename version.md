# nilmapper
#### Version 1.0.1
* Feature: Added to ingore case sensetive if the mapper couldnt find Filed name
* support Slice copy
* you can use like this :
``` go
  // Base model
  type BaseModel struct {
      Id    int 
  }
  
  // Country model
  type Test struct {
      ID int 
  }
  src:= BaseModel {
        Id: 1
  }
  var dest  Test
  Copy(src,&dst)
```
* 2023-04-15 19:00 in ShangHai
