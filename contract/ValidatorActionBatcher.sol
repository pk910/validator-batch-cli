// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;
 
contract ValidatorActionBatcher {
    address public constant WITHDRAWAL_SYSTEM_CONTRACT = 0x00000961Ef480Eb55e80D19ad83579A64c007002;
    address public constant CONSOLIDATION_SYSTEM_CONTRACT = 0x0000BBdDc7CE488642fb579F8B00f3a590007251;
 
    modifier onlySelf {
        _checkSelf();
        _;
    }
 
    function _checkSelf() internal view virtual {
        if(msg.sender != address(this)) {
            revert("self call only");
        }
    }
 
    receive() external payable {
    }
 
    fallback() external payable {
    }
 
    function getWithdrawalFee() public view returns (uint256 fee) {
        (bool ok, bytes memory result) = WITHDRAWAL_SYSTEM_CONTRACT.staticcall("");
        require(ok);
 
        fee = abi.decode(result, (uint256));
    }
 
    function batchWithdraw(bytes[] calldata data, uint256 feeLimit) public onlySelf {
        uint256 fee = getWithdrawalFee();
 
        require(feeLimit == 0 || (fee <= feeLimit), "fee too high");
        require(address(this).balance >= fee * data.length, "balance too low");
 
        for (uint256 i; i < data.length; i++) {
            require(data[i].length == 56, "invalid withdrawal");
 
            (bool success, ) = WITHDRAWAL_SYSTEM_CONTRACT.call{value: fee}(data[i]);
            require(success, "call failed");
        }
    }
 
    function getConsolidationFee() public view returns (uint256 fee) {
        (bool ok, bytes memory result) = CONSOLIDATION_SYSTEM_CONTRACT.staticcall("");
        require(ok);
 
        fee = abi.decode(result, (uint256));
    }
 
    function batchConsolidate(bytes[] calldata data, uint256 feeLimit) public onlySelf {
        uint256 fee = getConsolidationFee();
 
        require(feeLimit == 0 || (fee <= feeLimit), "fee too high");
        require(address(this).balance >= fee * data.length, "balance too low");
 
        for (uint256 i; i < data.length; i++) {
            require(data[i].length == 96, "invalid consolidation");
 
            (bool success, ) = CONSOLIDATION_SYSTEM_CONTRACT.call{value: fee}(data[i]);
            require(success, "call failed");
        }
    }
 
}