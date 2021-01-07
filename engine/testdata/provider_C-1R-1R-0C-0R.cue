{
    type: "Server"
    action: "create"
    name: "C-1R-1R"
    system: {
        type: "Openstack"
    }
    constrains: [
        {
            name: "1R-1R"
            pre: [{
                action: "create"
                resource: {
                    type: "Storage"
                }
            }]
            post: [{
                action: "create"
                resource: {
                    type: "Storage"
                }
            }]
        }
    ]
}