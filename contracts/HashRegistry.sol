// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

contract HashRegistry {
    mapping(bytes32 => bool) public hashes;
    bytes32[] public hashList;
    mapping(bytes32 => uint256) public hashTimestamps;
    mapping(bytes32 => address) public hashOwners;
    mapping(bytes32 => string) public hashCIDs;

    event HashRegistrado(bytes32 hash, address owner, uint256 timestamp, string cid);

    function registrarHash(bytes32 hash, string memory cid) public {
        require(!hashes[hash], "Hash already registered!");
        hashes[hash] = true;
        hashList.push(hash);
        hashTimestamps[hash] = block.timestamp;
        hashOwners[hash] = msg.sender;
        hashCIDs[hash] = cid;
        emit HashRegistrado(hash, msg.sender, block.timestamp, cid);
    }

    function verificarHash(bytes32 hash) public view returns (bool) {
        return hashes[hash];
    }

    function getTotalHashes() public view returns (uint256) {
        return hashList.length;
    }

    function getHashByIndex(uint256 index) public view returns (bytes32) {
        require(index < hashList.length, "Index out of bounds");
        return hashList[index];
    }

    function getHashDetails(bytes32 hash) public view returns (
        bool exists,
        uint256 timestamp,
        address owner,
        string memory cid
    ) {
        return (
            hashes[hash],
            hashTimestamps[hash],
            hashOwners[hash],
            hashCIDs[hash]
        );
    }

    function getAllHashes() public view returns (bytes32[] memory) {
        return hashList;
    }

    function getCID(bytes32 hash) public view returns (string memory) {
        require(hashes[hash], "Hash not found");
        return hashCIDs[hash];
    }
}
