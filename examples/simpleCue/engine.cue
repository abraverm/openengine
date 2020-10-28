package engine

#action: *"create" | "delete" | "read" | "update"

_action: #action @tag(action)

_resources: [
  {
    name: "my_server"
    imageRef: "Fedora-31"
    key_name: "ccit"
    networks: "ccit-net"
  }
]

_systems: [
  {
    type: "openstack"
    nova: 1.3
  }
]

#Server: _openstack_server

_openstack_server: {
  action: "create"
  system: {
    type: "openstack"
    nova: >1.0
  }
  resource: {
    name: string
    imageRef: string
    flavorRef?: string
    key_name: string
    networks: string
  }
  ...
}


  
#provisioner: #Linchpin
#provider: #Server

solutions: [{
  for _resource in _resources {
    for _system in _systems {
      solution: #provisioner
      solution: #provider
      solution: { 
        resource: _resource
        system: _system
      }
      solution: { action: _action }
    }
  }
}]


