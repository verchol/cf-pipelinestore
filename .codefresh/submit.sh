
    echo clone started
    if [ ! -d "/codefresh/volume/$REPO_NAME" ]; then
           git clone https://${{GIT_TOKEN}}@github.com/${{REPO_OWNER}}/${{REPO_NAME}}.git
    fi
    cd  /codefresh/volume/$REPO_NAME
    mv submit.json submit.new.json
    git checkout HEAD~1 submit.json 
    mv submit.json > submit.old.json