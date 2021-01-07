{
    name: "Mixed"
    type: "Server"
    action: "create"
    system: {
        type: "Openstack"
    }
    properties: {
        name: string | *null
        _random: bool | *null
        memory?: string | *null
        disk: string | *null
    }
    implicit: {
        name: [{script: "something", args: { a: properties._random }, resolved: len([ for a in args if a == null {null}]) == 0}]
    }
}