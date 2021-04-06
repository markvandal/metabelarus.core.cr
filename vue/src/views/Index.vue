<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form path="metabelarus.mbcorecr.crconsent" type="request" :fields="[ ['creator', 1, 'string'] , ['initiator', 2, 'string'] , ['recipient', 3, 'string'] , ['requestType', 4, 'string'] , ['status', 5, 'string'] , ['value', 6, 'int32'] , ['memo', 7, 'string'] , ['promoUrl', 8, 'string'] , ['creationDt', 9, 'string'] , ['finalDt', 10, 'string'] ]" />
      <sp-type-form path="metabelarus.mbcorecr.crsign" type="auth" :fields="[ ['creator', 1, 'string'] , ['identity', 2, 'string'] , ['service', 3, 'string'] , ['key', 4, 'string'] , ['status', 5, 'string'] , ['creationDt', 6, 'string'] , ['availabilityDt', 7, 'string'] ]" />
      <sp-type-form path="metabelarus.mbcorecr.crsign" type="id2sign" :fields="[ ['creator', 1, 'string'] , ['identity', 2, 'string'] , ['signature', 3, 'string'] ]" />
      <sp-type-form path="metabelarus.mbcorecr.crsign" type="signature" :fields="[ ['creator', 1, 'string'] , ['identity', 2, 'string'] , ['service', 3, 'string'] , ['key', 4, 'string'] , ['secret', 5, 'string'] , ['creationDt', 6, 'string'] , ['availabilityDt', 7, 'string'] ]" />
      <sp-type-form path="metabelarus.mbcorecr.mbcorecr" type="invite" :fields="[ ['creator', 1, 'string'] , ['inviter', 2, 'string'] , ['invitee', 3, 'string'] , ['level', 4, 'string'] , ['key', 5, 'string'] , ['creationDt', 6, 'string'] ]" />
      <sp-type-form path="metabelarus.mbcorecr.mbcorecr" type="identity" :fields="[ ['creator', 1, 'string'] , ['AccountID', 2, 'string'] , ['IdentityType', 3, 'string'] , ['Details', 4, 'string'] , ['CreationDt', 5, 'string'] ]" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
