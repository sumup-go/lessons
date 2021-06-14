# Introduction
In preparation for the first Go lesson, please make sure to prepare the following:

* Laptop

We will be using a laptop in order to do practical exercises. If you don't have a laptop, you'll still be able to follow along.

Please also follow the below instructions depending on the operating system for the laptop you have.

## MacOS
1. Download and install [VSCode](https://code.visualstudio.com/download)
2. Download and install the [Go Compiler](https://golang.org/doc/install?download=go1.13.4.darwin-amd64.pkg)
3. Open the terminal application `Applications > Terminal` and run the following command:

```
mkdir ~/go && echo "export GOPATH=\"$HOME/go\"" >> ~/.bashrc;
```

Done!

## Windows
1. Download and install [VSCode](https://code.visualstudio.com/download)
2. Download and install the [Go Compiler](https://golang.org/doc/install?download=go1.13.4.windows-amd64.msi)
4. Create folder at `C:\go`
5. Right click on "Start" and click on "Control Panel". Select "System and Security", then click on "System"
6. From the menu on the left, select the "Advanced systems settings"
7. Click the "Environment Variables" button at the bottom
8. Click "New" from the "User variables" section
9. Type GOPATH into the "Variable name" field
10. Type `C:\go` into the "Variable value" field
11. Click "OK"

Done!

## Linux
1. Download and install [VSCode](https://code.visualstudio.com/download)
2. Download and install the [Go Compiler](https://golang.org/doc/install?download=go1.13.4.linux-amd64.tar.gz)
3. Open your terminal and run the following command:

```
mkdir ~/go && echo "export GOPATH=\"$HOME/go\"" >> ~/.bashrc;
```

Done!

# Further Resources
If you would like to prepare a little ahead of time, definitely check out the tour of Go. It provides a great introduction to the language in a nice, interactive way.

* [tour.golang.org/welcome/1](https://tour.golang.org/welcome/1)
