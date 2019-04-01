
    echo processing codefresh-io/plugins/submit.json 
    export REPO_OWNER=codefresh-io
    export REPO_NAME=plugins
 
    if [ ! -d "$PWD/$REPO_NAME" ]; then
           echo cloning ...
           git clone https://$GIT_TOKEN@github.com/$REPO_OWNER/$REPO_NAME.git
    fi
    cd  $PWD/$REPO_NAME
    echo $PWD 
    ls .
    mv submit.json submit.new.json
    git checkout HEAD~1 submit.json 
    mv submit.json submit.old.json
