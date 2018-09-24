changedBranch=$(git symbolic-ref HEAD | sed -e 's,.*/\(.*\),\1,')

if [ $changedBranch == "dev" ]; then
      echo "Not allowed to push in dev"
      exit 1
fi
