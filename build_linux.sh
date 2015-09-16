go install netfiles/pcshow
mkdir -p $GOPATH/bin/netfiles/
cp -r $GOPATH/src/netfiles/static $GOPATH/bin/netfiles/
mv $GOPATH/bin/pcshow $GOPATH/bin/netfiles
