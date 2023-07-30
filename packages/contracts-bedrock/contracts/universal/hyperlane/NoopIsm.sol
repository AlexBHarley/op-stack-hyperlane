// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

contract NoopIsm {
    constructor() {}

    function verify(bytes calldata, bytes calldata) public returns (bool) {
        return true;
    }
}
