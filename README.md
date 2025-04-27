# logging-challenge
logging challenge course

## Setup

### Virtual Machine
- For virtual machine, I'm using [orbstack](https://orbstack.dev/)
- Need to install rosetta first

```
/usr/sbin/softwareupdate --install-rosetta --agree-to-license
```

- To connect to the orbstack virtual machine, use

```sh
ssh <virtual-machine-name>@orb

ex:
ssh ubuntu-telemetry@orb
```

### Go

Follow this [instruction](https://www.cherryservers.com/blog/install-go-ubuntu) to install go 1.21

```sh
sudo apt update && sudo apt upgrade
wget https://go.dev/dl/go1.21.4.linux-amd64.tar.gz -O go.tar.gz
sudo tar -xzvf go.tar.gz -C /usr/local
echo export PATH=$HOME/go/bin:/usr/local/go/bin:$PATH >> ~/.profile
source ~/.profile
```

### Python

```sh
sudo apt update && sudo apt upgrade
sudo apt install build-essential software-properties-common -y\
sudo add-apt-repository ppa:deadsnakes/ppa
sudo apt install python3.11 -y
python3.11 --version
```

### Docker

Follow tutorial from -> [link](https://www.cherryservers.com/blog/install-docker-ubuntu)

### Docker Compose

Follow tutorial from -> [link](https://medium.com/@piyushkashyap045/comprehensive-guide-installing-docker-and-docker-compose-on-windows-linux-and-macos-a022cf82ac0b)
