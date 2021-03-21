[[{
    name: "R(ED-Other2)S(Other)PD(Other)PR(Other)"
    provisioner: "example.sh"
    resource: {
        type: "Other"
        name: "ED-Other2"
        solutions: [{
            name: "R(Other)S(Other)PD(Other)PR(Other)"
            resource: {
                type: "Other"
                name: "Other"
            }
            system: {
                type: "Other"
                name: "Other"
            }
            provisioner: "example.sh"
            resolved: true
        }]
        dependencies: ["Other"]
    }
    system: {
        type: "Other"
        name: "Other"
    }
}, {
    name: "R(Other)S(Other)PD(Other)PR(Other)"
    provisioner: "example.sh"
    resource: {
        type: "Other"
        name: "Other"
    }
    system: {
        type: "Other"
        name: "Other"
    }
}], [{
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