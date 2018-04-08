# Introduction

This is an automatic script for converting Microsoft docs from [API reference of Azure SDK for .NET](https://docs.microsoft.com/en-us/dotnet/api/overview/azure/?view=azure-dotnet) to [Dash](https://kapeli.com/dash), an API documentation browser for macOS and iOS.

We found that all the sub-link of API reference in the [API reference of Azure SDK for .NET](https://docs.microsoft.com/en-us/dotnet/api/overview/azure/?view=azure-dotnet) page can be get by visiting the Json file, [Azure .NET API list](https://docs.microsoft.com/api/apibrowser/dotnet/namespaces?moniker=azure-dotnet&api-version=0.2 ). 

We also use [dashing](https://github.com/technosophos/dashing), a Dash generator script for any HTML files to simply the building process. The project, dashing, can generate the SQL database file and other files and structure of file system which a Dash docset should contains. Dashing also provides CSS selectors and other useful tools, contributing to building a complex docset.

When we get all Microsoft docs HTML files in a flat file structure, we can simply use dashing to build a custom docset which can be imported into Dash app later.

# Dependencies

[simplejson](github.com/bitly/go-simplejson)

Use `go get -u github.com/bitly/go-simplejson` to get go-simplejson in your local GOROOT dir.

# Usage

Golang environment is required. If you have Go 1.4 or later, simply use `go run` to start the download process. When finished, you'll see many HTML files in the same directory with `main.go`.

Then use [dashing](https://github.com/technosophos/dashing) to build the docset. Full documentation of dashing can be seen from its [Github repo](https://github.com/technosophos/dashing/blob/master/README.md). We can only concentrate on the `dashing.json`. We need to enable external JS by modifying `AllowJS` entry in `Info.plist` in the docset. You get more information in [enableJavascript](https://kapeli.com/docsets#enableJavascript). Or simply use my own version of [dashing](https://github.com/yzs981130/dashing).

# Attention

Avoid running it in `$HOME` or other important folder. In case, use `rm *.html` to clean up.


# Performance problem
Because MS docs website will block connection when too frequently, so I only use one single thread to get all the HTML files. It's expected to finish in hours when there's a pool Internet connection.

# Future work

- Optimize searching function.

- Use multiple filename map to improve concurrency.