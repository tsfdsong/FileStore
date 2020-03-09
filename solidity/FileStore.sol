pragma solidity ^0.5.0;
import "./file.sol";

/// @title File storage
contract FileStore is MetaData {
    function updateHash(string memory hash) public {
        require(msg.sender == owner);
        update(hash);
    }

    function indexHash() public view returns (string  memory) {
        return core;
    }
}