{
    name: "Generic Implicit 2"
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
        name: [{
            resource: {
                type: "Example Resource Type"
                properties: {
                    name: "something"
                }
            }
            action: "create"
        }]
    }
}