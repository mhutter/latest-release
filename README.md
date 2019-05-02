# latest-release

[![Build Status](https://travis-ci.com/mhutter/latest-release.svg?branch=master)](https://travis-ci.com/mhutter/latest-release)

Check the latest release of a GitHub repository.

## Usage

    $ latest-release [<flags>] repository

    ARGUMENTS:
      repository
            Repository to check, either in the form of ORG/REPO or just REPO,
            in which case it will be expanded to REPO/REPO

    FLAGS:
      -P    ONLY include prereleases
      -p    Include prereleases
      -v    Print version and exit


## Installation

**Homebrew tap** (macOS)

    brew install mhutter/tap/latest-release


**go get** (any OS)

    go get -u github.com/mhutter/latest-release


Or [download a pre-built binary](https://github.com/mhutter/latest-release/releases).



## License

MIT (see LICENSE)


---
> [Manuel Hutter](https://hutter.io) -
> GitHub [@mhutter](https://github.com) -
> Twitter [@dratir](https://twiter.com/dratir)
