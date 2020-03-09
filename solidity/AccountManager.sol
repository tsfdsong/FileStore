pragma solidity ^0.5.0;

// account manager
contract AccountManager {
    //save smart contract accounts, user address => contract address
    mapping(address => address) public accounts;

    address public owner;

    event Record(address indexed account, address indexed store);
    event DeleteRecord(address indexed account);

    constructor() public {
        owner = msg.sender;
    }

    //set
    function set(address contractAddr) public {
        require(accounts[msg.sender] == address(0), "account store contract has exist");
        accounts[msg.sender] = contractAddr;
        emit Record(msg.sender, contractAddr);
    }

    //get
    function get(address account) public view returns (address) {
        return accounts[account];
    }

    //delete
    function clear(address account) public {
        require(msg.sender == owner, "you are not allowed to delete it");
        delete accounts[account];
        emit DeleteRecord(account);
    }
}