# pdf-printer

A super simple golang POC for a html to pdf web service with an HTTP api and build on wkhtml2pdf as a backend.

## Setup

### Requirements
 - A copy of the golang 1.18 for linux zipped tarball. Download from the official site [here](https://go.dev/dl/go1.18.linux-amd64.tar.gz).
 - Docker.
 - `curl` if you want to use the provided client.sh script to test the service.

### Build & Run the Service
 - Copy the downloaded golang for linux tarball to the project root.
 - Run the following commands from the project root:

```shell
docker build -t printer .
docker run -d printer
```

## Muck about with it a bit
 - Run `docker ps` now and you should see output resembling the follwing snippet

```
CONTAINER ID   IMAGE            COMMAND                  CREATED          STATUS          PORTS     NAMES
616916e3d438   printer:latest   "/bin/sh -c /pdf-pri…"   15 minutes ago   Up 15 minutes             frosty_heisenberg
```
 - Use the provided `client.sh` script in the project root to test the service like so
```shell
./client.sh www.isitchristmas.com
```
 - Then open the example.pdf file it (hopefully) creates in the project root and unless it's christmas it should look like this:

![image](https://user-images.githubusercontent.com/38912305/162385205-4a3c7e2c-9a43-4b4b-90fe-4ef393862920.png)
