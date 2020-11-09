package engine

#Linchpin: #_linchpin_openstack_server

#_linchpin_openstack_server: {
  action: "create"
  system: {
    type: "openstack"
    nova: >1.0
  }
  resource: {
    type: "Server"
    name: string
    imageRef: string
    flavorRef: string | [#implicit_task, ...#implicit_task] | null
    key_name: string
    networks: string
  }
  provisioner: "linchpin_server.tmpl"
} 
