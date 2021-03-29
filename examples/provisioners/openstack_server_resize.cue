{
    type: "Server"
    action: "update"
    system: {
        type: "Openstack"
    }
    provisioner: "openstack server resize --flavor {{ flavorRef }} {{ name }}"
    properties: {
        name: string | *null
        flavorRef: string | *null
    }
}