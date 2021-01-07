{
    type: "Server"
    action: "create"
    system: {
        type: "Openstack"
    }
    properties: {
        name: string | *null
        _random: bool | *null
    }
    implicit: {
        name: [{script: "something", args: { a: properties._random }}]
    }
}