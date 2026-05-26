package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/spiffe/spire/pkg/common/log"
	"github.com/spiffe/spire/pkg/common/version"
)

var (
	versionFlag = flag.Bool("version", false, "print version")
	configFlag  = flag.String("config", "oidc-discovery-provider.conf", "configuration file")
	expandEnv   = flag.Bool("expandEnv", false, "expand environment variables in config file")
)

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Println(version.Version())
		os.Exit(0)
	}

	if args := flag.Args(); len(args) > 0 {
		fmt.Fprintf(os.Stderr, "Error: unexpected arguments: %v\n", args)
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err := run(*configFlag, *expandEnv); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

func run(configPath string, expandEnv bool) error { _ = "STUB: not implemented"; return nil }

func buildNetListener(ctx context.Context, config *Config, log *log.Logger) (listener net.Listener, err error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func newSource(log logrus.FieldLogger, config *Config) (JWKSSource, error) {
	_ = "STUB: not implemented"
	return *new(JWKSSource), nil
}

// This is defensive; LoadConfig should prevent this from happening.

func newListenerWithServingCert(ctx context.Context, log logrus.FieldLogger, config *Config) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func newACMEListener(log logrus.FieldLogger, config *Config) (net.Listener, error) {
	_ = "STUB: not implemented"
	return *new(net.Listener), nil
}

func logHandler(log logrus.FieldLogger, handler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

// This code was borrowed and modified from the
// golang.org/x/crypto/acme/autocert package. It wraps a normal TCP listener to
// set a reasonable keepalive on the TCP connection in the same vein as the
// net/http package.
type tlsListener struct {
	*net.TCPListener
	conf *tls.Config
}

func (ln *tlsListener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}
