if [ ! -d "./cf-pipelinestore" ]; then
           echo cloning ...
           git clone https://$GIT_TOKEN@github.com/verchol/cf-pipelinestore.git
fi