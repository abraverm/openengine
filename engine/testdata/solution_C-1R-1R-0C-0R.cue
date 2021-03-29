[[{
    name: "R(Minimal)S(Minimal)PD(C-1R-1R)PR(Minimal)"
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "Minimal"
    }
    constrains: [{
        name: "1R-1R"
        pre: [{
            resolved: true
            action:   "create"
            resource: {
                type: "Storage"
            }
            solutions: [{
                name:     "S(Minimal)PD(ID-Storage)PR(ID-Storage)"
                resolved: true
                resource: {
                    type: "Storage"
                }
                system: {
                    type: "Openstack"
                    name: "Minimal"
                }
                provisioner: "example.sh"
            }]
        }]
        post: [{
            resolved: true
            action:   "create"
            resource: {
                type: "Storage"
            }
            solutions: [{
                name:     "S(Minimal)PD(ID-Storage)PR(ID-Storage)"
                resolved: true
                resource: {
                    type: "Storage"
                }
                system: {
                    type: "Openstack"
                    name: "Minimal"
                }
                provisioner: "example.sh"
            }]
        }]
    }]
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}]]