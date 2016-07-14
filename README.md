[This excellent document](http://www.techradar.com/us/news/software/how-ssl-and-tls-works-1047412)
explains the mechanism beneath TLS.


This
[namecheap.com's page](https://www.namecheap.com/support/knowledgebase/article.aspx/9474/69/how-do-i-create-a-pem-file-from-the-certificates-i-received-from-you)
briefly described the PEM file format defined in X.509 -- It is base64
encoded block of data encapsulated between

```
-----BEGIN CERTIFICATE ----- 
-----END CERTIFICATE -----
```


## Pieces of Wise

### Generate RSA Key

According to [this Gist](https://gist.github.com/denji/12b3a568f092ab951456):

```
openssl genrsa -out server.key 2048
```

The output file is in PEM format:

```
-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: DES-EDE3-CBC,DB98A9512DD7CBCF

yKTM+eoxBvptGrkEixhljqHSuE+ucTh3VqYQsgO6+8Wbh1docbFUKzLKHrferJBH
...
-----END RSA PRIVATE KEY-----
```

Above command doesn't need passphrase.  By
[this Heroku tutorial](https://devcenter.heroku.com/articles/ssl-endpoint#acquire-ssl-certificate),
the following command need passphrase:

```
openssl genrsa -des3 -out server.key 2048
```

This can strip off the passphrase:

```
openssl rsa -in server.pass.key -out server.key
```



### Generate Certificate Signing Request (CSR)

By [this Heroku tutorial](https://devcenter.heroku.com/articles/ssl-certificate-self):

```
openssl req -nodes -new -key server.key -out server.csr
```

This requires input of identification information.



### Self-Signed certificate

By
[this Heroku tutorial](https://devcenter.heroku.com/articles/ssl-certificate-self),
we can generate a self-signed certificate using aforementioned key and CSR:

```
openssl x509 -req -sha256 -days 365 -in server.csr -signkey server.key -out server.crt
```

By [this gist](https://gist.github.com/denji/12b3a568f092ab951456),
without generating the CSR, we can generate the certificate as well:

```
openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
```

This .crt file and aformentioned .key file can be passed to
`http.ListenAndServeTLS` to run an HTTPS server.

### curl and TLS

The Mac OS X contains a 7.43 version curl.  It doesn't work well 
