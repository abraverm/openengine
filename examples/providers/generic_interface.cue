{
    type: "Example Resource Interface Type"
    action: "create"
    system: {
        type: "Example Provider"
    }
    properties: {
        example: string | *null
    }
    interfaces:example:[
        {
            name: "Resolved from Example Resource Response Type"
            type: "Example Resource Response Type"
            action: "create"
            response: {
                other: string | *null
            }
            field: "other"
        }
    ]
}