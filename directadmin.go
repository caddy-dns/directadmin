package directadmin

import (
	"strconv"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/directadmin"
)

// Provider lets Caddy read and manipulate DNS records hosted by this DNS provider.
type Provider struct{ *directadmin.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.directadmin",
		New: func() caddy.Module { return &Provider{new(directadmin.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	p.Provider.Logger = ctx.Logger()
	p.Provider.ServerURL = caddy.NewReplacer().ReplaceAll(p.Provider.ServerURL, "")
	p.Provider.User = caddy.NewReplacer().ReplaceAll(p.Provider.User, "")
	p.Provider.LoginKey = caddy.NewReplacer().ReplaceAll(p.Provider.LoginKey, "")

	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	directadmin {
//	    host <host_url>
//	    user <user>
//	    login_key <login_key>
//		insecure_requests <true/false>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "host":
				if p.Provider.ServerURL != "" {
					return d.Err("host already set")
				}
				if d.NextArg() {
					p.Provider.ServerURL = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "user":
				if p.Provider.User != "" {
					return d.Err("user already set")
				}
				if d.NextArg() {
					p.Provider.User = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "login_key":
				if p.Provider.LoginKey != "" {
					return d.Err("login key already set")
				}
				if d.NextArg() {
					p.Provider.LoginKey = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "insecure_requests":
				if d.NextArg() {
					val, err := strconv.ParseBool(d.Val())
					if err != nil {
						return d.Err("login key already set")
					}
					p.Provider.InsecureRequests = val
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.ServerURL == "" {
		return d.Err("missing host")
	}
	if p.Provider.User == "" {
		return d.Err("missing user")
	}
	if p.Provider.LoginKey == "" {
		return d.Err("missing login key")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
