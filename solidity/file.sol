pragma solidity ^0.5.0;

contract MetaData {
    string public core;
    address public oracle;
    bool public locked;
    address public owner;
    bytes32 public codeTx;
    bool public hashTx;

    event LockEvent(bool);

    constructor() public {
        locked =false;
        oracle = 0x96216849c49358B10257cb55b28eA603c874b05E;
        owner = msg.sender;
        hashTx = false;
    }
  
    function update(string memory _data) public {
        if (locked) {
            require(msg.sender == oracle);
        } else {
            require(msg.sender == owner);
        }
        core = _data;
    }
    
    function lock() public {
        require(locked == false);
        require(msg.sender == oracle); 
        locked = true;
        emit LockEvent(locked);
    }

    function unlock() public {
        require(locked == true);
        require(msg.sender == oracle);
        locked = false;
        emit LockEvent(locked);
    } 
    
    function setCode(bytes32  _codeTx) public {
        require(msg.sender == owner); 
        require(hashTx == false);
        codeTx = _codeTx;
        hashTx = true;
    }
}
