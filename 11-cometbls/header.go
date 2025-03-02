package cometbls

import (
	"time"

	errorsmod "cosmossdk.io/errors"

	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/v8/modules/core/23-commitment/types"
	"github.com/cosmos/ibc-go/v8/modules/core/exported"
)

var _ exported.ClientMessage = (*Header)(nil)

// ConsensusState returns the updated consensus state associated with the header
func (h Header) ConsensusState() *ConsensusState {
	return &ConsensusState{
		Timestamp:          uint64(h.GetTime().UnixNano()),
		Root:               commitmenttypes.NewMerkleRoot(h.SignedHeader.AppHash),
		NextValidatorsHash: h.SignedHeader.NextValidatorsHash,
	}
}

// ClientType defines that the Header is a Tendermint consensus algorithm
func (Header) ClientType() string {
	return ClientType
}

// GetHeight returns the current height. It returns 0 if the tendermint
// header is nil.
// NOTE: the header.Header is checked to be non nil in ValidateBasic.
func (h Header) GetHeight() exported.Height {
	// revision := clienttypes.ParseChainID(h.SignedHeader.ChainID)
	return clienttypes.NewHeight(h.TrustedHeight.RevisionNumber, uint64(h.SignedHeader.Height))
}

// GetTime returns the current block timestamp. It returns a zero time if
// the tendermint header is nil.
// NOTE: the header.Header is checked to be non nil in ValidateBasic.
func (h Header) GetTime() time.Time {
	return h.SignedHeader.Time
}

// ValidateBasic calls the SignedHeader ValidateBasic function and checks
// that validatorsets are not nil.
// NOTE: TrustedHeight and TrustedValidators may be empty when creating client
// with MsgCreateClient
func (h Header) ValidateBasic() error {
	if h.SignedHeader == nil {
		return errorsmod.Wrap(clienttypes.ErrInvalidHeader, "tendermint signed header cannot be nil")
	}

	// TrustedHeight is less than Header for updates and misbehaviour
	if h.TrustedHeight.GTE(h.GetHeight()) {
		return errorsmod.Wrapf(ErrInvalidHeaderHeight, "TrustedHeight %d must be less than header height %d",
			h.TrustedHeight, h.GetHeight())
	}

	return nil
}
