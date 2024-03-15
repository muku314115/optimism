// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

// Testing utilities
import { CommonTest } from "test/setup/CommonTest.sol";

// Libraries
import { Encoding } from "src/libraries/Encoding.sol";

// Target contract
import { L1Block, NotDepositor, DependencySetSizeMismatch } from "src/L2/L1Block.sol";

contract L1BlockTest is CommonTest {
    address depositor;

    /// @dev Sets up the test suite.
    function setUp() public virtual override {
        super.setUp();
        depositor = l1Block.DEPOSITOR_ACCOUNT();
    }
}

contract L1BlockBedrock_Test is L1BlockTest {
    // @dev Tests that `setL1BlockValues` updates the values correctly.
    function testFuzz_updatesValues_succeeds(
        uint64 n,
        uint64 t,
        uint256 b,
        bytes32 h,
        uint64 s,
        bytes32 bt,
        uint256 fo,
        uint256 fs
    )
        external
    {
        vm.prank(depositor);
        l1Block.setL1BlockValues(n, t, b, h, s, bt, fo, fs);
        assertEq(l1Block.number(), n);
        assertEq(l1Block.timestamp(), t);
        assertEq(l1Block.basefee(), b);
        assertEq(l1Block.hash(), h);
        assertEq(l1Block.sequenceNumber(), s);
        assertEq(l1Block.batcherHash(), bt);
        assertEq(l1Block.l1FeeOverhead(), fo);
        assertEq(l1Block.l1FeeScalar(), fs);
    }

    /// @dev Tests that `setL1BlockValues` can set max values.
    function test_updateValues_succeeds() external {
        vm.prank(depositor);
        l1Block.setL1BlockValues({
            _number: type(uint64).max,
            _timestamp: type(uint64).max,
            _basefee: type(uint256).max,
            _hash: keccak256(abi.encode(1)),
            _sequenceNumber: type(uint64).max,
            _batcherHash: bytes32(type(uint256).max),
            _l1FeeOverhead: type(uint256).max,
            _l1FeeScalar: type(uint256).max
        });
    }

    /// @dev Tests that `setL1BlockValues` fails if sender address is not the depositor
    function test_updatesValues_notDepositor_fails() external {
        vm.expectRevert("L1Block: only the depositor account can set L1 block values");
        l1Block.setL1BlockValues({
            _number: type(uint64).max,
            _timestamp: type(uint64).max,
            _basefee: type(uint256).max,
            _hash: keccak256(abi.encode(1)),
            _sequenceNumber: type(uint64).max,
            _batcherHash: bytes32(type(uint256).max),
            _l1FeeOverhead: type(uint256).max,
            _l1FeeScalar: type(uint256).max
        });
    }
}

contract L1BlockEcotone_Test is L1BlockTest {
    /// @dev Tests that setL1BlockValuesEcotone updates the values appropriately.
    function testFuzz_setL1BlockValuesEcotone_succeeds(
        uint32 baseFeeScalar,
        uint32 blobBaseFeeScalar,
        uint64 sequenceNumber,
        uint64 timestamp,
        uint64 number,
        uint256 baseFee,
        uint256 blobBaseFee,
        bytes32 hash,
        bytes32 batcherHash
    )
        external
    {
        bytes memory functionCallDataPacked = Encoding.encodeSetL1BlockValuesEcotone(
            baseFeeScalar, blobBaseFeeScalar, sequenceNumber, timestamp, number, baseFee, blobBaseFee, hash, batcherHash
        );

        vm.prank(depositor);
        (bool success,) = address(l1Block).call(functionCallDataPacked);
        assertTrue(success, "Function call failed");

        assertEq(l1Block.baseFeeScalar(), baseFeeScalar);
        assertEq(l1Block.blobBaseFeeScalar(), blobBaseFeeScalar);
        assertEq(l1Block.sequenceNumber(), sequenceNumber);
        assertEq(l1Block.timestamp(), timestamp);
        assertEq(l1Block.number(), number);
        assertEq(l1Block.basefee(), baseFee);
        assertEq(l1Block.blobBaseFee(), blobBaseFee);
        assertEq(l1Block.hash(), hash);
        assertEq(l1Block.batcherHash(), batcherHash);

        // ensure we didn't accidentally pollute the 128 bits of the sequencenum+scalars slot that
        // should be empty
        bytes32 scalarsSlot = vm.load(address(l1Block), bytes32(uint256(3)));
        bytes32 mask128 = hex"FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF00000000000000000000000000000000";

        assertEq(0, scalarsSlot & mask128);

        // ensure we didn't accidentally pollute the 128 bits of the number & timestamp slot that
        // should be empty
        bytes32 numberTimestampSlot = vm.load(address(l1Block), bytes32(uint256(0)));
        assertEq(0, numberTimestampSlot & mask128);
    }

    /// @dev Tests that `setL1BlockValuesEcotone` succeeds if sender address is the depositor
    function test_setL1BlockValuesEcotone_isDepositor_succeeds() external {
        bytes memory functionCallDataPacked = Encoding.encodeSetL1BlockValuesEcotone(
            type(uint32).max,
            type(uint32).max,
            type(uint64).max,
            type(uint64).max,
            type(uint64).max,
            type(uint256).max,
            type(uint256).max,
            bytes32(type(uint256).max),
            bytes32(type(uint256).max)
        );

        vm.prank(depositor);
        (bool success,) = address(l1Block).call(functionCallDataPacked);
        assertTrue(success, "function call failed");
    }

    /// @dev Tests that `setL1BlockValuesEcotone` fails if sender address is not the depositor
    function test_setL1BlockValuesEcotone_notDepositor_fails() external {
        bytes memory functionCallDataPacked = Encoding.encodeSetL1BlockValuesEcotone(
            type(uint32).max,
            type(uint32).max,
            type(uint64).max,
            type(uint64).max,
            type(uint64).max,
            type(uint256).max,
            type(uint256).max,
            bytes32(type(uint256).max),
            bytes32(type(uint256).max)
        );

        (bool success, bytes memory data) = address(l1Block).call(functionCallDataPacked);
        assertTrue(!success, "function call should have failed");
        // make sure return value is the expected function selector for "NotDepositor()"
        assertEq(bytes4(data), NotDepositor.selector);
    }
}

contract L1BlockInterop_Test is L1BlockTest {
    /// @dev Tests that setL1BlockValuesInterop updates the values appropriately.
    function testFuzz_setL1BlockValuesInterop_succeeds(
        uint32 baseFeeScalar,
        uint32 blobBaseFeeScalar,
        uint64 sequenceNumber,
        uint64 timestamp,
        uint64 number,
        uint256 baseFee,
        uint256 blobBaseFee,
        bytes32 hash,
        bytes32 batcherHash,
        uint256[] calldata dependencySet
    )
        external
    {
        vm.assume(dependencySet.length <= type(uint8).max);

        bytes memory functionCallDataPacked = Encoding.encodeSetL1BlockValuesInterop({
            _baseFeeScalar: baseFeeScalar,
            _blobBaseFeeScalar: blobBaseFeeScalar,
            _sequenceNumber: sequenceNumber,
            _timestamp: timestamp,
            _number: number,
            _baseFee: baseFee,
            _blobBaseFee: blobBaseFee,
            _hash: hash,
            _batcherHash: batcherHash,
            _dependencySet: dependencySet
        });

        vm.prank(depositor);
        (bool success,) = address(l1Block).call(functionCallDataPacked);
        assertTrue(success, "Function call failed");

        assertEq(l1Block.baseFeeScalar(), baseFeeScalar);
        assertEq(l1Block.blobBaseFeeScalar(), blobBaseFeeScalar);
        assertEq(l1Block.sequenceNumber(), sequenceNumber);
        assertEq(l1Block.timestamp(), timestamp);
        assertEq(l1Block.number(), number);
        assertEq(l1Block.basefee(), baseFee);
        assertEq(l1Block.blobBaseFee(), blobBaseFee);
        assertEq(l1Block.hash(), hash);
        assertEq(l1Block.batcherHash(), batcherHash);
        assertEq(l1Block.dependencySetSize(), dependencySet.length);
        for (uint256 i = 0; i < dependencySet.length; i++) {
            assertEq(l1Block.dependencySet(i), dependencySet[i]);
            assertTrue(l1Block.isInDependencySet(dependencySet[i]));
        }

        // ensure we didn't accidentally pollute the 128 bits of the sequencenum+scalars slot that
        // should be empty
        bytes32 scalarsSlot = vm.load(address(l1Block), bytes32(uint256(3)));
        bytes32 mask128 = hex"FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF00000000000000000000000000000000";

        assertEq(0, scalarsSlot & mask128);

        // ensure we didn't accidentally pollute the 128 bits of the number & timestamp slot that
        // should be empty
        bytes32 numberTimestampSlot = vm.load(address(l1Block), bytes32(uint256(0)));
        assertEq(0, numberTimestampSlot & mask128);
    }

    /// @dev Tests that `setL1BlockValuesInterop` succeeds if sender address is the depositor
    function test_setL1BlockValuesInterop_isDepositor_succeeds() external {
        bytes memory functionCallDataPacked = Encoding.encodeSetL1BlockValuesInterop(
            type(uint32).max,
            type(uint32).max,
            type(uint64).max,
            type(uint64).max,
            type(uint64).max,
            type(uint256).max,
            type(uint256).max,
            bytes32(type(uint256).max),
            bytes32(type(uint256).max),
            new uint256[](0)
        );

        vm.prank(depositor);
        (bool success,) = address(l1Block).call(functionCallDataPacked);
        assertTrue(success, "function call failed");
    }

    /// @dev Tests that `setL1BlockValuesInterop` fails if sender address is not the depositor
    function test_setL1BlockValuesInterop_isDepositor_fails() external {
        bytes memory functionCallDataPacked = Encoding.encodeSetL1BlockValuesInterop(
            type(uint32).max,
            type(uint32).max,
            type(uint64).max,
            type(uint64).max,
            type(uint64).max,
            type(uint256).max,
            type(uint256).max,
            bytes32(type(uint256).max),
            bytes32(type(uint256).max),
            new uint256[](0)
        );

        (bool success, bytes memory data) = address(l1Block).call(functionCallDataPacked);
        assertTrue(!success, "function call should have failed");
        // make sure return value is the expected function selector for "NotDepositor()"
        assertEq(bytes4(data), NotDepositor.selector);
    }

    /// @dev Tests that `setL1BlockValuesInterop` fails if _dependencySetSize is the same as
    ///      the length of _dependencySet. (happy path)
    function testFuzz_setL1BlockValuesInterop_dependencySetSizeMatch_succeeds(uint256[] calldata dependencySet)
        external
    {
        vm.assume(dependencySet.length <= type(uint8).max);

        bytes memory functionCallDataPacked = Encoding.encodeSetL1BlockValuesInterop(
            type(uint32).max,
            type(uint32).max,
            type(uint64).max,
            type(uint64).max,
            type(uint64).max,
            type(uint256).max,
            type(uint256).max,
            bytes32(type(uint256).max),
            bytes32(type(uint256).max),
            dependencySet
        );

        vm.prank(depositor);
        (bool success,) = address(l1Block).call(functionCallDataPacked);
        assertTrue(success, "function call failed");
    }

    /// @dev Tests that `setL1BlockValuesInterop` fails if _dependencySetSize is not the same as
    ///      the length of _dependencySet. (bad path)
    function testFuzz_setL1BlockValuesInterop_dependencySetSizeMatch_fails(
        uint8 notDependencySetSize,
        uint256[] calldata dependencySet
    )
        external
    {
        vm.assume(dependencySet.length <= type(uint8).max);
        vm.assume(notDependencySetSize != dependencySet.length);

        bytes memory functionCallDataPacked = abi.encodePacked(
            bytes4(keccak256("setL1BlockValuesInterop()")),
            type(uint32).max,
            type(uint32).max,
            type(uint64).max,
            type(uint64).max,
            type(uint64).max,
            type(uint256).max,
            type(uint256).max,
            bytes32(type(uint256).max),
            bytes32(type(uint256).max),
            notDependencySetSize,
            dependencySet
        );

        vm.prank(depositor);
        (bool success, bytes memory data) = address(l1Block).call(functionCallDataPacked);
        assertTrue(!success, "function call should have failed");
        // make sure return value is the expected function selector for "DependencySetSizeMismatch()"
        assertEq(bytes4(data), DependencySetSizeMismatch.selector);
    }

    /// @dev Tests that an arbitrary dependency set can be set and that ìsInDependencySet returns
    ///      the expected results.
    function testFuzz_isInDependencySet_succeeds(
        uint32 baseFeeScalar,
        uint32 blobBaseFeeScalar,
        uint64 sequenceNumber,
        uint64 timestamp,
        uint64 number,
        uint256 baseFee,
        uint256 blobBaseFee,
        bytes32 hash,
        bytes32 batcherHash,
        uint256[] calldata dependencySet
    )
        external
    {
        vm.assume(dependencySet.length <= type(uint8).max);

        bytes memory functionCallDataPacked = Encoding.encodeSetL1BlockValuesInterop({
            _baseFeeScalar: baseFeeScalar,
            _blobBaseFeeScalar: blobBaseFeeScalar,
            _sequenceNumber: sequenceNumber,
            _timestamp: timestamp,
            _number: number,
            _baseFee: baseFee,
            _blobBaseFee: blobBaseFee,
            _hash: hash,
            _batcherHash: batcherHash,
            _dependencySet: dependencySet
        });

        vm.prank(depositor);
        (bool success,) = address(l1Block).call(functionCallDataPacked);
        assertTrue(success, "Function call failed");

        assertEq(l1Block.dependencySetSize(), dependencySet.length);

        for (uint256 i = 0; i < dependencySet.length; i++) {
            assertTrue(l1Block.isInDependencySet(dependencySet[i]));
        }
    }

    /// @dev Tests that `isInDependencySet` returns true when the current chain ID is passed as the input
    function test_isInDependencySet_isChainId_succeeds() external {
        assertTrue(l1Block.isInDependencySet(block.chainid));
    }

    /// @dev Tests that `isInDependencySet` fails when the input chain ID is not in the dependency set
    function testFuzz_isInDependencySet_fails(uint256 chainId) external {
        vm.assume(chainId != 1);

        uint256[] memory dependencySet = new uint256[](1);
        dependencySet[0] = 1;

        bytes memory functionCallDataPacked = Encoding.encodeSetL1BlockValuesInterop({
            _baseFeeScalar: 0,
            _blobBaseFeeScalar: 0,
            _sequenceNumber: 0,
            _timestamp: 0,
            _number: 0,
            _baseFee: 0,
            _blobBaseFee: 0,
            _hash: bytes32(0),
            _batcherHash: 0,
            _dependencySet: dependencySet
        });

        vm.prank(depositor);
        (bool success,) = address(l1Block).call(functionCallDataPacked);
        assertTrue(success, "Function call failed");

        assertFalse(l1Block.isInDependencySet(chainId));
    }

    /// @dev Tests that `isInDependencySet` returns false when the dependency set is empty
    function testFuzz_isInDependencySet_dependencySetEmpty_fails(uint256 chainId) external {
        assertTrue(l1Block.dependencySetSize() == 0);
        assertFalse(l1Block.isInDependencySet(chainId));
    }
}
