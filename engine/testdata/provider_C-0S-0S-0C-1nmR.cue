{
    type: "Server"
    action: "create"
    name: "C-empty"
    system: {
        type: "Openstack"
    }
    constrains: [
        {
            name: "0S-0S-0C-1nmR"
            pre: []
            post: []
            response: {
                no: bool | *null
            }
        }
    ]
    response: {
        id: string | *null
    }
}