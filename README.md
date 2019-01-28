# go-measure-things
An example project how to export application metrics to Prometheus


## Planned talks

*30. January 2019* 
- [Chemnitz Cloud Native Meetup #2](https://www.meetup.com/de-DE/Chemnitz-Cloud-Native-Meetup/events/257370642/) at Staffbase HQ 

*16. or 17. March 2019 (not yet accepted)*
- Chemnitz [CLT 2019](https://chemnitzer.linux-tage.de/2019/) as German talk "Prometheus in Aktion: Go, measure things!" 


# Compiling

## Documentation

Go to the `doc` folder and execute either:

```
pdflatex pres.tex
```
or 

```
latexmk -pdf pres.tex
```

to compile the Latex source into a PDF.


## Building

Install [GOLang](https://golang.org/doc/install) 1.11 or newer.
Just execute:

```
go build
```

to build all files.

Now you can start the binary using the executable file: 
```
./go-measure-things
```

An alternative is starting the program directly via go after compiling it
```
go run main.go sleep.go greetings.go
```


## Running Prometheus
Change the target ip address `192.168.0.4` in the file `prometheus-data/prometheus.yml`
to your local hosts IP address which is reachable via docker. 

```
chmod 777 -R prometheus-data
docker run -p 9090:9090 \
  --name prom \
  --volume $(pwd)/prometheus-data:/prometheus-data:rw \
  prom/prometheus \
  --config.file=/prometheus-data/prometheus.yml \
  --storage.tsdb.path=/prometheus-data/data
```

Now you can access Prometheus via http://localhost:9090/

Get logs 
```
docker logs -f prom
```

## Grafana

Simply run the following command to start Grafan locally in a docker container.

```
docker run \
  -d \
  -p 3000:3000 \
  --name=grafana \
  -e "GF_SERVER_ROOT_URL=http://localhost" \
  -e "GF_SECURITY_ADMIN_PASSWORD=secret" \
  grafana/grafana
```

## Metrics

* Average sleep time `things_sleep_requested_seconds_sum / things_sleep_requested_seconds_count`
* Personal Greetings `things_sent_greetings_personal_total`
