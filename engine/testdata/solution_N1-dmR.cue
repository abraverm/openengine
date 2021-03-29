[[{
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
}], [{
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