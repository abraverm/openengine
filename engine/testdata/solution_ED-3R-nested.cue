[[{
    name: "R(ED-Grampa)S(Minimal)PD(Minimal)PR(Minimal)"
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "ED-Grampa"
        solutions: [{
            name: "R(ED-Minimal)S(Minimal)PD(Minimal)PR(Minimal)"
            resource: {
                type: "Server"
                name: "ED-Minimal"
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
                }]
                dependencies: ["Minimal"]
            }
            system: {
                type: "Openstack"
                name: "Minimal"
            }
            provisioner: "example.sh"
            resolved: true
        }]
        dependencies: ["ED-Minimal"]
    }
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}, {
    name: "R(ED-Minimal)S(Minimal)PD(Minimal)PR(Minimal)"
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "ED-Minimal"
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
        }]
        dependencies: ["Minimal"]
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
}]]