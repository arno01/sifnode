// SPDX-License-Identifier: Apache-2.0
pragma solidity 0.8.0;

import "./CosmosBank.sol";
import "./EthereumWhitelist.sol";
import "./CosmosWhiteList.sol";
import "../Oracle.sol";
import "../CosmosBridge.sol";
import "./BankStorage.sol";
import "./Pausable.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title Bridge Bank
 * @dev Bank contract which coordinates asset-related functionality.
 *      CosmosBank manages the minting and burning of tokens which
 *      represent Cosmos based assets, while EthereumBank manages
 *      the locking and unlocking of Ethereum and ERC20 token assets
 *      based on Ethereum. WhiteList records the ERC20 token address 
 *      list that can be locked.
 */
contract BridgeBank is BankStorage,
    CosmosBank,
    EthereumWhiteList,
    CosmosWhiteList,
    Pausable {
    using SafeERC20 for IERC20;

    /**
     * @dev Has the contract been initialized?
     */
    bool private _initialized;

    /**
     * @dev Has the contract been reinitialized?
     */
    bool private _reinitialized;

    /**
     * @notice Initializer
     * @param _operator Manages the contract
     * @param _cosmosBridgeAddress The CosmosBridge contract's address
     * @param _owner Manages whitelists
     * @param _pauser Can pause the system
     * @param _networkDescriptor Indentifies the connected network
     */
    function initialize(
        address _operator,
        address _cosmosBridgeAddress,
        address _owner,
        address _pauser,
        uint256 _networkDescriptor
    ) public {
        require(!_initialized, "Init");

        CosmosWhiteList._cosmosWhitelistInitialize();
        EthereumWhiteList.initialize();

        contractName[address(0)] = "Ethereum";
        contractSymbol[address(0)] = "ETH";

        _initialized = true;

        _initialize(
            _operator,
            _cosmosBridgeAddress,
            _owner,
            _pauser,
            _networkDescriptor
        );
    }

    /**
     * @notice Allows the current operator to reinitialize the contract
     * @param _operator Manages the contract
     * @param _cosmosBridgeAddress The CosmosBridge contract's address
     * @param _owner Manages whitelists
     * @param _pauser Can pause the system
     * @param _networkDescriptor Indentifies the connected network
     */
    function reinitialize(
        address _operator,
        address _cosmosBridgeAddress,
        address _owner,
        address _pauser,
        uint256 _networkDescriptor
    ) public onlyOperator {
        require(!_reinitialized, "Already reinitialized");

        _reinitialized = true;

        _initialize(
            _operator,
            _cosmosBridgeAddress,
            _owner,
            _pauser,
            _networkDescriptor
        );
    }

    /**
     * @dev Internal function called by initialize() and reinitialize()
     * @param _operator Manages the contract
     * @param _cosmosBridgeAddress The CosmosBridge contract's address
     * @param _owner Manages whitelists
     * @param _pauser Can pause the system
     * @param _networkDescriptor Indentifies the connected network
     */
    function _initialize(
        address _operator,
        address _cosmosBridgeAddress,
        address _owner,
        address _pauser,
        uint256 _networkDescriptor
    ) private {
        Pausable._pausableInitialize(_pauser);

        operator = _operator;
        cosmosBridge = _cosmosBridgeAddress;
        owner = _owner;
        networkDescriptor = _networkDescriptor;
    }

    /**
     * @dev Modifier to restrict access to operator
     */
    modifier onlyOperator() {
        require(msg.sender == operator, "!operator");
        _;
    }

    /**
     * @dev Modifier to restrict access to owner
     */
    modifier onlyOwner {
        require(msg.sender == owner, "!owner");
        _;
    }

    /**
     * @dev Modifier to restrict access to the cosmos bridge
     */
    modifier onlyCosmosBridge {
        require(
            msg.sender == cosmosBridge,
            "!cosmosbridge"
        );
        _;
    }

    /**
     * @dev Modifier to only allow valid sif addresses
     */
    modifier validSifAddress(bytes calldata sifAddress) {
        require(verifySifAddress(sifAddress) == true, "INV_SIF_ADDR");
        _;
    }

    /**
     * @notice Allows the operator to add or remove `_token` to/from the Eth whitelist
     * @dev Set the token address in Eth whitelist
     * @param _token ERC 20's address
     * @param _inList Set the _token in list or not
     * @return New value of if _token in whitelist
     */
    function updateEthWhiteList(address _token, bool _inList)
        public
        onlyOperator
        returns (bool)
    {
        // Do not allow a token with the same address to be whitelisted
        if (_inList) {
            // if we want to add it to the whitelist, make sure it's not there yet
            require(!getTokenInEthWhiteList(_token), "already in eth whitelist");
            require(!getCosmosTokenInWhiteList(_token), "already in cosmos whitelist");
        } else {
            // if we want to de-whitelist it, make sure that the token is already whitelisted 
            require(getTokenInEthWhiteList(_token), "!whitelisted");
        }
        return setTokenInEthWhiteList(_token, _inList);
    }

    /**
     * @notice Allows the operator to add or remove a list of tokens to/from the Eth whitelist
     * @dev Set N token addresses in Eth whitelist
     * @param _tokens List of ERC 20's addresses
     * @param _inList List of booleans for each address: set _tokens[i] in list or not
     */
    function batchUpdateEthWhiteList(address[] calldata _tokens, bool[] calldata _inList) public onlyOperator {
        require(_tokens.length == _inList.length, "INV_LEN");

        for (uint256 i = 0; i < _tokens.length; i++) {
            updateEthWhiteList(_tokens[i], _inList[i]);
        }
    }

    /**
     * @dev Set the token address in Cosmos whitelist
     * @param token ERC 20's address
     * @param inList Set the token in list or not
     * @return New value of if token is in whitelist
     */
    function setTokenInCosmosWhiteList(address token, bool inList)
        internal returns (bool)
    {
        _cosmosTokenWhiteList[token] = inList;
        emit LogWhiteListUpdate(token, inList);
        return inList;
    }

    /**
     * @notice Transfers ownership of this contract to `newOwner`
     * @dev Cannot transfer ownership to the zero address
     * @param newOwner The new owner's address
     */
    function changeOwner(address newOwner) public onlyOwner {
        require(newOwner != address(0), "invalid address");
        owner = newOwner;
    }

    /**
     * @notice Transfers the operator role to `_newOperator`
     * @dev Cannot transfer role to the zero address
     * @param _newOperator: the new operator's address
     */
    function changeOperator(address _newOperator) public onlyOperator {
        require(_newOperator != address(0), "invalid address");
        operator = _newOperator;
    }

    /**
     * @dev Validates if a sif address has a correct prefix
     * @param sifAddress The Sif address to check
     * @return Boolean: does it have the correct prefix?
     */
    function verifySifPrefix(bytes calldata sifAddress) private pure returns (bool) {
        bytes3 sifInHex = 0x736966;

        for (uint256 i = 0; i < sifInHex.length; i++) {
            if (sifInHex[i] != sifAddress[i]) {
                return false;
            }
        }
        return true;
    }

    /**
     * @dev Validates if a sif address has the correct length
     * @param sifAddress The Sif address to check
     * @return Boolean: does it have the correct length?
     */
    function verifySifAddressLength(bytes calldata sifAddress) private pure returns (bool) {
        return sifAddress.length == 42;
    }

    /**
     * @dev Validates if a sif address has a correct prefix and the correct length
     * @param sifAddress The Sif address to be validated
     * @return Boolean: is it a valid Sif address?
     */
    function verifySifAddress(bytes calldata sifAddress) private pure returns (bool) {
        return verifySifAddressLength(sifAddress) && verifySifPrefix(sifAddress);
    }

    /**
     * @notice Validates whether `_sifAddress` is a valid Sif address
     * @dev Function used only for testing
     * @param _sifAddress Bytes representation of a Sif address
     * @return Boolean: is it a valid Sif address?
     */
    function VSA(bytes calldata _sifAddress) external pure returns (bool) {
        return verifySifAddress(_sifAddress);
    }

    /**
     * @notice CosmosBridge calls this function to create a new BridgeToken
     * @dev Only CosmosBridge can create a new BridgeToken 
     * @param name The new BridgeToken's name
     * @param symbol The new BridgeToken's symbol
     * @param decimals The new BridgeToken's decimals
     * @param cosmosDenom The new BridgeToken's denom
     * @return The new BridgeToken contract's address
     */
    function createNewBridgeToken(
        string calldata name,
        string calldata symbol,
        uint8 decimals,
        string calldata cosmosDenom
    ) external onlyCosmosBridge returns (address) {
        address newTokenAddress = deployNewBridgeToken(
            name,
            symbol,
            decimals,
            cosmosDenom
        );
        setTokenInCosmosWhiteList(newTokenAddress, true);
        contractDenom[newTokenAddress] = cosmosDenom;

        return newTokenAddress;
    }

    /**
     * @notice Allows the owner to add `contractAddress` as an existing BridgeToken
     * @dev Adds the token to Cosmos Whitelist
     * @param contractAddress The token's address
     * @return New value of if token is in whitelist (true)
     */
    function addExistingBridgeToken(
        address contractAddress    
    ) external onlyOwner returns (bool) {
        return setTokenInCosmosWhiteList(contractAddress, true);
    }

    /**
     * @notice CosmosBridge calls this function to mint or unlock tokens
     * @dev Controlled tokens will be minted, others will be unlocked
     * @param ethereumReceiver Tokens will be sent to this address
     * @param tokenAddress The BridgeToken's address
     * @param amount How much should be minted or unlocked 
     */
    function handleUnpeg(
        address payable ethereumReceiver,
        address tokenAddress,
        uint256 amount
    ) external onlyCosmosBridge whenNotPaused {
        // if this is a bridge controlled token, then we need to mint
        if (getCosmosTokenInWhiteList(tokenAddress)) {
            mintNewBridgeTokens(
                ethereumReceiver,
                tokenAddress,
                amount
            );
        } else {
            // if this is an external token, we unlock
            unlock(ethereumReceiver, tokenAddress, amount);
        }
    }

    /**
     * @notice Burns `amount` `token` tokens for `recipient`
     * @dev Burns BridgeTokens representing native Cosmos assets
     * @param recipient Bytes representation of destination address
     * @param token Token address in origin chain (0x0 if ethereum)
     * @param amount How much will be burned
     */
    function burn(
        bytes calldata recipient,
        address token,
        uint256 amount
    ) external validSifAddress(recipient) onlyCosmosTokenWhiteList(token) whenNotPaused {
        // revert if the token doesn't have a denom
        string memory denom = getDenom(token);
        require(keccak256(abi.encodePacked(denom)) != keccak256(abi.encodePacked("")), "INV_DENOM");

        // burn the tokens
        BridgeToken(token).burnFrom(msg.sender, amount);
        
        // decimals defaults to 18 if call to decimals fails
        uint8 decimals = getDecimals(token);

        lockBurnNonce = lockBurnNonce + 1;

        emit LogBurn(
            msg.sender,
            recipient,
            token,
            amount,
            lockBurnNonce,
            decimals,
            networkDescriptor,
            denom
        );
    }

    /**
     * @dev Fetches the name of a token by address
     * @param token The BridgeTokens's address
     * @return The bridgeTokens's name or ''
     */
    function getName(address token) private returns (string memory) {
        string memory name = contractName[token];

        // check to see if we already have this name stored in the smart contract
        if (keccak256(abi.encodePacked(name)) != keccak256(abi.encodePacked(""))) {
            return name;
        }

        try BridgeToken(token).name() returns (string memory _name) {
            name = _name;
            contractName[token] = _name;
        } catch {
            // if we can't access the name function of this token, return an empty string
            name = "";
        }

        return name;
    }

    /**
     * @dev Fetches the symbol of a token by address
     * @param token The bridgeTokens's address
     * @return The bridgeTokens's symbol or ''
     */
    function getSymbol(address token) private returns (string memory) {
        string memory symbol = contractSymbol[token];

        // check to see if we already have this name stored in the smart contract
        if (keccak256(abi.encodePacked(symbol)) != keccak256(abi.encodePacked(""))) {
            return symbol;
        }

        try BridgeToken(token).symbol() returns (string memory _symbol) {
            symbol = _symbol;
            contractSymbol[token] = _symbol;
        } catch {
            // if we can't access the symbol function of this token, return an empty string
            symbol = "";
        }

        return symbol;
    }

    /**
     * @dev Fetches de decimals of a token by address
     * @param token The bridgeTokens's address
     * @return The bridgeTokens's decimals or 18
     */
    function getDecimals(address token) private returns (uint8) {
        uint8 decimals = contractDecimals[token];
        if (decimals > 0) {
            return decimals;
        }

        try BridgeToken(token).decimals() returns (uint8 _decimals) {
            require(decimals < 100, "invalid decimals");
            decimals = _decimals;
            contractDecimals[token] = _decimals;
        } catch {
            // if we can't access the decimals function of this token,
            // assume that it has 18 decimals
            decimals = 18;
        }

        return decimals;
    }

    /**
     * @dev Fetches the denom of a token by address
     * @param token The bridgeTokens's address
     * @return The bridgeTokens's denom or ''
     */
    function getDenom(address token) private returns (string memory) {
        string memory denom = contractDenom[token];

        // check to see if we already have this denom stored in the smart contract
        if (keccak256(abi.encodePacked(denom)) != keccak256(abi.encodePacked(""))) {
            return denom;
        }

        try BridgeToken(token).cosmosDenom() returns (string memory _denom) {
            denom = _denom;
            contractDenom[token] = _denom;
        } catch {
            denom = "";
        }

        return denom;
    }

    /**
     * @notice Locks `amount` `token` tokens for `recipient`
     * @dev Locks received Ethereum/ERC20 funds
     * @param recipient Bytes representation of destination address
     * @param token Token address in origin chain (0x0 if ethereum)
     * @param amount Value of deposit
     */
    function lock(
        bytes calldata recipient,
        address token,
        uint256 amount
    ) external payable validSifAddress(recipient) whenNotPaused {
        if (token == address(0)) {
            return handleNativeCurrencyLock(recipient, amount);
        }
        require(msg.value == 0, "INV_NATIVE_SEND");

        lockBurnNonce += 1;
        _lockTokens(recipient, token, amount, lockBurnNonce);
    }

    /**
     * @notice Locks multiple tokens in the bridge contract in a single tx
     * @dev This is used to handle the case where the user is sending tokens
     * @param recipient Bytes representation of destination address
     * @param token Token address
     * @param amount Value of deposit
     */
    function multiLock(
        bytes[] calldata recipient,
        address[] calldata token,
        uint256[] calldata amount
    ) external whenNotPaused {
        require(recipient.length == token.length, "M_P");
        require(token.length == amount.length, "M_P");

        // use intermediate lock burn nonce to distinguish between different lock calls
        // this alows us to track the lock calls in the logs
        // and to prevent double locking
        // (i.e. if a user calls lock twice with the same amount, we don't want to lock twice)
        // This pattern of using the intermediate value will save us gas
        // by utilizing the stack for all intermediate values
        uint256 intermediateLockBurnNonce = lockBurnNonce;

        for (uint256 i = 0; i < recipient.length; i++) {
            intermediateLockBurnNonce++;

            _lockTokens(
                recipient[i],
                token[i],
                amount[i],
                intermediateLockBurnNonce
            );
        }
        lockBurnNonce = intermediateLockBurnNonce;
    }

    /**
     * @notice Locks or burns multiple tokens in the bridge contract in a single tx
     * @param recipient Bytes representation of destination address
     * @param token Token address
     * @param amount Value of deposit
     * @param isBurn Should the tokens be burned?
     */
    function multiLockBurn(
        bytes[] calldata recipient,
        address[] calldata token,
        uint256[] calldata amount,
        bool[] calldata isBurn
    ) external whenNotPaused {
        // all array inputs must be of the same length
        // else throw malformed params error
        require(recipient.length == token.length, "M_P");
        require(token.length == amount.length, "M_P");
        require(token.length == isBurn.length, "M_P");

        uint256 intermediateLockBurnNonce = lockBurnNonce;

        for (uint256 i = 0; i < recipient.length; i++) {
            intermediateLockBurnNonce++;

            if (isBurn[i]) {
                _burnTokens(
                    recipient[i],
                    token[i],
                    amount[i],
                    intermediateLockBurnNonce
                );
            } else {
                _lockTokens(
                    recipient[i],
                    token[i],
                    amount[i],
                    intermediateLockBurnNonce
                );
            }
        }
        lockBurnNonce = intermediateLockBurnNonce;
    }

    /**
     * @dev Locks a token in the bridge contract
     * @param recipient Bytes representation of destination address
     * @param tokenAddress Token address
     * @param tokenAmount Value of deposit
     * @param _lockBurnNonce A global nonce for locking an burning tokens
     */
    function _lockTokens(
        bytes calldata recipient,
        address tokenAddress,
        uint256 tokenAmount,
        uint256 _lockBurnNonce
    ) private onlyEthTokenWhiteList(tokenAddress) validSifAddress(recipient) {
        IERC20 tokenToTransfer = IERC20(tokenAddress);
        // lock tokens
        tokenToTransfer.safeTransferFrom(
            msg.sender,
            address(this),
            tokenAmount
        );

        // decimals defaults to 18 if call to decimals fails
        uint8 decimals = getDecimals(tokenAddress);

        // Get name and symbol
        string memory name = getName(tokenAddress);
        string memory symbol = getSymbol(tokenAddress);

        emit LogLock(
            msg.sender,
            recipient,
            tokenAddress,
            tokenAmount,
            _lockBurnNonce,
            decimals,
            symbol,
            name,
            networkDescriptor
        );
    }

    /**
     * @dev Burns a token
     * @param recipient Bytes representation of destination address
     * @param tokenAddress Token address
     * @param tokenAmount How much should be burned
     * @param _lockBurnNonce A global nonce for locking an burning tokens
     */
    function _burnTokens(
        bytes calldata recipient,
        address tokenAddress,
        uint256 tokenAmount,
        uint256 _lockBurnNonce
    ) private onlyCosmosTokenWhiteList(tokenAddress) validSifAddress(recipient) {
        BridgeToken tokenToTransfer = BridgeToken(tokenAddress);
        
        // burn tokens
        tokenToTransfer.burnFrom(msg.sender, tokenAmount);

        // revert if the token doesn't have a denom
        string memory denom = getDenom(tokenAddress);
        require(keccak256(abi.encodePacked(denom)) != keccak256(abi.encodePacked("")), "INV_DENOM");

        // decimals defaults to 18 if call to decimals fails
        uint8 decimals = getDecimals(tokenAddress);

        emit LogBurn(
            msg.sender,
            recipient,
            tokenAddress,
            tokenAmount,
            _lockBurnNonce,
            decimals,
            networkDescriptor,
            denom
        );
    }

    /**
     * @dev Locks received EVM native tokens
     * @param recipient Bytes representation of destination address
     * @param amount Value of deposit
     */
    function handleNativeCurrencyLock(
        bytes calldata recipient,
        uint256 amount
    ) internal validSifAddress(recipient) {
        require(msg.value == amount, "amount mismatch");

        address token = address(0);

        // decimals defaults to 18 if call to decimals fails
        uint8 decimals = 18;

        // Get name and symbol
        string memory name = getName(token);
        string memory symbol = getSymbol(token);

        lockBurnNonce = lockBurnNonce + 1;

        emit LogLock(
            msg.sender,
            recipient,
            token,
            amount,
            lockBurnNonce,
            decimals,
            symbol,
            name,
            networkDescriptor
        );
    }

    /**
     * @dev Unlocks native or ERC20 tokens
     * @param recipient Recipient's Ethereum address
     * @param token Token contract address
     * @param amount Wei amount or ERC20 token count
     */
    function unlock(
        address payable recipient,
        address token,
        uint256 amount
    ) internal {
        // Transfer funds to intended recipient
        if (token == address(0)) {
            (bool success,) = recipient.call{value: amount}("");
            require(success, "error sending ether");
        } else {
            IERC20 tokenToTransfer = IERC20(token);
            tokenToTransfer.safeTransfer(recipient, amount);
        }

        emit LogUnlock(recipient, token, amount);
    }

    function setBridgeTokenDenom(
      address _token, string memory _denom
    ) external onlyOperator returns (bool) {
      contractDenom[_token] = _denom;
      return BridgeToken(_token).setDenom(_denom);
    }
}
