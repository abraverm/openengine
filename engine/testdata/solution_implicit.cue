[[{
    name: "S(Minimal)"
    properties: {
        name: [{
            resolved: true
            script:   "something"
            args: {
                a: true
            }
        }]
    }
    provisioner: "example.sh"
    resource: {
        type: "Server"
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}]]