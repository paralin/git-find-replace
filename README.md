# git-find-replace

[Go Report Card Widget]: https://goreportcard.com/badge/github.com/paralin/git-find-replace
[Go Report Card]: https://goreportcard.com/report/github.com/paralin/git-find-replace

This is a small tool which can be used to find and replace literal strings in a
Git repository.

It uses `git grep -l -F` to search for the pattern and `strings.ReplaceAll` from
Go to replace the matches within the files.

Usage:

```
go install -v github.com/paralin/git-find-replace@main
cd ./my/git/repo
git find-replace 'SearchString' 'ReplaceString'
```

Searches and replaces in files in and below the working dir.

Ignores files ignored by .gitignore.

## License

MIT
