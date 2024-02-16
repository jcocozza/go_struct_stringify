# Go Struct Stringify

This package will take a struct instance and convert it into a string.

The goal is to provide a very easy way to take a struct that exists as an instance and create the code that corresponds to that struct.

```
type myStruct struct {
    A string
    B int
}

myStruct := &{A: "a", B: 1}

myStructString := StructStringify(myStruct)
fmt.Print(myStructString)
--
&myStruct{A: "a", B: 1}
```