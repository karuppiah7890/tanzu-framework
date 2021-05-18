## tanzu pinniped-auth login

Login using an OpenID Connect provider

```
tanzu pinniped-auth login [flags]
```

### Examples

```

    # pinniped-auth login using OpenID Connect provider
    tanzu pinniped-auth login  --issuer https://issuer.example.com --client-id tanzu-cli

    # pinniped-auth login using OpenID Connect provider with TCP port for local host listener (authorization code flow only)
    tanzu pinniped-auth login  --issuer https://issuer.example.com --client-id tanzu-cli --listen-port=48095
```

### Options

```
      --ca-bundle strings                     Path to TLS certificate authority bundle (PEM format, optional, can be repeated).
      --ca-bundle-data strings                Base64 endcoded TLS certificate authority bundle (base64 encoded PEM format, optional, can be repeated)
      --client-id string                      OpenID Connect client ID. (default "pinniped-cli")
      --concierge-authenticator-name string   Concierge authenticator name
      --concierge-authenticator-type string   Concierge authenticator type (e.g., 'webhook', 'jwt')
      --concierge-ca-bundle-data string       CA bundle to use when connecting to the concierge
      --concierge-endpoint string             API base for the Pinniped concierge endpoint
      --concierge-namespace string            Namespace in which the concierge was installed (default "pinniped-concierge")
      --enable-concierge                      Exchange the OIDC ID token with the Pinniped concierge during login
  -h, --help                                  help for login
      --issuer string                         OpenID Connect issuer URL.
      --listen-port uint16                    TCP port for localhost listener (authorization code flow only).
      --request-audience string               Request a token with an alternate audience using RFC8693 token exchange
      --scopes strings                        OIDC scopes to request during login. (default ["offline_access, openid, pinniped:request-audience"])
      --session-cache string                  Path to session cache file. (default "~/.config/tanzu/pinniped/sessions.yaml")
      --skip-browser                          Skip opening the browser (just print the URL).
```

### SEE ALSO

* [tanzu pinniped-auth](tanzu_pinniped-auth.md)	 - Pinniped authentication operations (usually not directly invoked)

###### Auto generated by spf13/cobra on 10-May-2021