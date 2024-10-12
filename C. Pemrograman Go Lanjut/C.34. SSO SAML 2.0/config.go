package main

import "fmt"

var (
	samlCertificatePath = "./server.crt"
	samlPrivateKeyPath  = "./server.key"
	samlIDPMetadata     = "https://samltest.id/saml/idp"

	webserverPort    = 9000
	webserverRootURL = fmt.Sprintf("http://localhost:%d", webserverPort)
)
