package utils

import "context"

/*
	Operations on structures in proto files
*/

// PbJson2String Output pb parameters in the form of JSON strings
func PbJson2String(ctx context.Context, pbStruct any) string {
	s, err := ToJsonString(pbStruct)
	if err != nil {
		return ""
	}

	return s
}
