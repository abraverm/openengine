{
    type: "Distro"
    action: "read"
    system: {
        type: "Beaker"
    }
    properties: {
        name: string | *null
    }
    provisioner: "bkr distros-list --name {{ name }}"
}