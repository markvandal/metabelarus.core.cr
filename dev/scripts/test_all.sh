
INVITE=$(mbcorecrd tx mbcorecr create-invite LevelSuper SERVICE \
 --from $(mbcorecrd keys show user1 -a) -y)

TMPADDR=$(echo $INVITE | jq -r '.logs[0].events[0].attributes[1].value')
INVITEID=$(echo $INVITE | jq -r '.logs[0].events[0].attributes[2].value')

NEWACC=$(mbcorecrd tx mbcorecr accept-invite $INVITEID \
 --from $(mbcorecrd keys show $TMPADDR -a) -y)

echo $NEWACC