# ORN KIT
This kit support :
```
1. Golang version 1.13.x or latest
2. Python version 3.7.x or latest
3. NodeJS latest
```

For futher information:
```
Sofyan Saputra | sofyansaputrawn48@gmail.com
```

## Foreword
---
### Install Golang
**Darwin**
```
wget https://dl.google.com/go/go1.14.3.darwin-amd64.pkg
```
double click to install

**Linux**
```
curl -fLo /usr/local https://dl.google.com/go/go1.14.3.linux-amd64.tar.gz
```
export path and write to .bashrc
```
export PATH=$PATH:/usr/local/go/bin
```

### Getting Started
```
mkdir ~/.orn/
cp config.example ~/.orn/
```
edit your config


### Installing

#### From Binary

```
curl -fLo /usr/local/bin/orn https://github.com/orn-id/wa-cli/blob/master/bin/orn.$OS-$ARCH
```

#### From Source
**Darwin**
```
make build-darwin ARCH=$ARCH
make install 'or use sudo' sudo make install
make clean
```
**Linux**
```
make build-linux ARCH=$ARCH
make install 'or use sudo' sudo make install
make clean
```
**Windows**
```
This tools not support for windows
```


Usage tools see for help
```
NAME:
   orn - orn [command]

USAGE:
   orn [global options] command [command options] [arguments...]

VERSION:
   0.1.1

AUTHOR:
   Service API <api.team@orn.com>

COMMANDS:
   migrate  migrate [option]
   seed     seed [option]
   ssm      ssm [option]
   swagger  swagger [option]
   builder  builder [option]
   setup    setup [option]
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE  Load environtment config from FILE
   --help, -h              show help
   --version, -v           print the version
```

### Bulider
Builder tool build your service in golang, python and nodejs 
#### Golang
#### Python
#### NodeJS

