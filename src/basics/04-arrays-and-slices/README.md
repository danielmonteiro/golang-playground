**Array** has a fixes length
**Slice** can grow or shrink
To loop through a slice:
```
for i, card := range cards {
  fmt.Println(i, card)
}
```
If we are not going to use the index variable, we have to replace by an underscore
```
for _, card := range cards {
  fmt.Println(card)
}
```

## Slice range syntax
```
fruits[startIndexIncluding:upToNotIncluding]
```
`fruits[0:2]` will get the values from index 0 and 1 (index 2 is not included).  `fruits[:2]` Has the same behavior. `fruits[2:]` will get the values from index 2 until the end of the slice

## Length of a slice
```
len(cards)
```

## Swap elements of a slice
```
cards[i], cards[j] = cards [j], card[i]
```