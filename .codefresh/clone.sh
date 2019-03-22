if [ ! -d "/codefresh/volume/cf-pipelinestore" ]; then
           git clone https://${{GIT_TOKEN}}@github.com/verchol/cf-pipelinestore.git
fi