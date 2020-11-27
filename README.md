# Prometheus training application

This training app provides to features:

* A sample app with a target which can be registered in Prometheus
* A webhook app which is used together with Alertmanager and acts as a webhook receiver

## Start the sample app
Download a release from https://github.com/rotscher/prometheus-training-app/releases and unpack the archive.
Then start the sample app as follows:

```
./prometheus-training-app sampleapp [--port 8080]
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
curl http://localhost:8080/business/random
```

Application specific metrics are prefixed with `sampleapp_`.

## Start the webhook app

```
./prometheus-training-app webhook
```

This app runs on port `5001` and outputs requests to standard error.
