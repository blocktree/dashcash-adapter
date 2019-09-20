/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package dashcash

import (
	"github.com/assetsadapterstore/dashcash-adapter/dscTransaction"
	"github.com/blocktree/go-owcrypt"
)

const (
	//币种
	Symbol    = "DSC"
	CurveType = owcrypt.ECC_CURVE_SECP256K1
	Decimals  = int32(8)
)

var (
	MainNetAddressPrefix = dscTransaction.AddressPrefix{P2PKHPrefix: []byte{0x1e}, P2WPKHPrefix: []byte{0x0d}, P2SHPrefix: nil, Bech32Prefix: "dsc"}
	TestNetAddressPrefix = dscTransaction.AddressPrefix{P2PKHPrefix: []byte{0x6f}, P2WPKHPrefix: []byte{0x13}, P2SHPrefix: nil, Bech32Prefix: "dsc"}
)
