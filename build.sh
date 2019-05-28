go get -u github.com/hashicorp/terraform/helper/schema
go get -u github.com/danielmabbett/terraform-provider-jenkinsci/jenkins
# go build -o terraform-provider-jenkinsci
go build -o terraform-provider-example

# mkdir -p ~/.terraform.d/plugins
# mkdir -p ~/.terraform.d/plugins/darwin_amd64
# 
# mv -f ./terraform-provider-jenkinsci ~/.terraform.d/plugins/darwin_amd64/example
# 
# ls ~/.terraform.d/plugins/darwin_amd64/