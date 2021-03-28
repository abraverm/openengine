{
    type: "Image"
    action: "read"
    system: {
        type: "Openstack"
    }
    provisioner: "openstack image get --name {{ name }}"
    properties: {
      name?: string | *null
    }
}