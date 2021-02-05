# An identity and a service create passport records for the identity

SCRIPTPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

NODECRYPT=$SCRIPTPATH/../../../node/crypt.js

# [SETUP TEST] Create service and user identities

SERVICE=$(. $SCRIPTPATH/../invite/test_all.sh)
IDENTITY=$(. $SCRIPTPATH/../invite/test_all.sh)

SERV_ADDR=$(echo $SERVICE | jq -r '.address')
SERV_ID=$(echo $SERVICE | jq -r '.id')

IDEN_ADDR=$(echo $IDENTITY | jq -r '.address')
IDEN_ID=$(echo $IDENTITY | jq -r '.id')

# [TEST]

# Identity creates a private changable record for itself 

# TODO Encrypt value, record with the same key shouldn't be created
RECORD_REQ=$(mbcorecrd tx crsign create-record \
 some-passport-record \
 some-ecnrypted-record-value \
 IDENTITY_RECORD \
 PRIVATE \
 0 \
 --from $(mbcorecrd keys show $IDEN_ADDR -a) -y)

RECORD_ID=$(echo $RECORD_REQ | jq -r '.logs[0].events[0].attributes[0].value')

echo "Private mutable record ID:"$RECORD_ID