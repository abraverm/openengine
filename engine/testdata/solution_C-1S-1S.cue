[[{
    name: "R(Minimal)S(Minimal)PD(C-1S-1S)PR(Minimal)"
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "Minimal"
    }
    constrains: [{
        name: "1S-1S"
        pre: [{
            resolved: true
            script:   "pre.sh"
        }]
        post: [{
            resolved: true
            script:   "pre.sh"
        }]
    }]
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}]]