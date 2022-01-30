# Structs
Defining a struct:
```
type person struct {
	firstName string
	lastName  string
}
```
There are 3 ways to declare a struct:

### 1st
```
person1 := person{"Daniel", "Monteiro"}
```
Arguments passed to `person` will follow the same order as the fields are declared in the struct.

### 2nd
```
person2 := person{firstName: "Daniel", lastName: "Monteiro"}
```
Declare a value to given field names.

### 3rd
```
var person3 person
person3.firstName = "Daniel"
person3.lastName = "Monteiro"
```
Create an empty struct. Initially the fields will have zero/default values ("" for strings, 0 for int and float, false for boolean). Then access each field and assign the value.

### Printing structs
There's a "special" way to print structs where it'll display the field names:
```
fmt.Printf("%+v", person1)
```
Will print:
```
{firstName:Daniel lastName:Monteiro}
```

## Embedded structs
It's possible to create a field in a struct that is type of another struct:
```
type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}
```
To declare it:
```
embedded := person{
  firstName: "Embedded",
  lastName:  "Structs",
  contact: contactInfo{
    email:   "embedded-structs@email.com",
    zipCode: 12345,
  },
}
```
It is possible to omit the field name in case it will have the same name as the struct:
```
type person struct {
	firstName string
	lastName  string
	contactInfo
}
```
In the above example the `person` struct will have a `contactInfo` field type of `contactInfo`


## Receiver functions for structs
It is possible to create a receiver function for a struct:
```
func (p person) print() {
	fmt.Printf("%+v", p)
}
```
Then a struct of type `person` can call the `print` function:
```
person1 := person{firstName: "Daniel", lastName: "Monteiro"}
person1.print()
```