package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

func (k Keeper) GetRecordsFromId(ctx sdk.Context, identity string) types.Id2Record {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2RecordKey))
	var id2record types.Id2Record
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.Id2RecordKey+identity)), &id2record)
	return id2record
}

func (k Keeper) AddRecord2Id(ctx sdk.Context, identity string, id string) {
	record2id := k.GetRecordsFromId(ctx, identity)
	for _, record := range record2id.Records {
		if record == id {
			return
		}
	}

	if record2id.Records == nil || len(record2id.Records) == 0 {
		record2id.Records = []string{id}
	} else {
		record2id.Records = append(record2id.Records, id)
	}

	prefix.NewStore(
		ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2RecordKey),
	).Set(
		types.KeyPrefix(types.Id2RecordKey+identity),
		k.cdc.MustMarshalBinaryBare(&record2id),
	)
}
