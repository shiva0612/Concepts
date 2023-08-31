cd simpleserver && go run server.go -host "localhost" -srvcert "../out/ localhost.crt" -srvkey "../out/localhost.key"
cd client && go run client.go -cacert "../out/ExampleCA.crt"


cd advserver && go run server.go  -host "localhost" -srvcert "../out/localhost.crt" -srvkey "../out/localhost.key" -cacert "../out/ExampleCA.crt" -port 443 -certopt 4
cd client && go run client.go -clientcert "../out/client.crt" -clientkey "../out/client.key" -cacert "../out/ExampleCA.crt"
volumeMount: 
    - mountPath: /opt/src/certs
        name: core-tool-cert
        readOnly: true
    - mountPath: /opt/src/certs
        name: core-tool-key
        readOnly: true
volumes: 
    - name: core-tool-cert
    secret:
        defaultMode: 420
        secretName: core-tool-cert
    - name: core-tool-key
    secret:
        defaultMode: 420
        secretName: core-tool-key