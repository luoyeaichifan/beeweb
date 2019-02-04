git pull
ps aufx | grep beeweb | grep -v grep | awk -F " " '{print $2}' | xargs -I {} echo {}
ps aufx | grep beeweb | grep -v grep | awk -F " " '{print $2}' | xargs -I {} kill -9 {}
cd $GOPATH/src/github.com/luoyeaichifan/beeweb
go build
 nohup ./beeweb >> proxy.log 2>&1 &
