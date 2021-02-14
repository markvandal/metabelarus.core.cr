package mbutils

import (
	"strconv"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ModuleName   = "MetaBelarus.Utils"
	ErrSizeLimit = sdkerrors.Register(ModuleName, 10110, "Value size limit is reached")
	ErrValueInt  = sdkerrors.Register(ModuleName, 10120, "Value should be number")
)

func EnumMapToList(mp map[int32]string) []string {
	arr := []string{}
	for _, v := range mp {
		arr = append(arr, v)
	}

	return arr
}

func ValidateId(id string, param string) error {
	if len([]rune(id)) > 20 {
		return sdkerrors.Wrapf(ErrSizeLimit, "%s can't be longer than 20 symbols", param)
	}

	if _, err := strconv.Atoi(id); err != nil {
		return sdkerrors.Wrapf(
			ErrValueInt,
			"%s value should be a number given: %s, err: %s",
			param, id, err,
		)
	}

	return nil
}

func ValidateKey(id string, param string) error {
	if len([]rune(id)) > 1024 {
		return sdkerrors.Wrapf(ErrSizeLimit, "%s can't be longer than 1024 symbols", param)
	}

	return nil
}

func ValidateData(value string, param string) error {
	if len([]rune(value)) > 2048 {
		return sdkerrors.Wrapf(ErrSizeLimit, "%s can't be longer than 2048 symbols", param)
	}

	return nil
}
