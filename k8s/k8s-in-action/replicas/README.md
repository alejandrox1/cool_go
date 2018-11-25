# Managing pods

## Liveness probe
### Command
```
livenessProbe:
    exec:
        command:
            - cat
            - /tmp/healthy
    initialDelaySeconds: 5
    periodSeconds: 5
```

### HTTP
```
livenessProbe:
    httpGet:
        path: /health
        port: 8080
        httpHeaders:
        - name: X-Custom-Header
          value: Awesome
    initialDelaySeconds: 5
    timeoutSeconds: 1
    periodSeconds: 10
    failureThreshold: 3
```

To make things better, you can use a named port:
```
ports:
    - name: liveness-port
    containerPort: 8080
    hostPort: 8080

livenessProbe:
    httpGet:
        path: /health
        port: liveness-port
```

For more, see [K8s: configure liveness and readiness probes](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-probes/).

## ReplicaSets and labels
```
selector:
    matchExpressions:
        - key: app
        operator: In
        values:
            - kubia
```

* `In` - specify `values`
* `NotIn` - specify `values`
* `Exists` - do not specify `values`. Checks if key exists.
* `DoesNotExist` - do not specify `values`. Checks if key does not exist.
