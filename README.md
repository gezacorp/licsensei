# LicSensei

![Build Status](https://github.com/gezacorp/licsensei/workflows/CI/badge.svg?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/gezacorp/licsensei?style=flat-square)](https://goreportcard.com/report/github.com/gezacorp/licsensei)

**A lightweight tool for verifying license headers in Go source code files.**

*The project draws inspiration and some actual code from the header-checking functionality of [Licensei](github.com/goph/licensei). It provides greater flexibility, including support for multiple configurable license header options, built-in recognition of commonly used licenses, and the ability to manage multiple allowed copyright disclaimers within each header configuration.*

Built-in SPDX license identifiers

* MIT
* Apache-2
* BSD-2-Clause
* BSD-3-Clause
* MPL-2.0
* GPL-2.0-only
* GPL-2.0-or-later
* GPL-3.0-only
* GPL-3.0-or-later

## Installation

Download the latest pre-built binary from the [Releases](https://github.com/gezacorp/licsensei/releases) page.

Alternatively, add the following code to your Makefile:

```makefile
LICSENSEI_VERSION = 0.0.1
bin/licsensei: bin/licsensei-${LICSENSEI_VERSION}
    @ln -sf licsensei-${LICSENSEI_VERSION} bin/licsensei
bin/licsensei-${LICSENSEI_VERSION}:
    @mkdir -p bin
    curl -sfL https://raw.githubusercontent.com/gezacorp/licsensei/master/install.sh | bash -s v${LICSENSEI_VERSION}
    @mv bin/licsensei $@
```

## Usage

```
Usage of licsensei:
  licsensei

Flags:
  -c, --config string   config file path
  -v, --verbosity int   log verbosity, higher value produces more output
```

## Configuration example

Place the following content into `.licsensei.yaml`.

```yaml
ignorePaths:
  - vendor
  - .gen
ignoreFiles:
  - "mock_*.go"
  - "*_gen.go"
headerConfigs:
  - copyrights:
      - "Copyright (c) :YEAR: :AUTHOR:"
    licenseTypes:
      - MIT
    authors:
      - Geza Corp and authors
    licenseTexts:
      - |
        The MIT License (MIT)

        Permission is hereby granted, free of charge, to any person obtaining a copy of this software
        and associated documentation files (the "Software"), to deal in the Software without restriction,
        including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
        and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so,
        subject to the following conditions:

        The above copyright notice and this permission notice shall be included in all copies or substantial
        portions of the Software.

        THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED
        TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
        THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
        OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
