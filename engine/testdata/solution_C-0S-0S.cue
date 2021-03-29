[[{
    name: "R(Minimal)S(Minimal)PD(C-0S-0S)PR(Minimal)"
    provisioner: "example.sh"
    resource: {
        type: "Server"
        name: "Minimal"
    }
    constrains: [{
        name: "0S-0S"
    }]
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}]]