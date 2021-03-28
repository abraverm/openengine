{
    type: "Server"
    action: "update"
    system: {
        type: "Openstack"
    }
    properties: {
      name: string & != ""| *null
      stop: true | *null
    }
}