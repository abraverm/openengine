{
    type: "Volume"
    action: "create"
    system: {
        type: "Openstack"
    }
    properties: {
        size: string | *null
    }
    provisioner: "openstack volume create --size {{ size }}"
    response: {
        id: ""
    }
}