### Secret Key Example

The example creates secret key. The key name parameter in variables.tf can let you set secret key name.

### Get up and running

* Set up access key and secret:

        export NETEASE_ACCESS_KEY="***"
        export NETEASE_ACCESS_SECRET="***"

* Planning phase

		TF_LOG=TRACE TF_LOG_PATH=tf-`date +%Y-%m-%d-%H-%M-%S`.log terraform plan

* Apply phase

		TF_LOG=TRACE TF_LOG_PATH=tf-`date +%Y-%m-%d-%H-%M-%S`.log terraform apply

* Destroy

		TF_LOG=TRACE TF_LOG_PATH=tf-`date +%Y-%m-%d-%H-%M-%S`.log terraform destroy