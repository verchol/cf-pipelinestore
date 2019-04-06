export REPO_BRANCH=new-pipeline
if [ ! -d "$REPO_NAME" ]; then
           echo cloning ...
           git clone https://$GIT_TOKEN@github.com/$REPO_OWNER/$REPO_NAME.git
fi

cd ./$REPO_NAME
git checkout $REPO_BRANCH 
cd ..

