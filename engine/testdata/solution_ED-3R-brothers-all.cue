[[{
    name: "R(ED-Minimal2)S(Minimal)PD(Minimal)PR(Minimal)"
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "ED-Minimal2"
        solutions: [{
            name: "R(Minimal)S(Minimal)PD(Minimal)PR(Minimal)"
            resource: {
                type: "Server"
                name: "Minimal"
            }
            system: {
                type: "Openstack"
                name: "Minimal"
            }
            provisioner: "example.sh"
            resolved: true
        }, {
            name: "R(Two)S(Minimal)PD(Minimal)PR(Minimal)"
            resource: {
                type: "Server"
                name: "Two"
            }
            system: {
                type: "Openstack"
                name: "Minimal"
            }
            provisioner: "example.sh"
            resolved: true
        }]
        dependencies: ["Minimal", "Two"]
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}, {
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
}, {
    name: "R(Two)S(Minimal)PD(Minimal)PR(Minimal)"
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "Two"
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}]]