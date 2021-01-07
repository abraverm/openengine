{
    type: "Server"
    action: "create"
    name: "C-1S-1S"
    system: {
        type: "Openstack"
    }
    constrains: [
        {
            name: "1S-1S"
            pre: [ { script: "pre.sh" } ]
            post: [ { script: "pre.sh" } ]
        }
    ]
}