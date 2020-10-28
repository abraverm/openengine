#action: *"create" | "delete" | "read" | "update"

_my_resource: {
  resource: {
    name: "my_server"
    imageRef: "Fedora-31"
    key_name: "ccit"
    networks: 8080
  }
}

_my_system: {
  system: {
    type: "openstack"
    nova: 1.3
  }
}

#Server: _openstack_server

_openstack_server: {
  system: {
    type: "openstack"
    nova: >1.0
  }
  resource: {
    name: string
    action: action
    imageRef: string
    flavorRef: string | *null
    key_name: string
    networks: { port: string } | { network: string }
  }
  ...
}

#Linchpin: _linchpin_openstack_server

_linchpin_openstack_server: {
  system: {
    type: "openstack"
    nova: >1.0
  }
  resource: {
    name: string
    action: action
    imageRef: string
    flavorRef: string | *null
    key_name: string
    networks: string
  }
  provisioner: "linchpin_server.tmpl"
  ...
}
  
#provisioner: #Linchpin
#provider: #Server

solution: #provisioner & #provider & _my_resource & _my_system