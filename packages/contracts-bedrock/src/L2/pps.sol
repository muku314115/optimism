// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract PPSContract {
    uint256 public PPS = 1;

    function updatePPS(uint256 amount) public { // this is not meant to be public, just made this way for easy testing
        require(
            msg.sender == 0xDeaDDEaDDeAdDeAdDEAdDEaddeAddEAdDEAd0001,
            "Only system address can update PPS"
        );
    PPS = amount;
    }
}
