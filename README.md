# GO EPO-EPS
[![Go Report Card](https://goreportcard.com/badge/github.com/max-planck-innovation-competition/go-epo-eps)](https://goreportcard.com/report/github.com/max-planck-innovation-competition/go-epo-eps)
[![Go Reference](https://pkg.go.dev/badge/github.com/max-planck-innovation-competition/go-epo-eps.svg)](https://pkg.go.dev/github.com/max-planck-innovation-competition/go-epo-eps)

Go API Client for the European Publication Server Web Service

The European Publication Server's REST API enables access to XML, HTML, TIFF images, and PDF/A versions of European A
and B publications.

## Status

Alpha Version

**⚠️ Experimental - Not ready for production.**

## Installation

Add the package to your project via the following command:

```shell
go get github.com/max-planck-innovation-competition/go-epo-eps
```

## Usage

The following function calls can be used to retrieve the publication dates, the patent ids and the patent data.

### Get publication dates

```go
import eps
dates, err := eps.GetPublicationDates()
```

### Get patents ids of a publication dates

```go
import eps
patentIds, err := eps.GetPublicationDatePatents(date)
```

### Get patent by id

```go
import eps
patentXMLData, err := eps.GetPatentXML(patentID)
patentHTMLData, err := eps.GetPatentHTML(patentID)
patentZIPData, err := eps.GetPatentZIP(patentID)
patentPDFData, err := eps.GetPatentPDF(patentID)
```

### Transform xml data to golang struct

```go
import eps
epPatentDocumentSimple, err := eps.ProcessXMLSimple(patentXMLData)
```


## Environment

```
PROXY=http...
HTTP_PROXY=http...
```
