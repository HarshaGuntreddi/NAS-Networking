ubi / switching layout

- gateway:
  - wan
  - lan (vlan 20)
  - media (vlan 30)
  - mgmt (vlan 10)
  - guest (vlan 40)
- poe switch:
  - aps on vlan 30/40 trunk
  - cameras on vlan 30
- truenas box:
  - nic on mgmt + media vlans (tagged)
- firewall:
  - allow plex from wan only via https reverse proxy
  - block guest -> mgmt
  - allow lan/media -> truenas

documented here so recruiter sees network thinking
