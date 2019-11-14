import React, { Component } from 'react';
import metamask from '../util/web3';
import Header from '../containers/main/header';
import Menu from '../containers/main/menu';
import Footer from '../containers/main/foot';
import '../css/metamask.less';
import RegisterHeader from '../components/registerHeader';
import RegisterFooter from '../components/registerFooter';
import metanet from '../img/metanet.png';
import eth0 from '../img/eth0.png';
import {Redirect} from 'react-router-dom';

export default class MetaMask extends Component {

    constructor(props) {
        super(props);
        this.state = {
            token: null,
            isInstalled: false,
            islogin: false,
            netId : 1
        }
    }
    componentDidMount() {
        metamask.hasProvider();
        var web3 = metamask.getWeb3();
        if (web3 === null) {//是否安装metamask
            this.setState({
                token: null,
                isInstalled: false,
                islogin: false,
                netId : 1
            });
        } else {
            this.setState({
                token: null,
                isInstalled: true,
                islogin: false,
                netId : 1
            });
            web3.version.getNetwork((err, netId) => {
                switch (netId) {
                  case "1":
                    this.setState({
                        netId : 1
                    })
                    break
                  case "2":
                    console.log('This is the deprecated Morden test network.')
                    break
                  case "3":
                    this.setState({
                     netId : 3
                    })
                    break
                  case "4":
                    this.setState({
                        netId : 4
                    })
                    break
                  case "42":
                    this.setState({
                        netId : 42
                    })
                    break
                  default:
                    this.setState({
                        netId : 0
                    })
                }
              })
            var result = metamask.isWalletAvailable();
            if (result.errNum !== 0) {
                this.setState({
                    token: null,
                    isInstalled: true,
                    islogin: false,
                });
            }else{
                result.obj.eth.getAccounts((errNum, accounts) => {
                    if (accounts.length === 0) {
                        this.setState({
                            token: null,
                            isInstalled: true,
                            islogin: false,
                        });
                    } else {
                        this.setState({
                            token: accounts[0],
                            isInstalled: true,
                            islogin: true,
                        });
                    }
                })
            }
        }
        setInterval(() => {
            this.update();
        }, 1000);
    }

    update() {
        var web3 = metamask.getWeb3();
        if (web3 === null) {//是否安装metamask
            this.setState({
                token: null,
                isInstalled: false,
                islogin: false,
                netId : 1
            });
        } else {
            web3.version.getNetwork((err, netId) => {
                switch (netId) {
                  case "1":
                    this.setState({
                        netId : 1
                    })
                    break
                  case "2":
                    console.log('This is the deprecated Morden test network.')
                    break
                  case "3":
                    this.setState({
                     netId : 3
                    })
                    break
                  case "4":
                    this.setState({
                        netId : 4
                    })
                    break
                  case "42":
                    this.setState({
                        netId : 42
                    })
                    break
                  default:
                    this.setState({
                        netId : 0
                    })
                }
              })
            var result = metamask.isWalletAvailable();
            if (result.errNum !== 0) {
                this.setState({
                    token: null,
                    isInstalled: true,
                    islogin: false,
                });
            }else{
                result.obj.eth.getAccounts((errNum, accounts) => {
                    if (accounts.length === 0) {
                        this.setState({
                            token: null,
                            isInstalled: true,
                            islogin: false,
                        });
                    } else {
                        this.setState({
                            token: accounts[0],
                            isInstalled: true,
                            islogin: true,
                        });
                    }
                })
            }
        }
        // console.log('state:'+JSON.stringify(this.state));
    }

    render() {
        const { token, isInstalled, islogin, netId } = this.state;
        if (!this.state.isInstalled) {
            return (
                <div className="metamask-main">
                    <RegisterHeader />
                    <div className="sub-title">进入Forever Bird游戏</div>
                    <div className="sub-eng-title">Please unlock your MetaMask wallet or open this page in Toshi or Trust Wallet.</div>
                    <div className="plugin-wrap">
                        <div className="plugin-img">
                        
                        </div>
                        <div className="plugin-desc">
                            <div className="plugin-title-desc">Fishbank游戏需要火狐浏览器或谷歌Chrome安装MetaMask插件。</div>
                            <a className="plugin-btn" href="https://metamask.io/">下载MetaMask</a>
                            <div className="plugin-title-desc second-title">You can play Fishbank on any IOS or Android device with Trust Wallet or Toshi digital mobile Ethereum wallets.</div>
                            <div className="plugin-btn-wrap">
                                <a className="plugin-gray-btn" href="https://metamask.io/">Get Toshi</a>
                                <a className="plugin-gray-btn" href="https://metamask.io/">Trust Wallet</a>
                            </div>
                        </div>
                    </div>
                    <RegisterFooter />
                </div>
            )
        } else if (!islogin /* || netId !== 4 */) {
            return (
                <div className="metamask-main">
                    <RegisterHeader />
                    <div  className="cc" id="page">
                        <div  className="netting _mt50 _mb50">
                            <div  className="netting_12 _mb25 prefix_2">
                                <h1 className="cf">欢迎来到Forever Bird!</h1>
                                <p className="_tb6">
                                    开始游戏前还需进行一小步操作。
                                </p>
                            </div>
                            <div className="netting_4 prefix_2">
                                <p>
                                    请如右边所示切换你的MetaMask到Rinkeby网络。
                                </p>
                            </div>
                            <div  className="netting_4 prefix_2" style={{paddingLeft:25}}>
                                <img  className="_tc _mb15 shadow_6" src={metanet} width="100%" />
                            </div>
                        </div>
                    </div>
                    <RegisterFooter />
                </div>
            )
        } else {
            return (
                <Redirect to = {{
                        pathname : '/register',
                        state : {token :this.state.token}
                    }}
                />
            )
        }
    }
}