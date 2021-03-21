[[{
    name: "S(Minimal)PD(C-empty)PR(Minimal)"
    properties: {
        name: "explicit"
    }
    provisioner: "example.sh"
    resource: {
        type: "Server"
        properties: {
            name: "explicit"
        }
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}]]