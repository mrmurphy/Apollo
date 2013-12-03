# Apollo
A tool to make docker on as easy to use as possible on OS X.

Apollo tries to minimize the difficulty possible when developing on a Virtual Machine.

Here's a taste of how it works:
```
# Specify the Vagrantfile to be used by vagrant commands:
export VAGRANT_CWD="path/to/Vagrantfile"

# Launch the VM, and mount its filesystem over sshfs.
apollo up

# Run an interactive shell on the VM
apollo i

# Navigate to the mounted directory:
cd ~/apollo

# Run a command on the VM without launching an interactive shell
# Apollo runs the command on the remote machine in the directory
# that corresponds to the location of the shell in the mounted
# directory.
# For example, `pwd` from within ~/apollo/foobar on the host will
# return "/home/vagrant/foobar" on the VM
apollo ls -l ~
```

---

Apollo is under heavy development. Currently being used and deveoped simultaneously in production at [Space Monkey Inc.](http://www.spacemonkey.com) by Murphy Randle.

Expect a cleaner codebase and further documentation in the future.

---

# Installation
## Building
In order to build the executable, a *go* development environment is needed. Here are instructions for getting started with *golang*: [http://golang.org/doc/install](http://golang.org/doc/install)

Also, [homebrew](http://brew.sh/) is recommended for installing the operational dependencies.

After golang is installed and set up, Apollo can be installed like this:

```
go get github.com/murphyrandle/apollo
cd $GOROOT/src/github.com/murphyrandle/apollo
go build
```

## Dependencies
**Apollo** Is meant to be run on OS X and connecto to a [Virtualbox](https://www.virtualbox.org/) image managed by [Vagrant](http://www.vagrantup.com/)

The following are runtime dependencies:

- Vagrant
- Virtualbox
- sshfs
- ssh-copy-id

```
go get github.com/murphyrandle/apollo
cd $GOROOT/src/github.com/murphyrandle/apollo
brew bundle
# Read and follow the instructions in Brewfile!!
```

# Notes on usage

 - Apollo assumes that the `VAGRANT_CWD` environment variable is set, so that `vagrant ssh` can be called from any directory.
 - At the current time, the directory for the VM mount: `~/apollo` is hard-coded into the source and must be created by hand before running Apollo. This will be configured by a flag or ENV variable in future releases.
 - Apollo has only been tested on Mac OS X. Windows is not supported (unless you want to send a pull request!), and Linux is untested.
 - Apollo *should* work with any VM back-end supported by Vagrant. 

# License

The MIT License (MIT)

Copyright (c) 2013 (Murphy) Jackson Randle

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
