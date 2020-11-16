package example

#Linchpin: [
  for provider in _linchpin
  for resource in provider
  for action in resource
  for system_match in action
  { system_match }
]

_linchpin: openstack: keypair: create: [{
  action: "create"
  system: {
    type: "openstack"
    nova: >1.0
  }
  resource: {
    type: "Keypair"
    name: string | [#solution, ...#solution]
  }
  response: {
    name: string | *""
  }
  provisioner: "linchpin_keypair.tmpl"
}]

_linchpin: openstack: server: create: [{
  action: "create"
  system: {
    type: "openstack"
    nova: >1.0
  }
  resource: {
    type: "Server"
    name: string
    imageRef: string| [#solution, ...#solution]
    flavorRef: string | [#implicit_task, ...#implicit_task] | null
    key_name: string | [#solution, ...#solution]
    networks: string
  }
  provisioner: "linchpin_server.tmpl"
  response: {
    name: string | *""
  }
}]

_linchpin: openstack: flavor: get: [{
  action: "get"
  system: {
    type: "openstack"
    nova: >1.0
  }
  resource: {
    type:       "Flavor"
    minDisk: string
    minMemory: string
  }
  provisioner: "linchpin_flavor.tmpl"
}]
