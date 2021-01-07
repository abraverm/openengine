{
    type: "Server"
    action: "create"
    name: "ID-Minimal2"
    system: {
        type: "Openstack"
    }
    provisioner: "example.sh"
    properties:network: string | *null
    properties:storage: string | *null
}