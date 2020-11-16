package example

R2 = {
  type: "Keypair"
  name: *((_solution & {xaction: _action, xresource: R3, xsystem: _systems[0]}).out & [#solution, ...#solution]) |  null
}

R1 = {
  name:     "my_server"
  type:     "Server"
  imageRef: "Fedora-31"
  key_name: *((_solution & {xaction: _action, xresource: R2, xsystem: _systems[0]}).out & [#solution, ...#solution]) |  null
  networks: "ccit-net"
  flavorRef: "Sfd"
  _memory: "4g"
  _disk: "10g"
}

R3 = {
  type: "Keypair"
  name: "ccit"
}

R4 = {
  name:     "my_server 2"
  type:     "Server"
  imageRef: "Fedora-31"
  key_name: *((_solution & {xaction: _action, xresource: R2, xsystem: _systems[0]}).out & [#solution, ...#solution]) |  null
  networks: "ccit-net"
  flavorRef: "Sfd"
  _memory: "4g"
  _disk: "10g"
}

_resources: [R1, R4]
_systems: [
	{
		type: "openstack"
		nova: 1.3
	},
]

#provisioners: #Linchpin
#providers: #Openstack
