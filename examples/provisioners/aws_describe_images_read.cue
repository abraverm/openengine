{
    type: "DescribeImages"
    action: "read"
    system: {
        type: "AWS"
    }
    properties: {
        name: string | *null
    }
    provisioner: "curl https://ec2.amazonaws.com/?Action=DescribeImages&Filter.1.Name=name&Filter.1.Value.1={{ name }}"
}