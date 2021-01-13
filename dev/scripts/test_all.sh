
mbcorecrd keys export user1 2>&1 | tee ./tmppk

PKEY=$(sed -n 2,10p ./tmppk)

PKEYD=$(mbcorecrd query mbcorecr unpack-privkey -- "$PKEY" 11112222)

INVITE=$(mbcorecrd tx mbcorecr create-invite LevelSuper SERVICE \
 --from $(mbcorecrd keys show user1 -a))

PAYLOAD=$(echo $INVITE | jq -r '.logs[0].events[0].attributes[1].value')

PAYLOADD=$(mbcorecrd query mbcorecr decrypt-payload $PKEYD $PAYLOAD)

PAYLOADD=$(echo "$PAYLOADD" | sed 's/\n/\\n/g')

INVITEID=$(echo $PAYLOADD | jq -r '.inviteID')

SEQUENCE=$(echo $PAYLOADD | jq -r '.mnemonic')

ACCEPTEDACC=$(mbcorecrd tx mbcorecr accept-invite $INVITEID $SEQUENCE)

echo "$ACCEPTEDACC" > ./tmpinvited

PKEY=$(sed -n 26p ./tmpinvited)

echo $PKEY > ./tmppk2

PKEY=$(<./tmppk2)

ACCOUNT=$(echo $(sed -n 14p ./tmpinvited) | sed 's/.*://')

PKEYD=$(mbcorecrd query mbcorecr unpack-privkey -- "$PKEY" 11112222)

mbcorecrd query mbcorecr decrypt-payload $PKEYD $ACCOUNT
