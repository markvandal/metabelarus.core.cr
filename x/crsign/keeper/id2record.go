package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

func (k Keeper) GetRecordsFromId(ctx sdk.Context, identity string) types.Id2Record {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2RecordKey))
	var id2record types.Id2Record
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.Id2RecordKey+identity)), &id2record)
	return id2record
}

func (k Keeper) AddRecord2Id(ctx sdk.Context, identity string, id string) error {
	record2id := k.GetRecordsFromId(ctx, identity)
	if len(record2id.Records) > 255 {
		return sdkerrors.Wrap(
			types.ErrRecordLimit,
			fmt.Sprintf(" One needs a special administrative rights to create more records for this ID %s", identity),
		)
	}

	for _, record := range record2id.Records {
		if record == id {
			return nil
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

	return nil
}

func (k Keeper) DeleteRecordFromId(ctx sdk.Context, identity string, id string) {
	record2id := k.GetRecordsFromId(ctx, identity)

	found := false
	var foundRecord int
	for idx, record := range record2id.Records {
		if record == id {
			foundRecord = idx
			found = true
			break
		}
	}
	if !found {
		return
	}

	record2id.Records[foundRecord] = record2id.Records[len(record2id.Records)-1]
	record2id.Records = record2id.Records[:len(record2id.Records)-1]

	prefix.NewStore(
		ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2RecordKey),
	).Set(
		types.KeyPrefix(types.Id2RecordKey+identity),
		k.cdc.MustMarshalBinaryBare(&record2id),
	)
}
