ghr
====

[![GitHub release](http://img.shields.io/github/release/tcnksm/ghr.svg?style=flat-square)][release]
[![Wercker](http://img.shields.io/wercker/ci/54393fe184570fc622001411.svg?style=flat-square)][wercker]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)][godocs]

[release]: https://github.com/tcnksm/ghr/releases
[wercker]: https://app.wercker.com/project/bykey/a181c474f1e25e1870d0ba387723046b
[license]: https://github.com/tcnksm/ghr/blob/master/LICENSE
[godocs]: http://godoc.org/github.com/tcnksm/ghr

`ghr` is a simple tool to create GitHub Release and upload artifacts by only one command. `ghr` uploads multiple artifacts in parallel. 

## Demo

This demo shows simple usage of `ghr`. It creates GitHub Release page for version `v1.0.0` and uploads cross-compiled golang binaries there. 

![](doc/ghr.gif)

You can see release page [here](https://github.com/tcnksm/ghr-demo/releases/tag/v1.0.0).

## Usage

Run it in your project directory,

```bash
$ ghr [option] <tag> <artifacts>
```

You need to set `GITHUB_TOKEN` environmental variable:

```bash
$ export GITHUB_TOKEN="....."
```

Or you can set it in `github.token` in gitconfig:

```bash
$ git config --global github.token "....."
```

Enviromental variable takes priority over gitconfig value.

### GitHub Enterprise

You can use `ghr` for GitHub Enterprise. Change API endpoint via the enviromental variable.

```bash
$ export GITHUB_API=http://github.company.com/api/v3
```

## Example

To upload all package in `pkg` directory with tag `v0.1.0`

```bash
$ ghr v0.1.0 pkg/
--> Uploading: pkg/0.1.0_SHASUMS
--> Uploading: pkg/ghr_0.1.0_darwin_386.zip
--> Uploading: pkg/ghr_0.1.0_darwin_amd64.zip
--> Uploading: pkg/ghr_0.1.0_linux_386.zip
--> Uploading: pkg/ghr_0.1.0_linux_amd64.zip
--> Uploading: pkg/ghr_0.1.0_windows_386.zip
--> Uploading: pkg/ghr_0.1.0_windows_amd64.zip
```

Or if you want to replace artifact which is already uploaded:

```bash
$ ghr --replace v0.1.0 pkg/
```

## Options

You can set some options:

```bash
$ ghr \
    -t <token> \       # Set Github API Token
    -u <username> \    # Set Github username
    -r <repository> \  # Set repository name
    -c <commitish> \   # Set target commitish, branch or commit SHA
    -p <num> \         # Set amount of parallelism (Default is number of CPU)
    --replace \        # Replace asset if target is already exists
    --delete \         # Delete release and its git tag in advance if it exists
    --draft \          # Release as draft (Unpublish)
    --prerelease \     # Crate prerelease
    <tag> <artifacts>
```

## Install

If you are OSX user, you can use [Homebrew](http://brew.sh/):

```bash
$ brew tap tcnksm/ghr
$ brew install ghr
```

If you are in another platform, you can download binary from [release page](https://github.com/tcnksm/ghr/releases) and place it in `$PATH` directory.

## Integration with CI-as-a-Service

You can integrate ghr with CI-as-a-Service to release your artifacts after test passed. It's very easy to provide latest build to your user continuously.

See [Integrate ghr with CI as a Service](https://github.com/tcnksm/ghr/wiki/Integrate-ghr-with-CI-as-a-Service) page.

## VS.

- [aktau/github-release](https://github.com/aktau/github-release) - `github-release` can also create and edit releases and upload artifacts. It has many options. `ghr` is a simple alternative. And `ghr` will parallelize upload artifacts.

## Contribution

1. Fork ([https://github.com/tcnksm/ghr/fork](https://github.com/tcnksm/ghr/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `make test` command and confirm that it passes
1. Run `gofmt -s`
1. Create new Pull Request

You can get source with `go get`:

```bash
$ go get -d github.com/tcnksm/ghr
$ cd $GOPATH/src/github.com/tcnksm/ghr
$ make install
```

`ghr` uses vendoring feature for some third party packages. If you use `go1.5`, then set `GO15VENDOREXPERIMENT=1`. 

## Author

[Taichi Nakashima](https://github.com/tcnksm)
