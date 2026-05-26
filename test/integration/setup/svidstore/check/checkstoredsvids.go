package main

import (
	"context"
	"crypto/x509"
	"fmt"
	"os"
	"time"

	entryv1 "github.com/spiffe/spire-api-sdk/proto/spire/api/server/entry/v1"
	"github.com/spiffe/spire-api-sdk/proto/spire/api/types"
	svidstorev1 "github.com/spiffe/spire-plugin-sdk/proto/spire/plugin/agent/svidstore/v1"
	"github.com/spiffe/spire/test/integration/setup/itclient"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: checkstoredsvids storageFile")
		os.Exit(1)
	}
	storageFile := os.Args[1]

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	client := itclient.NewLocalServerClient()
	defer client.Release()

	entriesResp := getEntries(ctx, client)

	storedSVIDS := getSVIDsFromFile(storageFile)

	currentBundle := getCurrentBundle(ctx, client)

	assertStoredSVIDs(entriesResp, storedSVIDS, currentBundle)
}

func getCurrentBundle(ctx context.Context, client *itclient.LocalServerClient) []*x509.Certificate {
	_ = "STUB: not implemented"
	return nil
}

func getEntries(ctx context.Context, client *itclient.LocalServerClient) *entryv1.ListEntriesResponse {
	_ = "STUB: not implemented"
	return nil
}

func assertStoredSVIDs(entries *entryv1.ListEntriesResponse, svids map[string]*svidstorev1.X509SVID, currentBundle []*x509.Certificate) {
	_ = "STUB: not implemented"
	return
}

// decode ASN.1 DER bundle

// decode certChain

// decode private key

// check spiffe id

func getSVIDsFromFile(storageFile string) map[string]*svidstorev1.X509SVID {
	_ = "STUB: not implemented"
	return nil
}

func getSecretName(selectors []*types.Selector) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func assertNoError(err error, format string, v ...any) { _ = "STUB: not implemented"; return }

func assertEqual(expected, actual any, format string, v ...any) { _ = "STUB: not implemented"; return }

func assertEqualCerts(expected, actual []*x509.Certificate, format string, v ...any) {
	_ = "STUB: not implemented"
	return
}
