// Code generated by "stringer -type=contractTypeImpl -linecomment"; DO NOT EDIT.

package testutil

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OriginType-0]
	_ = x[MessageHarnessType-1]
	_ = x[BaseMessageHarnessType-2]
	_ = x[RequestHarnessType-3]
	_ = x[OriginHarnessType-4]
	_ = x[StateHarnessType-5]
	_ = x[SnapshotHarnessType-6]
	_ = x[AttestationHarnessType-7]
	_ = x[TipsHarnessType-8]
	_ = x[HeaderHarnessType-9]
	_ = x[DestinationHarnessType-10]
	_ = x[SummitHarnessType-11]
	_ = x[SummitType-12]
	_ = x[DestinationType-13]
	_ = x[AgentsTestContractType-14]
	_ = x[TestClientType-15]
	_ = x[PingPongClientType-16]
	_ = x[LightManagerHarnessType-17]
	_ = x[BondingManagerHarnessType-18]
	_ = x[LightManagerType-19]
	_ = x[BondingManagerType-20]
	_ = x[GasDataHarnessType-21]
	_ = x[GasOracleType-22]
	_ = x[InboxType-23]
	_ = x[LightInboxType-24]
}

const _contractTypeImpl_name = "OriginMessageHarnessBaseMessageHarnessRequestHarnessOriginHarnessStateHarnessTypeSnapshotHarnessTypeAttestationHarnessTypeTipsHarnessTypeHeaderHarnessTypeDestinationHarnessSummitHarnessSummitDestinationAgentsTestContractTestClientPingPongClientLightManagerHarnessBondingManagerHarnessLightManagerBondingManagerGasDataHarnessTypeGasOracleInboxLightInbox"

var _contractTypeImpl_index = [...]uint16{0, 6, 20, 38, 52, 65, 81, 100, 122, 137, 154, 172, 185, 191, 202, 220, 230, 244, 263, 284, 296, 310, 328, 337, 342, 352}

func (i contractTypeImpl) String() string {
	if i < 0 || i >= contractTypeImpl(len(_contractTypeImpl_index)-1) {
		return "contractTypeImpl(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _contractTypeImpl_name[_contractTypeImpl_index[i]:_contractTypeImpl_index[i+1]]
}