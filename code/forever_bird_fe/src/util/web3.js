
var metamask = {
    
    test:function(){
        alert('aaa');
    },
    getWeb3:function() {
        let Web3 = require('web3');
        let web3;

        //当前存在web3实例
        if (typeof window.web3 !== 'undefined') {
            web3 = new Web3(window.web3.currentProvider)
            return web3;
        }
        return null
    },
    isWalletAvailable:function() {
        var web3 = this.getWeb3();
        if (web3 === null) {
            return { errNum: 1, obj: null }
        }

        if (web3.currentProvider.isMetaMask === true) {
            return { errNum: 0, obj: web3 }
        }
        else {
            return { errNum: 2, obj: web3 }
        }
    },
    getEthAccounts:function() {
        var result = this.isWalletAvailable();
        if (result.errNum !== 0) {
            return null
        }

        result.obj.eth.getAccounts((errNum, accounts) => {
            if (accounts.length === 0) {
                return null
            } else {
                return accounts
            }
        })
    },
    hasProvider:function(){
        window.addEventListener('load', () => {
            //是否存在钱包
            var result = this.isWalletAvailable();
            switch (result.error) {
                case 0:
                    result.obj.eth.getAccounts((error, accounts) => {
                        if (accounts.length == 0) {
                            console.log(accounts);
                        }
                        else {
                            console.log(accounts)
                        }
                    });
                    break;
                case 1:
                case 2:
                    console.log('no wallet available')
                    break;
            }
        });
    }
}

export default metamask;