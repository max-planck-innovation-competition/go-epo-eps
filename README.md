# Go API client for the European Publication Server web service

The European Publication Server's REST API enables access to XML, HTML, TIFF images, and PDF/A versions of European A
and B publications.

## Status

Alpha Version

## Installation

Put the package under your project folder and add the following in import:

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
patentXMLData, err := GetPatentXML(patentID)
patentHTMLData, err := GetPatentHTML(patentID)
patentZIPData, err := GetPatentZIP(patentID)
patentPDFData, err := GetPatentPDF(patentID)
```