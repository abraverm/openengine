{
    type: "Server"
    action: "update"
    system: {
        type: "Openstack"
    }
    provisioner: "openstack server stop {{ server_id }}"
    properties: {
      name: string | *null
      stop: true | *null
    }
}