If a function does not return any value it can be declared as `func newCard()` but if it returns a value, the type of it has to be declared `func newCard() string`
If many files belong to the same package, the functions declared on them are "shared" between all the files (because they are part of the same package)

## Receiver functions
Receiver functions are functions attached to a given type
```
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
```
In the above example, all variables of type `deck` will have the `print()` function

## Multiple return values
It is possible to return multiple values in a function, annotating the type of each value in the function declaration and adding multiple values in the return statement separated by a coma:
```
func getFruits() (string, string) {
	return "apple", "banana"
}
```