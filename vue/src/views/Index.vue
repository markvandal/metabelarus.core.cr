<template>
  <div>
    <div class="sp-container">
      <sp-sign-in />
      <sp-bank-balances />
      <sp-token-send />
      <!-- this line is used by starport scaffolding # 4 -->
		<sp-type-form type="allowance" :fields="['passportId', 'allowanceType', 'resolution', ]" module="mbpasstrust" />
		<sp-type-form type="consent" :fields="['extserviceId', 'passportId', 'resolution', ]" module="mbgovperm" />
		<sp-type-form type="extservice" :fields="['alias', 'detailsUrl', ]" module="mbgovperm" />
    </div>
  </div>
</template>

<script>
import * as sp from "@tendermint/vue";

export default {
  components: { ...sp },
};
</script>
