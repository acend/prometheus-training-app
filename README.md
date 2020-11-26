# Prometheus sample application

## Usage

### Start the server
Download a release from https://github.com/rotscher/prometheus-sample-app/releases and unpack the archive.
Then start the sample app as follows:

```
./prometheus-sample-app server
```

### Browse metric endpoint

```
curl http://localhost:8080/metrics
```

To demonstrate the promtool linter you can execute the following command:

```
curl -s http://localhost:8080/metrics/invalid | promtool check metrics
```

### Pseudo business methods

```
curl http://localhost:8080/business/readfile
```

This business function as two related metrics:

```
sampleapp_readfile_bytes_total
sampleapp_readfile_total
```