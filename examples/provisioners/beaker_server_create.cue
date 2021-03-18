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
    provisioner: "generate_xml {{ name }} {{ family }} {{ method }} {{ arch }} && bkr job-submit --wait my-beaker-job.xml"
    properties: {
      name: string | *null
      family: string | *null
      method: string | *null
      arch: string | *null
    }
}