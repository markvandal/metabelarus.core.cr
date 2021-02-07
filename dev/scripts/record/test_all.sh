# An identity and a service create passport records for the identity

SCRIPTPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# [SETUP TEST] Create service and user identities

. $SCRIPTPATH/../auth/test_all.sh

# [TEST]

# Identity creates a private unchangable record for itself
# This kind of records required to depict personal identification
# data, like name, phisical passport number etc. 

# TODO Encrypt value
RECORD_REQ=$(mbcorecrd tx crsign create-record \
 some-passport-record$RANDOM some-ecnrypted-record-value \
 IDENTITY_RECORD PRIVATE 0 \
 --from $(mbcorecrd keys show $IDEN_ADDR -a) -y)

RECORD_ID=$(echo $RECORD_REQ | jq -r '.logs[0].events[0].attributes[0].value')

echo "Private record ID: "$RECORD_ID

SEAL_REQ=$(mbcorecrd tx crsign update-record \
 $RECORD_ID "" 0 REOCRD_UPDATE_SEAL \
 --from $(mbcorecrd keys show $IDEN_ADDR -a) -y)

SEAL_RES=$(echo $SEAL_REQ | jq -r '.logs[0].events[1].attributes[0].value')

echo "Sealing reasult: "$SEAL_RES

REOPEN_REQ=$(mbcorecrd tx crsign update-record \
 $RECORD_ID "" 0 REOCRD_UPDATE_REOPEN \
 --from $(mbcorecrd keys show $IDEN_ADDR -a) -y)

ERROR=$(echo $REOPEN_REQ | jq -r '.raw_log')

echo "Shouldn't be able to reopen the record: "$ERROR

# Only service, that an Idnetity autheticated in, should be able to 
# create public records, for example, to store KYC badges 

RECORD_REQ=$(mbcorecrd tx crsign create-record \
 service-passport-record$RANDOM some-ecnrypted-record-value \
 PROVIDER_RECORD PUBLIC 0 $IDEN_ID \
 --from $(mbcorecrd keys show $SERV_ADDR -a) -y)

RECORD_ID=$(echo $RECORD_REQ | jq -r '.logs[0].events[0].attributes[0].value')

echo "Provided record ID: "$RECORD_ID

SEAL_REQ=$(mbcorecrd tx crsign update-record \
 $RECORD_ID "" 0 REOCRD_UPDATE_SEAL \
 --from $(mbcorecrd keys show $SERV_ADDR -a) -y)

SEAL_RES=$(echo $SEAL_REQ | jq -r '.logs[0].events[1].attributes[0].value')

echo "Sealing reasult: "$SEAL_RES

REOPEN_REQ=$(mbcorecrd tx crsign update-record \
 $RECORD_ID "" 0 REOCRD_UPDATE_REOPEN \
 --from $(mbcorecrd keys show $SERV_ADDR -a) -y)

REOPEN_RES=$(echo $SEAL_REQ | jq -r '.logs[0].events[1].attributes[0].value')

echo "Reopen result: "$REOPEN_RES