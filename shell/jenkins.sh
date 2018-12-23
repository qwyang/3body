function install_java()
{
    tar -xzf jdk-8u65-linux-x64.gz
    sed -i '/export JAVA_HOME/d' /etc/profile
    sed -i '/export PATH=${JAVA_HOME}/d' /etc/profile
    sed -i '/export CLASSPATH/d' /etc/profile
    echo 'export JAVA_HOME=/root/jenkins/jdk1.8.0_65' >> /etc/profile
    echo 'export PATH=${JAVA_HOME}/bin:${PATH}' >> /etc/profile
    echo 'export CLASSPATH=.:${JAVA_HOME}/lib/dt.jar:${JAVA_HOME}/lib/tools.jar' >> /etc/profile
    source /etc/profile
    java -version || exit 1
    echo "=== please run source /etc/profile for first time use or reopen a new shell section. ==="
}

function run_jenkins(){
    pidfile=/var/run/jenkins.pid
    logfile=/tmp/jenkins.log
    if [ -f ${pidfile} ];then
        kill -9 `cat ${pidfile}`
    fi
    java -jar jenkins.war &> ${logfile} &
    pid=$!
    echo "$pid" > ${pidfile}
    echo "jenkins running, pidfile:${pidfile}, log file: ${logfile}"
}

install_java
run_jenkins