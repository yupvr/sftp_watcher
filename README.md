# SFTP watcher

With help of this repo, you can build a simple SFTP server and notification system, by combining two Docker containers.
To run those containers, you should use [docker-compose](https://docs.docker.com/compose/):

- service `sftp` runs the SFTP server;
- service `watcher` runs service that monitors the directory served by service `sftp` for new files and prints their path on host machine to *STDOUT*.

## Requirements
- [Docker Engine](https://docs.docker.com/engine/) version 17.06.0+
- [Docker Compose](https://docs.docker.com/compose/) version 1.14.0+

## Usage

Define path to upload directory on host machine and run docker-compose from curent directory:

``` bash
$ export UPLOAD_DIR=/home/bob/impraise/upload && \
	docker-compose up
```
