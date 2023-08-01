// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Script } from "forge-std/Script.sol";
import { console2 as console } from "forge-std/console2.sol";

import { Predeploys } from "../contracts/libraries/Predeploys.sol";
import {
    OptimismMintableERC20Factory
} from "../contracts/universal/OptimismMintableERC20Factory.sol";
import { L2StandardBridge } from "../contracts/L2/L2StandardBridge.sol";

contract Deploy is Script {
    function run() public {
        vm.startBroadcast();
        console.log("Deploying ERC20");

        OptimismMintableERC20Factory factory = OptimismMintableERC20Factory(
            Predeploys.OPTIMISM_MINTABLE_ERC20_FACTORY
        );

        address erc20 = factory.createOptimismMintableERC20(
            0x6523Fd7e1B28113aB7350B7C312b541fd4866492,
            "l2_MyToken",
            "l2_TOKE"
        );
        console.log("Deployed mintable erc20 to", erc20);

        L2StandardBridge bridge = L2StandardBridge(payable(Predeploys.L2_STANDARD_BRIDGE));
        bridge.setTokenEnabledForMailboxWithdrawals(erc20, true);

        vm.stopBroadcast();
    }
}
