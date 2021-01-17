# Handshake between Wallet\Client of some user, server side of the service,
# and the Citizen Register

SCRIPTPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# [SETUP TEST] Create service and user identities

SERVICE=$(. $SCRIPTPATH/../invite/test_all.sh)
IDENTITY=$(. $SCRIPTPATH/../invite/test_all.sh)

SERV_ADDR=$(echo $SERVICE | jq -r '.address')
SERV_ID=$(echo $SERVICE | jq -r '.id')

IDEN_ADDR=$(echo $IDENTITY | jq -r '.address')
IDEN_ID=$(echo $IDENTITY | jq -r '.id')

# [TEST]

# Service should encrypt some token with user's public key to make sure
# that the client represents the user.
# OPTIONAL FLOW: Key can contain any metainfo, while user should sign
# his requests with his Private Key, while the signature can be always checked
# with the PubKey.
# This version of the flow is more secure, but it requires more computational
# power.

ENCRYPTED_KEY=$(mbcorecrd query mbcorecr encrypt some-uid-generated-by-service \
 $(mbcorecrd keys show $IDEN_ADDR -p))

echo "Ecnrypted key: "$ENCRYPTED_KEY

# The service requests an authentication handshake with the user

AUTH_REQ=$(mbcorecrd tx crsign request-auth $SERV_ID $IDEN_ID $ENCRYPTED_KEY \
 --from $(mbcorecrd keys show $SERV_ADDR -a) -y)

AUTH_ID=$(echo $AUTH_REQ | jq -r '.logs[0].events[0].attributes[0].value')

echo "Auth ID: "$AUTH_ID

# The user can sign (confirm) the requested authentication.
# At this moment the recommended experation time is set.
# PURPOSE 1: This experation time will define the time
# when the service has access to other user related actions inside the
# Citizen Register.
# PURPOSE 2: To check if user's client still has access to service functions
# from the name of this user.
# In different cases it can be required to check if the service could
# perform specific actions with user data in specific period of time.
AUTH_CONFIRM=$(mbcorecrd tx crsign confirm-auth $IDEN_ID $SERV_ID \
--from $(mbcorecrd keys show $IDEN_ADDR -a) -y)

AUTH_RESULT=$(echo $AUTH_CONFIRM | jq -r '.logs[0].events[0].type')

echo "Auth confirmation result: "$AUTH_RESULT

# As a part of handshake the user should send to the service the 
# unencrypted token, to prove that this is really the client who
# communicates with the service may represent the user.
AUTH_RECORD=$(mbcorecrd query crsign show-auth $AUTH_ID --output json)

KEYTODECRYPT=$(echo $AUTH_RECORD | jq -r '.Auth.key')

DECRYPTEDKEY=$(mbcorecrd tx mbcorecr decrypt $KEYTODECRYPT \
 --from $(mbcorecrd keys show $IDEN_ADDR -a) -y)

echo "Decrypted Key: "$DECRYPTEDKEY
