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

type OrderStatus uint8

const (
	ORDER_UNKNOWN OrderStatus = iota
	ORDER_NEW
	ORDER_PARTIAL
	ORDER_FINISHED
	ORDER_CANCEL
	ORDER_CUTOFF
)

type Order interface {
	GetProtocol() Address
	GetTokenS() Address
	GetTokenB() Address
	GetAmountS() *big.Int
	GetAmountB() *big.Int
	GetTimestamp() int64
	GetTtl() *big.Int
	GetLrcFee() *big.Int
	GetBuyNoMoreThanAmountB() bool
	GetMarginSplitPercentage() uint8
	GetPrice() *big.Rat

	// functions
	GenerateHash() Hash
	GeneratePrice()
	GenerateAndSetSignature() error
	ValidateSignatureValues() bool
}
