# ws-pdf-publish: WebSite PDF Publish

@author Carles Clavero i Matas - carles.clavero@gmail.com

@date 2021-11-28

![Go Version](https://img.shields.io/badge/go-ver%3E=1.16-informational)
![Docker](https://img.shields.io/badge/docker-ver%3E=20-informational)
![License](https://img.shields.io/badge/license-MIT-green)
[![Go Report Card](https://goreportcard.com/badge/github.com/cclavero/ws-pdf-publish)](https://goreportcard.com/report/github.com/cclavero/ws-pdf-publish)
![Tests](https://img.shields.io/badge/tests-passed-green)
![Coverage](https://img.shields.io/badge/coverage-80%25-green)

WebSite PDF Publish: Simple command line tool to publish a set of HTML pages to PDF.

Current version: 1.0-alpha




--- TEMPORAL

- https://goreportcard.com/report/github.com/cclavero/ws-pdf-publish

- Example:


[![Go Report Card](https://goreportcard.com/badge/github.com/spf13/viper?style=flat-square)](https://goreportcard.com/report/github.com/spf13/viper)


--- TEMPORAL

- https://raw.githubusercontent.com/spf13/viper/master/README.md

--- TEMPORAL: REVISAR

Diari personal utilitzant Hugo per les Notes d'inxes

## Recursos

- https://github.com/spf13/viper
- https://github.com/onsi/ginkgo
- https://github.com/pdfcpu/pdfcpu
- https://wkhtmltopdf.org



## Prerequisits

- Tenir docker instal·lat:

```bash
$ docker version
Client:
 Version:           20.10.7
...
```

- Tenir golang instal·lat, versió 1.16 o superior:

```bash
$ go version
go version go1.16 linux/amd64
```

## Notes i configuració

1. Ubicació de les entrades i imatges per les Notes d'inxes

- Els fitxers 'md' amb el contingut de les futures URLs per les notes, es troben al directori '/docs/content/docs'.
- Les diferents imatges dels fitxers 'md', es troben al directori '/docs/static/img'.
- El menú és automàtic. Només cal afegir un nou fitxer 'md' per afegir-lo a les notes.

2. Generació del fitxer PDF final i llistat de URLs

El procés de generació del fitxer PDF amb totes les URLs de les Notes d'inxes utilitza la configuració definida al fitxer '/publish-pdf/publish-urls.yaml':

```yaml
publish:
  file: notes-inxes.pdf # Name of the final PDF file
  urls: # List of URLs to process
    - url: http://notes-inxes:1313/
      file: portada.pdf
    - url: http://notes-inxes:1313/docs/esquemes-generals
      file: esquemes-generals.pdf
    - url: http://notes-inxes:1313/docs/do-met-gallec/
      file: do-met-gallec.pdf
    - url: http://notes-inxes:1313/docs/do-met-escoces/
      file: do-met-escoces.pdf
    - url: http://notes-inxes:1313/docs/sol-met-gallec/
      file: sol-met-gallec.pdf
  wkhtmltopdfParams: --print-media-type --margin-top 20mm --margin-bottom 20mm # Parameters for the wkhtmltopdf utility    
```

Per tant, si s'afegeix o s'elimina un fitxer 'md', s'haurà d'actualitzar aquest fitxer.

3. Directori de generació dels fitxers PDF per cada URL i el fitxer PDF final

El directori a on es generen el fitxers PDF per cada URL (parcials) i el fitxer PDF final amb tots els continguts de les Notes d'inxes es el directori '/out'.

Tots els fitxers es regeneren en cada execució.

## Utilització

Totes les tasques definides al Makefile:

1. Obtenir ajuda de les tasques

```bash
$ make
...
```

2. Generar el docker amb el servidor web Hugo

```bash
$ make build
...
$ docker images | grep 'notes-inxes'
notes-inxes    1.0          0712a14d81d5   34 seconds ago   862MB
```

3. Eliminar la imatge de docker amb el servidor web Hugo

```bash
$ make clean
...
Untagged: notes-inxes:1.0
Deleted: sha256:21a32a6395d01a582d3723e2dcb430eaa035dd1ce05efba536f1349e3202ceb4
Deleted: sha256:b7af55c6fae99b16fafd76ae707bacfe7fbadc36185e031ad4260072aa47b938
```

4. Arrencar el docker amb el servidor web Hugo

```bash
$ make start
...
3f3b18a7609e0c3f595a02581072b408fc34c7bfc8cc09818b6f71055b2aac5a
$ docker logs -f notes-inxes
Start building sites …
...
Press Ctrl+C to stop
```

Un cop arrencat el servidor web Hugo, utilitzar el navegador per accedir a: 

- http://localhost:1313/

5. Publicar les Notes d'inxes com a fitxer PDF

Si el servidor web Hugo està arrencat, podem publicar els continguts de les Notes d'inxes com a un sol fitxer PDF (veure apartat 'Notes i configuració'):

```bash
$ make publish
...
- Starting
...
- Done
```

Un cop finalitzada la publicació, pots obtenir el fitxer PDF al directori '/out/notes-inxes.pdf'.

6. Aturar el docker amb el servidor web Hugo

```bash
$ make stop
...
```



--- TEMPORAL

$ make == make help

$ make clean

$ make ci
$ make test
$ make lint
$ make build

$ make install

--- TEMPORAL

$ go test -v -failfast -count=1 -tags="test" .

