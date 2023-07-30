// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Script } from "forge-std/Script.sol";
import { console2 as console } from "forge-std/console2.sol";
import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

import { L1StandardBridge } from "../contracts/L1/L1StandardBridge.sol";
import { NoopIsm } from "../contracts/universal/hyperlane/NoopIsm.sol";

contract L1Token is ERC20 {
    constructor() ERC20("L1 Token", "TOKE") {
        _mint(msg.sender, 1_000_000 ether);
    }
}

contract Deploy is Script {
    function run() public {
        vm.startBroadcast();

        ERC20 token = new L1Token();
        console.log("Token deployed at:", address(token));

        NoopIsm ism = new NoopIsm();
        console.log("ISM deployed at:", address(ism));

        L1StandardBridge l1Bridge = L1StandardBridge(
            payable(0x0165878A594ca255338adfa4d48449f69242Eb8F)
        );

        l1Bridge.setWithdrawalIsm(address(token), address(ism));

        vm.stopBroadcast();
    }
}
