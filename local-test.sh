#go get -u github.com/DanielMabbett/gojenkins
##go get -u github.com/bndr/gojenkins
#go get -u github.com/DanielMabbett/terraform-provider-jenkinsci/jenkinsci
#go get -u github.com/hashicorp/terraform/helper/schema
#go get -u github.com/danielmabbett/terraform-provider-jenkinsci/jenkins

go build -o terraform-provider-jenkinsci
mv terraform-provider-jenkinsci ./test


cd ./test
terraform init 
terraform plan 
cd ..