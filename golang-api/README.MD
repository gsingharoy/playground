# Golang API
This is a simple API server build in Go.

## Installation
Install go using brew.
```
brew update
brew install go
```
Setup your go path.
```
mkdir $HOME/work
```
Add the following lines to your `~/.bash_profile` or `~/.profile`.
```
# Add the following lines to the
export GOPATH=$HOME/work
export PATH=$PATH:$GOPATH/bin
```
Here `$HOME/work` is your workspace location. feel free to choose anything.
<br>
<br>
Move the contents of this project to the `$GOPATH/src folder`.
You should now have the project under `$GOPATH/src/golang-api`.
<br>
<br>
Fetch all the dependencies and build the project
```
go get
```
## Running
Build the project
```
go build
```
Execute the newly created executable file
```
./golang-api
```
Now go to http://localhost:8080.
