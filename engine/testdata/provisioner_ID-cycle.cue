{
    type: "Server"
    action: "create"
    name: "ID-cycle"
    system: {
        type: "Openstack"
    }
    provisioner: "example.sh"
    properties:chicken: string | *null
    response: {
        chicken: string | *null
    }
}