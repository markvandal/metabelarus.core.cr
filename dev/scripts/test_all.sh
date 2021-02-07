
ROOT_SCRIPTPATH=$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )

. $ROOT_SCRIPTPATH/invite/test_all.sh
. $ROOT_SCRIPTPATH/auth/test_all.sh
. $ROOT_SCRIPTPATH/record/test_all.sh