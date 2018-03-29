# Introduction

This is an automatic script for converting Microsoft docs from [API reference of Azure SDK for .NET](https://docs.microsoft.com/en-us/dotnet/api/overview/azure/?view=azure-dotnet) to [Dash](https://kapeli.com/dash), an API documentation browser for macOS and iOS.

We found that all the sublink of API reference in the [API reference of Azure SDK for .NET](https://docs.microsoft.com/en-us/dotnet/api/overview/azure/?view=azure-dotnet) page can be get by visiting the Json file, [Azure .NET API list](https://docs.microsoft.com/api/apibrowser/dotnet/namespaces?moniker=azure-dotnet&api-version=0.2 ). 

We also use [dashing](https://github.com/technosophos/dashing), a Dash generator script for any HTML files to simply the building process. The project, dashing, can generate the SQL database file and other files and structure of file system which a Dash docset should contains. Dashing also provides CSS selectors and other useful tools, contributing to building a complex docset.

When we get all Microsoft docs HTML files in a flat file structure, we can simply use dashing to build a custom docset which can be imported into Dash app later.

# Usage

If you have Go 1.4 or later, simply use `go run` to start the download process. When finished, use dashing to build docset.

# Future work

- Optimize searching function.

- Use multiple filename map to improve concurrency.