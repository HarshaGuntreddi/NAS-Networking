nas-networking

what this is

all-in-one sketch of a home / small-lab nas + network stack
kept flat, no pretty folder tree
reads like:
 "i designed, wired, and automated a poe-backed truenas/plex + web stack on real gear"

high level

- truenas core box
  - zfs pool, smb/nfs shares
  - jails/containers for plex + misc services
- ubiquiti gear
  - vlan split: mgmt / lan / media / guests
  - poe for ap + cameras
- router + reverse proxy
  - nginx / caddy front for public site + internal services
  - tls, sane ciphers
- full-stack site hosted on same platform
  - simple golang+react sample here
- infra-as-config
  - ansible to push basic config
  - docker compose to stand up app/plex/metrics
- observability
  - prometheus + node_exporter + grafana (optional stub)
- hardening
  - fail2ban + ufw sketch
  - backups noted

not production ready
no secrets committed
enough moving parts to look real on github
