
# Process of inviting one user or service by another user of the system.

# The inviter should create the invite record. As a result he or she:
# 1. Creates temporary account (mnemonics, address, pk, pubkey);
# 2. Creates invite that is linked to this temporary account;
# 3. Gets invite id.
INVITE=$(mbcorecrd tx mbcorecr create-invite LevelSuper SERVICE \
 --from $(mbcorecrd keys show user1 -a) -y)

TMPADDR=$(echo $INVITE | jq -r '.logs[0].events[0].attributes[1].value')
INVITEID=$(echo $INVITE | jq -r '.logs[0].events[0].attributes[2].value')

# The invitee gets from the inviter the tmp account credentials (as a set 
# of mnemonics). 
# Then the invitee:
# 1. Creates a permanent account for him- or herself;
# 2. Sends the accept request from the temporary account with the permanent
# account address and the invite id;
# 3. Receives his or her Identity Id;
NEWACC=$(mbcorecrd tx mbcorecr accept-invite $INVITEID \
 --from $(mbcorecrd keys show $TMPADDR -a) -y)

ADDR=$(echo $NEWACC | jq -r '.logs[0].events[0].attributes[1].value')
ID=$(echo $NEWACC | jq -r '.logs[0].events[0].attributes[2].value')

echo "{\"address\": \"$ADDR\", \"id\": \"$ID\"}"