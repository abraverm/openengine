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
    constrains: [{
        name: "0S-0S-1nmC-0R"
    }]
    system: {
        type: "Openstack"
        name: "Minimal"
    }
}]]