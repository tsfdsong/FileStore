pragma solidity ^0.5.0;
contract Storage {
    address public owner;
    address public oracle;
    uint public expireAt;
    string public indexHash;
    State public stat;
    enum State {
        Closed,
        Authorized,
        Opened
    }
    event StateChange(State target);
    // This is the constructor whose code is
    // run only when the contract is created.
    constructor() public {
        owner = msg.sender;
        stat = State.Closed;
    }
    function updateHash(string memory hash) public {
        if (stat == State.Opened) {
            require(msg.sender == oracle);
        } else { // Authorized or Closed
            require(msg.sender == owner);
        }
        indexHash = hash;
    }
    function isExpriy() public view returns(bool) {
        return (expireAt <= now);
    }
    function updateStat(State target) private {
        stat = target;
        emit StateChange(stat);
    }
    function authorize(address account) public {
        require(owner == msg.sender && account != owner);
        require(stat == State.Closed || stat == State.Authorized);
        oracle = account;
        updateStat(State.Authorized);
    }
    function open(uint16 duration) public {
        require(stat == State.Authorized);
        require(duration > 0);
        expireAt = now + duration;
        updateStat(State.Opened);
    }
    function close(string memory hash) public {
        require(msg.sender == oracle);
        require(stat == State.Opened);
        require(expireAt <= now);
        indexHash = hash;
        expireAt = 0;
        updateStat(State.Closed);
    }
}