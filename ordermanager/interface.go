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

package ordermanager

import "math/big"

type OrderManager interface {
	MinerOrders(protocol, tokenS, tokenB string, length int, startBlockNumber, endBlockNumber int64)

	GetOrderBook(protocol, tokenS, tokenB string, length int)

	GetOrders(query map[string]interface{}, pageIndex, pageSize int)

	GetOrderByHash(hash string)

	UpdateBroadcastTimeByHash(hash string, bt int) error

	FillsPageQuery(query map[string]interface{}, pageIndex, pageSize int)

	RingMinedPageQuery(query map[string]interface{}, pageIndex, pageSize int)

	IsOrderCutoff(protocol, owner string, createTime *big.Int) bool

	IsOrderFullFinished() bool

	GetFrozenAmount(owner, token string, statusSet []int) (*big.Int, error)

	GetFrozenLRCFee(owner string, statusSet []int) (*big.Int, error)
}
