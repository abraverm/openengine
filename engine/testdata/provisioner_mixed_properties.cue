{
    name: "Mixed"
    type: "Server"
    action: "create"
    system: {
        type: "Openstack"
    }
    provisioner: "example.sh"
    properties: {
        name: string | *null
        memory?: string | *null
        disk: string | *null
    }
}