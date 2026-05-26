package main

type DomainPolicy = func(domain string) error

// DomainAllowlist returns a policy that allows any domain in the given domains
func DomainAllowlist(domains ...string) (DomainPolicy, error) {
	_ = "STUB: not implemented"
	return *new(DomainPolicy), nil
}

// AllowAnyDomain returns a policy that allows any domain
func AllowAnyDomain() DomainPolicy { _ = "STUB: not implemented"; return *new(DomainPolicy) }

func toDomainKey(domain string) (string, error) { _ = "STUB: not implemented"; return "", nil }
