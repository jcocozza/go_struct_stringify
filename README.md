# Go Struct Stringify

This package will take a struct instance and convert it into a string.

```
type myStruct struct {
    A string
    B int
}

myStruct := &{ A: "a", B: 1}

myStructString := StructStringify(myStruct)
fmt.Print(myStructString)
--
&myStruct{ A: "a", B: 1}
```