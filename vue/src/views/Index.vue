<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form path="metabelarus.mbcorecr.crsign" type="signatureList" :fields="[ ['creator', 1, 'string'] , ['rootSignatureId', 2, 'string'] , ['lastSignatureId', 3, 'string'] , ['nextSignatureId', 4, 'string'] , ['metadata', 5, 'string'] ]" />
		<sp-type-form path="metabelarus.mbcorecr.crsign" type="id2auth" :fields="[ ['creator', 1, 'string'] , ['identity', 2, 'string'] , ['auth', 3, 'string'] ]" />
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
