package jenkinsci

import (
	// "io/ioutil"
	// "crypto/tls"
	// "crypto/x509"
	// "net/http"
	jenkins "github.com/DanielMabbett/gojenkins"
)

// Config is the set of parameters needed to configure the JenkinsCI provider.
type Config struct {
	jenkinsEndpoint      string
	jenkinsAdminUsername string
	jenkinsAdminPassword string
	insecure             bool
}

// Client Config
func (c *Config) Client() (*jenkins.Jenkins, error) {
	client := jenkins.CreateJenkins(nil, c.jenkinsEndpoint, c.jenkinsAdminUsername, c.jenkinsAdminPassword)

	_, err := client.Init()
	if err != nil {
		return nil, err
	}

	return client, nil
}
