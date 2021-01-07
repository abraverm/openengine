{
    type: "Server"
    action: "create"
    name: "C-empty"
    system: {
        type: "Openstack"
    }
    constrains: [
        {
            name: "0S-0S-1nmC-0R"
            pre: []
            post: []
            properties: {
                name: string | *null
            }
        }
    ]
    properties: {
        name: string | *null
    }
}