{
    type: "Server"
    action: "create"
    system: {
        type: "Beaker"
        properties: {
            version: number | *24
            version: >=24
        }
    }
    properties: {
      name: string | *null
      family: string | *null
      method: string | *null
      arch: "x86_64" | "x86" | *null
    }
}