// Copyright (c) 2016 The Go Authors. All rights reserved.
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

// nolint //forked code
package autocert

import (
	"context"
	"crypto"
	"sync"
	"time"
)

// renewJitter is the maximum deviation from Manager.RenewBefore.
const renewJitter = time.Hour

// domainRenewal tracks the state used by the periodic timers
// renewing a single domain's cert.
type domainRenewal struct {
	m   *Manager
	ck  certKey
	key crypto.Signer

	timerMu sync.Mutex
	timer   *time.Timer
}

// start starts a cert renewal timer at the time
// defined by the certificate expiration time exp.
//
// If the timer is already started, calling start is a noop.
func (dr *domainRenewal) start(exp time.Time) { _ = "STUB: not implemented"; return }

// stop stops the cert renewal timer.
// If the timer is already stopped, calling stop is a noop.
//
//nolint:unused
func (dr *domainRenewal) stop() { _ = "STUB: not implemented"; return }

// renew is called periodically by a timer.
// The first renew call is kicked off by dr.start.
func (dr *domainRenewal) renew() { _ = "STUB: not implemented"; return }

// TODO: rotate dr.key at some point?

// updateState locks and replaces the relevant Manager.state item with the given
// state. It additionally updates dr.key with the given state's key.
func (dr *domainRenewal) updateState(state *certState) { _ = "STUB: not implemented"; return }

// do is similar to Manager.createCert, but it doesn't lock a Manager.state item.
// Instead, it requests a new certificate independently and, upon success,
// replaces dr.m.state item with a new one and updates cache for the given domain.
//
// It may lock and update the Manager.state if the expiration date of the currently
// cached cert is far enough in the future.
//
// The returned value is a time interval after which the renewal should occur again.
func (dr *domainRenewal) do(ctx context.Context) (time.Duration, error) {
	_ = "STUB: not implemented"
	// a race is likely unavoidable in a distributed environment
	// but we try nonetheless
	return *new(time.Duration), nil
}

func (dr *domainRenewal) next(expiry time.Time) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// add a bit of randomness to renew deadline

var testDidRenewLoop = func(next time.Duration, err error) {}
