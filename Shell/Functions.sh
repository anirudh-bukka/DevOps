## Functions

# Without arguments and without return
function just_display() {
    echo "Just display"
}

# With arguments and without return
function display_given_name() {
    echo "Hello, $1!"
}

# Without arguments and with return
function timeFunc() {
    date +"%T"
}

# With arguments and with return
function calculate() {
    result=$(( ($1 + $2) * $3 ))
    echo $result
}

just_display

display_given_name "ABC"

time=timeFunc
echo "Time: $time"

calculate 1 2 3