{
    type: "Flavor"
    action: "read"
    system: {
        type: "Openstack"
    }
    provisioner: "openstack flavor get --minram {{ minRam }}"
    properties: {
      minRam?: string | *null
    }
}