Name: {{.serviceName}}.rpc
ListenOn: 0.0.0.0:8080
Etcd:
  Hosts:
  - 127.0.0.1:2379
  Key: {{.serviceName}}.rpc
Log:
  Encoding: plain
  TimeFormat: 2006-01-02 15:04:05
  Stat: false