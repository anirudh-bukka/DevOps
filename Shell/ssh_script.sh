$log_file="/tmp/script_logs.log"

function connect_to_server() {
    hostname=$1
    password=$2
    user=$3
    echo "connecting to host $user@$hostname with provided password..."
}

function run_helper_command() {
    sudo su -c "
    source /opt/assure1/.bashrc
    cd /opt/assure1/core/collection
    echo '$A1BASEDIR' >> $log_file
    
    "
}

function run_lnav_command() {

}

# read -p "Enter hostname: " hostname
# read -sp "Enter password: " password
# read -p "Enter username: " user

connect_to_server "$hostname" "$password" "$user"