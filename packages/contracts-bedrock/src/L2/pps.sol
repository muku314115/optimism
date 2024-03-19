// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract PPSContract {
    uint32 public PPS = 1;

    function updatePPS(uint32 amount) public {
        require(
            msg.sender == 0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001,
            "Only system address can update PPS"
        );
    PPS = amount;
    }
}
