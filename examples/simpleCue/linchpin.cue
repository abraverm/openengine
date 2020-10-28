package engine

#Linchpin: _linchpin_openstack_server

_linchpin_openstack_server: {
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
  provisioner: "linchpin_server.tmpl"
  ...
}
