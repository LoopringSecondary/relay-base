/*

  Copyright 2017 Loopring Project Ltd (Loopring Foundation).

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

package types

import (
	"math/big"
)

type TokenRegisterEvent struct {
	Token           Address
	ContractAddress Address
	Symbol          string
	BlockNumber     *big.Int
	Time            *big.Int
}

type TokenUnRegisterEvent struct {
	Token           Address
	ContractAddress Address
	Symbol          string
	BlockNumber     *big.Int
	Time            *big.Int
}

type RingHashSubmittedEvent struct {
	RingHash        Hash
	RingMiner       Address
	ContractAddress Address
	TxHash          Hash
	BlockNumber     *big.Int
	Time            *big.Int
}

type AddressAuthorizedEvent struct {
	Protocol        Address
	ContractAddress Address
	Number          int
	BlockNumber     *big.Int
	Time            *big.Int
}

type AddressDeAuthorizedEvent struct {
	Protocol        Address
	ContractAddress Address
	Number          int
	BlockNumber     *big.Int
	Time            *big.Int
}

type TransferEvent struct {
	From            Address
	To              Address
	ContractAddress Address
	Value           *big.Int
	BlockNumber     *big.Int
	Time            *big.Int
}

type ApprovalEvent struct {
	Owner           Address
	Spender         Address
	ContractAddress Address
	Value           *big.Int
	BlockNumber     *big.Int
	Time            *big.Int
}

type OrderFilledEvent struct {
	RingHash        Hash
	PreOrderHash    Hash
	OrderHash       Hash
	NextOrderHash   Hash
	TxHash          Hash
	ContractAddress Address
	Owner           Address
	TokenS          Address
	TokenB          Address
	RingIndex       *big.Int
	Time            *big.Int
	BlockNumber     *big.Int
	AmountS         *big.Int
	AmountB         *big.Int
	LrcReward       *big.Int
	LrcFee          *big.Int
	SplitS          *big.Int
	SplitB          *big.Int
	Market          string
	FillIndex       *big.Int
}

type OrderCancelledEvent struct {
	OrderHash       Hash
	TxHash          Hash
	ContractAddress Address
	Time            *big.Int
	BlockNumber     *big.Int
	AmountCancelled *big.Int
}

type CutoffEvent struct {
	Owner           Address
	ContractAddress Address
	TxHash          Hash
	Time            *big.Int
	BlockNumber     *big.Int
	Cutoff          *big.Int
}

type RingMinedEvent struct {
	RingIndex          *big.Int
	Time               *big.Int
	BlockNumber        *big.Int
	TotalLrcFee        *big.Int
	TradeAmount        int
	RingHash           Hash
	TxHash             Hash
	Miner              Address
	FeeRecipient       Address
	ContractAddress    Address
	IsRingHashReserved bool
}

type ApproveMethodEvent struct {
	From            Address
	To              Address
	ContractAddress Address
	TxHash          Hash
	Spender         Address
	Value           *big.Int
	Time            *big.Int
	BlockNumber     *big.Int
	Owner           Address
	Success         bool
}

type SubmitRingMethodEvent struct {
	TxHash       Hash
	UsedGas      *big.Int
	UsedGasPrice *big.Int
	Err          error
}

type RingHashSubmitMethodEvent struct {
	RingMiner    Address
	RingHash     Hash
	TxHash       Hash
	UsedGas      *big.Int
	UsedGasPrice *big.Int
	Err          error
}

type BatchSubmitRingHashMethodEvent struct {
	RingHashMinerMap map[Hash]Address
	TxHash           Hash
	UsedGas          *big.Int
	UsedGasPrice     *big.Int
	Err              error
}

type RingSubmitFailedEvent struct {
	RingHash Hash
	Err      error
}

type ForkedEvent struct {
	DetectedBlock *big.Int
	DetectedHash  Hash
	ForkBlock     *big.Int
	ForkHash      Hash
}

type BlockEvent struct {
	BlockNumber *big.Int
	BlockHash   Hash
}
