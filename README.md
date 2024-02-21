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

## Motivation

You're writing some integration tests for your code. You trust things to work-- these tests are more of a future proofing endeavor.
Topping it off, you have to load in some complex objects from the database. Structs with sub-structs, lists of structs, custom types etc. It's going to manually create these objects to verify things are working properly.

Here's were this package comes in, just use the `StructStringify()` function to turn a struct instance (that you just loaded in from your database, for example) into a string. Then all you need to do is print that string, and it will be ready for you to copy and paste where ever you please.