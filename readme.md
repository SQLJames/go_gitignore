# go_gitignore

go_gitignore is a go program that reaches out to https://github.com/github/gitignore and pulls the appropriate gitignore file for your language.


## Usage

```goalng
go build 
.\main.exe python
```

This will return a results folder which has all the files for your language in the repo.
Note this will only return files that are named with that language. So, for go, if you pass in golang the files will not be returned as the file is named go.gitignore.