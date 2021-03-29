{
    name: "Generic Implicit 1"
    type: "Example Resource Type"
    action: "create"
    system: {
        type: "Example Provider"
    }
    properties: {
        name: string | *null
        _name: string | *null
    }
    implicit: {
        name: [{script: "to_upperscore", args: { a: properties._name }}]
    }
}