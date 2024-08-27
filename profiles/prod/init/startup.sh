#!/usr/bin/env bash
ROOT=/myntra

CURRENT_BUILD="$ROOT"/accio

ENV="production"

ACCIO_LOGS="$CURRENT_BUILD/logs"
ACCIO_BIN="$CURRENT_BUILD/bin"

cd "$CURRENT_BUILD" || exit
mkdir -p "$ACCIO_LOGS"
mkdir -p "$ACCIO_BIN"
mkdir -p "$ROOT"/pid
mkdir -p "$ROOT"/switch/go/switchcache-2297.accio/

RUNNING_PID="$ROOT"/pid/accio.pid
echo "Checking for pid in $RUNNING_PID"
if [ -f $RUNNING_PID ]; then
    pid=$(cat $RUNNING_PID)
    echo "Last know pid is $pid"
    process=$(ps -ef | grep "$pid" | grep -v grep | wc -l)
    if [ "$process" -ne 0 ]; then
        echo "accio already running"
        exit
    fi
fi

# ENABLE_NEWRELIC=true
# NEWRELICHOST=pas-style0
# if [ "$HOSTNAME" = $NEWRELICHOST ]; then
#     echo "Enabling Newrelic"
#     ENABLE_NEWRELIC=true
# fi

>$RUNNING_PID
echo "installing & Starting accio"
mkdir -p "$CURRENT_BUILD"/bin

APP_ENV_PATH="$CURRENT_BUILD/bin/setenv.sh"
. "$APP_ENV_PATH" "$ENABLE_NEWRELIC"

echo "newrelic value is $ENABLE_NEWRELIC"

# add gogc
export GOGC=7000
export GOMAXPROCS=16
chmod +x "$CURRENT_BUILD"/webapps/accio

echo $ACCIO_LOGS

if [ -d "$ACCIO_LOGS" ]; then
  # Take action if $DIR exists. #
  echo "Directory Exists"
fi

nohup "$CURRENT_BUILD"/webapps/accio -logPath="$ACCIO_LOGS" -env="$ENV" -config="$CURRENT_BUILD"/profiles/k8prodCmsFetch/conf/config-prod.yml >"$ACCIO_LOGS"/init.log 2>&1 &

echo $! >>$RUNNING_PID
sleep 10
tail "$ACCIO_LOGS"/init.log
echo "accio started"
