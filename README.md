# go-measure-things
An example project how to export application metrics to Prometheus


## Planned talks

30. January 2019 
- Chemnitz Cloud Native Meetup #2 at Staffbase HQ https://www.meetup.com/de-DE/Chemnitz-Cloud-Native-Meetup/events/257370642/

16. or 17. March 2019 (not yet accepted)
- Chemnitz CLT 2019 as German talk "Prometheus in Aktion: Go, measure things!" 


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

Go to the `src` directory and execute:

```
go build main.go
go build sleep.go
```

to build all files.