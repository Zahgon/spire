// Copyright (c) 2017 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// nolint // forked code
package autocert

import (
	"crypto/tls"
	"net"
)

// NewListener returns a net.Listener that listens on the standard TLS
// port (443) on all interfaces and returns *tls.Conn connections with
// LetsEncrypt certificates for the provided domain or domains.
//
// It enables one-line HTTPS servers:
//
//	log.Fatal(http.Serve(autocert.NewListener("example.com"), handler))
//
// NewListener is a convenience function for a common configuration.
// More complex or custom configurations can use the autocert.Manager
// type instead.
//
// Use of this function implies acceptance of the LetsEncrypt Terms of
// Service. If domains is not empty, the provided domains are passed
// to HostWhitelist. If domains is empty, the listener will do
// LetsEncrypt challenges for any requested domain, which is not
// recommended.
//
// Certificates are cached in a "golang-autocert" directory under an
// operating system-specific cache or temp directory. This may not
// be suitable for servers spanning multiple machines.
//
// The returned listener uses a *tls.Config that enables HTTP/2, and
// should only be used with servers that support HTTP/2.
//
// The returned Listener also enables TCP keep-alives on the accepted
// connections. The returned *tls.Conn are returned before their TLS
// handshake has completed.
func NewListener(domains ...string) net.Listener {
	_ = "STUB: not implemented"
	return *new(net.Listener)
}

// Listener listens on the standard TLS port (443) on all interfaces
// and returns a net.Listener returning *tls.Conn connections.
//
// The returned listener uses a *tls.Config that enables HTTP/2, and
// should only be used with servers that support HTTP/2.
//
// The returned Listener also enables TCP keep-alives on the accepted
// connections. The returned *tls.Conn are returned before their TLS
// handshake has completed.
//
// Unlike NewListener, it is the caller's responsibility to initialize
// the Manager m's Prompt, Cache, HostPolicy, and other desired options.
func (m *Manager) Listener() net.Listener { _ = "STUB: not implemented"; return *new(net.Listener) }

type listener struct {
	conf *tls.Config

	tcpListener  net.Listener
	tcpListenErr error
}

func (ln *listener) Accept() (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// Because Listener is a convenience function, help out with
// this too.  This is not possible for the caller to set once
// we return a *tcp.Conn wrapping an inaccessible net.Conn.
// If callers don't want this, they can do things the manual
// way and tweak as needed. But this is what net/http does
// itself, so copy that. If net/http changes, we can change
// here too.

func (ln *listener) Addr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// net.Listen failed. Return something non-nil in case callers
// call Addr before Accept:

func (ln *listener) Close() error { _ = "STUB: not implemented"; return nil }

func homeDir() string { _ = "STUB: not implemented"; return "" }

func cacheDir() string { _ = "STUB: not implemented"; return "" }

// Worst case:
