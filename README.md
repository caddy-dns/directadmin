DirectAdmin module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with DirectAdmin.

## Caddy module name

```
dns.providers.directadmin
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "directadmin",
				"host": "DIRECT_ADMIN_HOST",
				"user": "DIRECT_ADMIN_USER",
				"login_key": "DIRECT_ADMIN_LOGIN_KEY"
			}
		}
	}
}
```

or with the Caddyfile:

```
# globally
{
	acme_dns directadmin ...
}
```

```
# one site
tls {
	dns directadmin ...
}
```
