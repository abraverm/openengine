{
    type: "Server"
    action: "create"
    system: {
        type: "Openstack"
        properties: {
            nova: number | *2.19
            nova: >=2.19
            nova: <2.37
        }
    }
    provisioner: "openstack server create --image {{ imageRef }} --flavor  {{ flavorRef }} {{ name }}"
    properties: {
      name: string | *null
      flavorRef: string | *null
      imageRef: string | *null
    }
}