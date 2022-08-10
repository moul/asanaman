# asanaman

:smile: asanaman

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/moul.io/asanaman)
[![License](https://img.shields.io/badge/license-Apache--2.0%20%2F%20MIT-%2397ca00.svg)](https://github.com/moul/asanaman/blob/main/COPYRIGHT)
[![GitHub release](https://img.shields.io/github/release/moul/asanaman.svg)](https://github.com/moul/asanaman/releases)
[![Docker Metrics](https://images.microbadger.com/badges/image/moul/asanaman.svg)](https://microbadger.com/images/moul/asanaman)
[![Made by Manfred Touron](https://img.shields.io/badge/made%20by-Manfred%20Touron-blue.svg?style=flat)](https://manfred.life/)

[![Go](https://github.com/moul/asanaman/workflows/Go/badge.svg)](https://github.com/moul/asanaman/actions?query=workflow%3AGo)
[![Release](https://github.com/moul/asanaman/workflows/Release/badge.svg)](https://github.com/moul/asanaman/actions?query=workflow%3ARelease)
[![PR](https://github.com/moul/asanaman/workflows/PR/badge.svg)](https://github.com/moul/asanaman/actions?query=workflow%3APR)
[![GolangCI](https://golangci.com/badges/github.com/moul/asanaman.svg)](https://golangci.com/r/github.com/moul/asanaman)
[![codecov](https://codecov.io/gh/moul/asanaman/branch/main/graph/badge.svg)](https://codecov.io/gh/moul/asanaman)
[![Go Report Card](https://goreportcard.com/badge/moul.io/asanaman)](https://goreportcard.com/report/moul.io/asanaman)
[![CodeFactor](https://www.codefactor.io/repository/github/moul/asanaman/badge)](https://www.codefactor.io/repository/github/moul/asanaman)

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/moul/asanaman)

## Usage

[embedmd]:# (.tmp/usage.txt console)
```console
foo@bar:~$ asanaman
USAGE
  asanaman [global flags] <subcommand> [flags] [args]

SUBCOMMANDS
  me              
  workspace-list  
  project-list    

FLAGS
  -cache-path .asanaman-cache  cache path
  -debug false                 debug mode
  -domain ...                  Asana workspace
  -json false                  return JSON output when possible
  -token ...                   Asana token
exit status 1
```

## Install

### Using go

```sh
go get moul.io/asanaman
```

### Releases

See https://github.com/moul/asanaman/releases

## Contribute

![Contribute <3](https://raw.githubusercontent.com/moul/moul/main/contribute.gif)

I really welcome contributions.
Your input is the most precious material.
I'm well aware of that and I thank you in advance.
Everyone is encouraged to look at what they can do on their own scale;
no effort is too small.

Everything on contribution is sum up here: [CONTRIBUTING.md](./.github/CONTRIBUTING.md)

### Dev helpers

Pre-commit script for install: https://pre-commit.com

### Contributors ✨

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg)](#contributors)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="http://manfred.life"><img src="https://avatars1.githubusercontent.com/u/94029?v=4" width="100px;" alt=""/><br /><sub><b>Manfred Touron</b></sub></a><br /><a href="#maintenance-moul" title="Maintenance">🚧</a> <a href="https://github.com/moul/asanaman/commits?author=moul" title="Documentation">📖</a> <a href="https://github.com/moul/asanaman/commits?author=moul" title="Tests">⚠️</a> <a href="https://github.com/moul/asanaman/commits?author=moul" title="Code">💻</a></td>
    <td align="center"><a href="https://manfred.life/moul-bot"><img src="https://avatars1.githubusercontent.com/u/41326314?v=4" width="100px;" alt=""/><br /><sub><b>moul-bot</b></sub></a><br /><a href="#maintenance-moul-bot" title="Maintenance">🚧</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors)
specification. Contributions of any kind welcome!

### Stargazers over time

[![Stargazers over time](https://starchart.cc/moul/asanaman.svg)](https://starchart.cc/moul/asanaman)

## License

© 2021   [Manfred Touron](https://manfred.life)

Licensed under the [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)
([`LICENSE-APACHE`](LICENSE-APACHE)) or the [MIT license](https://opensource.org/licenses/MIT)
([`LICENSE-MIT`](LICENSE-MIT)), at your option.
See the [`COPYRIGHT`](COPYRIGHT) file for more details.

`SPDX-License-Identifier: (Apache-2.0 OR MIT)`
