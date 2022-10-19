# Go Lesson

## First Dive: Run Your First Program

Go (or Golang) is an open source programming language created by Google.

What do you need to run a Go program?

    - Go compiler. (How to install)[https://go.dev/doc/install]
    - Code editor: (Goland)[https://www.jetbrains.com/go/promo/?source=google&medium=cpc&campaign=10156130801&gclid=CjwKCAiA55mPBhBOEiwANmzoQqsFRbbsB67NkuZM5EHt4qd33iv9XBcbxbnQEEfei2q0nwPT6tRI_xoCvZYQAvD_BwE] or [VSCode](https://code.visualstudio.com/)

Or you can start even simpler and run the program in the go playground. Just follow this (link)[https://goplay.tools/].

As not all your installed the Go compiler, let's jump into the playground (and please install the Go compiler by next lesson). 

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's do it")
}
```

Please open the playground in your browser and run the program.

We will be learning by doing, and you are expected to write the code during the lessons.

## Output

Try to change the output string and print "Hello world".

---spo

```go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's do it")
}
```

## Package

Now, the fun part. One of the most useful things that you can during the learning is to try to break the things.

Try to remove the very first line about package and run the program.

It seems to be not happy about the change and displays an error message:

    3:1: expected 'package', found 'import' (and 2 more errors)

Now, we can see that package name must always be provided.

First - **what is a package**?

Every Go file must belong to some package. Packages is a way to structure source code in Go. We don't need to worry about packages when we write a simple program as this one, but imagine that we have 10K lines of code. In this case, we must separate the source code into different files.

What we need to know for know is that every Go file must start with string `package $packagename`.

Now, try to change package name and call it `package somepackage` or `package usual` instead of package `main`.

The Go compiler not about it either:

No, it seems to be not happy about it and shows an error message:

    "package name must be main"

What does it mean?

In this case, we are trying to execute a program. `package main` is a special name, and we can only execute main packages in Go. So why do need other packages at all? We will cover it in a minute.

## Change import

Now what happens if I remove and import?

Now we got another error:

    ./prog.go:4:2: undefined: Println

What does it mean?

The good thing about programming is that 99% or any program is already written. Our Go program work with the browser, and the browser works with the operational system. The task of printing out a string might seem trivial at the first glance but actually it involves a lot of heavy-lifting. Luckily we don't have to worry about it. Other people already took care and wrote a function (Println)[https://pkg.go.dev/fmt#Println] that deals with the printing.

In this case, we are reusing this code by `import` statement and import function `Println` from the standard library.

If click the function name in the doc it will direct us to the (source code)[https://cs.opensource.google/go/go/+/refs/tags/go1.17.6:src/fmt/print.go;l=273] of the standard library that implements Println. We don't need to worry about how this particular function was implemented.

Here we can find familiar statement `package fmt`. Go authors named this package as `fmt` and therefore we also import it as `package fmt.`

And this is how you use non-main packages. You can also create your own package `package myOwn`, put some code here and then import it from your `package main` and use it here.

## Change function name

`func` is a keyword that declares a function. Function allows you to reuse your code.

Let's rename our function and call it `func notmain`

    runtime.main_main·f: function main is undeclared in the main package

In this case, we got something interesting - `runtime` error. Notice, that the previous error messages didn't have this prefix. We will cover the difference between the runtime error and usual error soon.

The error message is quite straightforward - package `main` must have function called `main`, so the Go program knows where to start.

## Change imported function call

Let's change `fmt.Println` to `fmt.Write`. Now try to guess what is going to happen before executing the program.

    ./prog.go:6:2: undefined: fmt.Write

The program says to us that it does not know what is `fmt.Write`. It makes sense because we are trying to import function `Write` from package `fmt`. And package `fmt` does not contain this function, so the Go program does not know what to do with this function call.

OK, you can see that even in such a small program **a lot** can get wrong. As you will be starting be prepared that a lot will get wrong :)

## Read the input

Let's we want to say hello to a user of our program. How can we read the input?

There must a library to do so. So we google it first and find that [bufio](https://pkg.go.dev/bufio) package provides the reading functionality.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What's your name? ")
	scanner.Scan()

	fmt.Println("Hello,", scanner.Text())
}
```

Now we run in a playground we will see that the program actually does not wait for the input and simply outputs empty string instead of a name.

This happens because we are trying to read from the standard input and browser does not have access to it due to security reasons.

So, this program we need execute directly at our machines with the following command:

    go run hello.go

Or simply press play button in source code editor.

Now the program asks for out input and prints out the name.

# Conclusion

In the first session, we learned how to create and execute a Go program that can interacts with a user. Not bad for a start :)




