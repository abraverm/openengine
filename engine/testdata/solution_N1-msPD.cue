[[{
    name: "R(Minimal)S(Minimal)PD(Minimal)PR(Minimal)"
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "Minimal"
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}], [{
    name: "R(Minimal)S(Minimal)PD(Implicit4)PR(Minimal)"
    properties: {
        name: [{
            resolved: true
            script:   "something"
            args: {
                a: "always"
            }
        }]
    }
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "Minimal"
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}]]