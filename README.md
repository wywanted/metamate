# MetaMate

MetaMate is an open-source semantic service bus and provides you an api for everything.

This monorepo hosts
- `asg` abstract schema graph - an abstraction of MetaMate's data and communication layer
- `auth-svc` authenticates `ClientAccounts`, hashes password and verifies `Tokens`
- `gen` generated sdks
- `generic` generic representation of types, powers everything that needs to handle a lot of different entities
- `kubernetes-svc` discovery service for services running on kubernetes
- `mastodon-svc` gateway service for mastodon
- `metactl` MetaMate's cli
- `metamate` semantic service bus
- `sqlx-svc` storage service that maps mql to sql

## Installation

#### metamate

To spin up a metamate instance simply run `docker run metamatex/metamate:latest`. It's preconfigured with some sane defaults.

#### metactl

`metactl` is MetaMate's cli. It generates sdks and helps you explore the asg.

macOs `brew install metamatex/taps/metactl`

For all other platforms, please see our [releases](https://github.com/metamatex/metamate/releases)

## Development

MetaMate develops in multiple stages

`v0.x` MetaMate tries to provide an abstract layer for all network connected datastores, which can be databases, websites, apis etc. The challenge here is to cover all major use-cases. MetaMate needs to be able to handle different kinds of paginations, entity respresentations and authorization methods.