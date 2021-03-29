{
    type: "Server"
    action: "create"
    system: {
        type: "Openstack"
    }
    provisioner: "openstack server create --volume {{ volume }}"
    properties: {
        volume: string | *null
    }
}