#!/bin/bash
USER=lsb; PASSWORD=123456; KEYFILE=my
URL=http://192.168.3.29:4800/
MYDOMAIN=fangpi; GWADDR=leither.cn
MYAPP=lapp

function readInput {
    # take 2 params, 1st is the variable name for user to input, 2nd is the description of the variable
    local t=$1
    echo Please input $2 "(default: ${!t})"
    local i
    read i
    if [ -z "$i" ]; then
        echo Use default $2: ${!t}
    else
        echo $2: $i
    fi
}
readInput USER "User Name"
readInput PASSWORD "Password"
readInput KEYFILE "key file name"
readInput MYAPP "local APP directory"
readInput URL "local URL"

./Leither lssl runscript -s "local auth=require('auth'); return auth.Register('$USER', '$PASSWORD');"
./Leither lssl runscript -s "local node=require('mimei'); return node.MMSetRight(request.sid, 'mmroot', '', 0x07276707);"
echo "User "$USER" created and authorized"

./Leither lssl genkey -o $KEYFILE.key
./Leither lssl genca -k $KEYFILE.key -m "name=$KEYFILE" -o $KEYFILE.ca
./Leither lssl gencert -k $KEYFILE.key -c $KEYFILE.ca -m "name=forapp" -o $KEYFILE.cert
./Leither lssl signppt -c $KEYFILE.cert -m "CertFor=Self" -o ${KEYFILE}login.ppt
echo "Credential files created"

./Leither lssl reqservice -c $KEYFILE.cert -m RequestService=mimei -n $URL
./Leither deploy uploadapp -p ${KEYFILE}login.ppt -i ./$MYAPP -n $URL
./Leither.exe deploy backup -a $MYAPP -p ${KEYFILE}login.ppt -n $URL
echo "APP uploaded to service node"

./Leither.exe deploy setdomain -d $MYDOMAIN.$GWADDR -n $URL -a $MYAPP -p mylogin.ppt -m gwaddr=$GWADDR
./Leither.exe deploy backup -a $MYAPP -p ${KEYFILE}login.ppt -n $URL
echo "APP published successfully"