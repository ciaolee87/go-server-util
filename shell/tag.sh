version=$1 
msg=$2
git tag -a v${version} -m ${msg}
git push origin v${version}
