# Pointers in Go

### Some "good to know" facts about Go lang.
- When we call a function, the arguments are copied, instead of the same thing passed.
- Now, why does it matter ? it's pretty simple actully.
- if you have two copies of the same thing, and change one of them and not the other. This will cause inconsistency because the value will keep changing, depending on the variable being used.
- now, the reference of both the variables is same i.e. their name is going to be same, but the values are different. You can differntiate them only by the scope they are present in.
```
func main () {
	number := 10

	changeNum(number)
}

func changeNum(number int){
	number += 10
}
```
- The varaiable "number" in main function is 10 and in function changeNum, there is a copy of it with value 20. The only way to differentiate them is to check their "scope" or simply, "where are they present". For instance considering the above example, we can differentiate them as "number in main function" and "number in changeNum" function.
- But, what if want to use the changed number in main ? doesn't matter what we do with the number, eventually everything should be updated in the main function to be further used right ?
- There is one more thing different between both numbers, their "memory address".
```
func main () {
	number := 10
	fmt.Printf("The number present in main has memory address: %v", &number)
	changeNum(number)
}

func changeNum(number int){
	fmt.Printf("The number present in changeNum function has memory address: %v")
	number += 10
}
```
Output (example)
```
The number present in main has memory address: 0x42344598
The number present in changeNum function has address : 0x42344621
```
- As you can see, they are different addresses ! therefore if I want to change any one of them, I can, if I had their memory location.
- This is where pointers come in. They are variable who so insted of storing the value (reference), store the memory address.
```
number := 10
numAddr := &number
```
Output:
```

```