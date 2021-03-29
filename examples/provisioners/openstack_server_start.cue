{
    type: "Server"
    action: "update"
    system: {
        type: "Openstack"
    }
    provisioner: "openstack server start {{ server_id }}"
    properties: {
      name: string | *null
      start: true | *null
    }
}