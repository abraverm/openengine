{
    type: "Server"
    action: "create"
    name: "Implicit4"
    system: {
        type: "Openstack"
    }
    properties: {
        name: string | *null
    }
    implicit: {
        name: *[{script: "something", args: { a: "always" }}] | null
    }
}