% Challenge Time
% Tom Arrell
% January 13

---

# Last week...

We had a look at some concepts partially or wholly unique to Go.

- Defer
- Multiple return values
- Interfaces
- Goroutines
- Channels

---

# This week

We're going to go through a set of challenges to practice what we've learned.

We'll spend 10 minutes on each challenge, and then solve it together.

---

# Setup

Each challenge will ask you to write a function. The setup for each one should look something like this...

```go
package main

func solution() {
  // your code here...
}

func main() {
  solution()
}
```

---

# Challenge 1: Counting Sheep

Consider a slice of sheep where some sheep may be missing from their place. Write a function `countSheep` that counts the number of sheep present in the slice (true means present) and returns the count.

e.g.

```go
var sheep = []bool{
  true,  true, false, true,
  false, true, false, true,
  false, true, true,  true,
  true,  true, true,  false,
}

func main() {
  fmt.Println(countSheep(sheep)) // 11
}
```

---

# Challenge 2: Name to Initials

Write a function `getInitials` which takes a single string, which is a person's first and last name, and returns a string of the person's initials.

e.g. *"John Smith"* => *"J.S"

```go
func main() {
  fmt.Println(getInitials("John Smith")) // "J.S"
  fmt.Println(getInitials("Bob Ross")) // "B.R"
  fmt.Println(getInitials("Luke Skywalker")) // "L.S"
}
```

---

# Challenge 3: Paper, Scissors, Rock

Write a function `psr` which takes two strings, giving the move of two different players, and return a string with who won.

e.g.

```go
func main() {
  fmt.Println(psr("paper", "scissors")) // "Player 2 wins"
  fmt.Println(psr("rock", "scissors")) // "Player 1 wins"
  fmt.Println(psr("paper", "rock")) // "Player 2 wins"
}
```

---

# Challenge 4: Reverse a string

Write a function `rev` which takes a string, and returns a version of the string with all characters in reverse order.

e.g.

```go
func main() {
  fmt.Println(rev("use the force")) // "ecrof eht esu"
  fmt.Println(rev("hello world")) // "dlrow olleh"
}
```

---

# Challenge 5: Check for Anagrams

Write a function `isAnagram` which takes two strings, and returns true if all the letters in the first words can be found in the second word.

e.g.

```go
func main() {
  fmt.Println(isAnagram("casper", "parsec")) // true
  fmt.Println(isAnagram("dog", "cat")) // false
}
```

---

*lesson 5, fin*

If you had any trouble, now is the time to ask for help!

**Questions?**
