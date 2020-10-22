#### This is interview test simple golang HTTPS Server/Client

##### Generate cert(if certificate does'not exists)
```shell
# generate private key
openssl genrsa -out server.key 2048

# generate self-signed public key(x509)
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650

```

##### How to run unit test
```shell script
cd $GOPATH/src/interview-go/server_test
go test .

#### Assume input ---> result  ###
[golang]            --> [flase]
[golang]            --> [true]
[golang,python]     --> [true,false]
[golang,python,php] --> [true,true, false]
[golang,python,php] --> [true,true, true]

#### Test report should be like blow
ok  	interview-go/server_test	5.062s

```


<br/>

#### Thank you review my code!
