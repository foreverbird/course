import React, { Component } from 'react';
import Header from './header';
import Body from './body';
import Footer from './foot';
import home from '../../img/home.png';
import {Button} from 'antd';
import {Link} from 'react-router-dom';
import RegisterFooter from '../../components/registerFooter';
import './index.less';
//import Web3 from 'web3'

export default class MainFrame extends Component {

    state = {
        status: -1,
        token: 'ls'
    }


    constructor(props) {
        super(props)
    }

    getWeb3() {
        let Web3 = require('web3');
        let web3;

        //当前存在web3实例
        if (typeof window.web3 !== 'undefined') {
            console.log('no web3');
            web3 = new Web3(window.web3.currentProvider)
            return web3;
        }

        return null
    }

    isWalletAvailable() {
        var web3 = this.getWeb3();
        if (web3 === null) {
            return { error: 1, obj: null }
        }

        if (web3.currentProvider.isMetaMask === true) {
            return { error: 0, obj: web3 }
        }
        else {
            return { error: 2, obj: web3 }
        }
    }

    getEthAccounts(web3) {
        var result = this.isWalletAvailable();
        if (result.error !== 0) {
            return null
        }

        result.obj.eth.getAccounts((error, accounts) => {
            if (accounts.length == 0) {
                return null
            } else {
                return accounts
            }
        })
    }

    transaction(etherNumber) {
        //获取以太坊web3
        var web3 = this.getWeb3();
        if (web3 === null) {
            console.log('transaction web3 error');
            return null
        }

        //获取当前地址
        var accounts = this.getEthAccounts(web3);
        if (accounts === null) {
            console.log('transaction accounts error');
            return null
        }

        //获取gas
        //var gasPrice = web3.eth.gasPrice;
        //var gasVal = web3.eth.gasVal;

        var ethVal = web3.toWei(1, 'ether')

        var ethHex = '0x' + ethVal.toString(16);
        //var gasHex = '0x' + gasVal.toString(16);
        //var gpHex = '0x' + gasPrice.toString(16);


        web3.eth.sendTransaction({ from: '0x39E50914c7683bFC9431E2Fc59f431Dc58bBc2c5', to: '0xd24FD280A061c69032b0e87D936BB053eb279C50', value: ethVal }, function (err, transactionHash) {
            console.log('transactionHash:', err);
            if (!err)
                console.log('transactionHash:', transactionHash);
        });
    }

    hasProvider() {
        window.addEventListener('load', () => {
            //是否存在钱包
            var connect = this.getWeb3().isConnected();
            console.log('connect:'+connect);
            var result = this.isWalletAvailable();
            switch (result.error) {
                case 0:
                    result.obj.eth.getAccounts((error, accounts) => {
                        if (accounts.length == 0) {
                            this.setState(
                                {
                                    status: 3,
                                    token: 'null'
                                }
                            );
                        }
                        else {
                            console.log(accounts)
                            this.setState(
                                {
                                    status: 0,
                                    token: accounts
                                }
                            );
                            return
                        }
                    });

                    break;
                case 1:
                case 2:
                    console.log('no wallet available')
                    this.setState(
                        {
                            status: result.error,
                            token: 'null'
                        }
                    );
                    break;
            }
        });
    }

    updateWallet() {
        let Web3 = require('web3');
        if (typeof window.web3 !== 'undefined') {
            var web3 = new Web3(window.web3.currentProvider)
            web3.net.getListening((error, result) => {
                console.log("listener:" + result)
            })
        }

    }

    componentDidMount() {
        // this.hasProvider()
        // this.transaction(1)
        /*
                setInterval(() => {
                    this.updateWallet();
                }, 1000);
        */
    }

    render() {
        return (
            <div className="main">
                <div className="main-header">
                    <div className="home-logo"></div>
                    <div className="home-discorvery">发现</div>
                    <div className="home-intro">玩法</div>
                    <div className="home-qa">常见问题解答</div>
                    <div className="home-login begin-btn">
                        <Link to={{
                                pathname: '/metamask',
                            }}>登陆游戏
                        </Link>
                    </div>
                </div>
                <div className="main-top">
                    <div className="bird-intro">
                        <div className="wing-logo"></div>
                        <div className="bird-logo"></div>
                        <div className="begin-btn">
                        <Link to={{
                                pathname: '/metamask',
                            }}>开始游戏
                        </Link>
                        
                        </div>
                        <p className="index-title">全新区块链游戏，养成以及交易加密鸟代币</p>
                        <p className="index-sub-title">基于以太网智能合约的大型多人对战游戏</p>
                        <div className="buy-bird">
                            <div className="buy-bird-img"></div>
                            <div className="bird-num-wrap">
                              <div className="bird-row">
                                <div className="key">力量</div>
                                <div className="value">11</div>
                              </div>
                              <div className="bird-row">
                                <div className="key">速度</div>
                                <div className="value">123</div>
                              </div>
                              <div className="bird-row">
                                <div className="key">经验</div>
                                <div className="value">123</div>
                              </div>
                            </div>
                        </div>
                    </div>
                    <div className="margin-buffer">可收藏，可繁殖，可讨人喜欢</div>
                    <div className="bird-qukuai">
                        <div className="qukuai-logo"></div>
                        <div className="qukuai-title">每只鸟都是存放在以太坊区块链上的加密代币</div>
                        <div className="qukuai-sub-title">数字资产100%的玩家所拥有。他可以像普通的加密货币一样管理，转移或出售给任何其他玩家。它不能被销毁，删除或更换。</div>
                    </div>
                    <div className="bird-qa">
                        <div className="qa">
                            <p className="title">Forever Bird是什么？</p>
                            <p className="ctx">点对点（P2P）玩家对玩家（PVP）的游戏，通过在以太坊区块链上运行的智能合约，通过社区驱动的经济和不可变的游戏资产来增长，对抗和交易独特的数字鸟。</p>
                            <div className="qa-btn">常见问题解答</div>
                        </div>
                        <div className="logo"></div>
                    </div>
                    <div className="bird-tech">
                        <div className="qa">
                            <p className="title">该如何玩</p>
                            <p className="ctx">玩foreverbird的要求：Chrome或Firefox浏览器安装MetaMask扩展。</p>
                            <div className="qa-btn">更多阅读</div>
                        </div>
                        <div className="logo"></div>
                    </div>
                    <div className="bird-fight">
                        <div className="title">与区块链上的加密鸟展开战斗</div>
                        <div className="info">
                            <div className="info-bird"></div>
                            <div className="fork"></div>
                            <div className="small-bird"></div>
                            <div className="equal"></div>
                            <div className="big-bird"></div>
                        </div>
                        <div className="sub-title">攻击其他鸟，以增加你的鸟重量，并达到全球排行榜的顶部。</div>
                    </div>
                    <div className="bird-video">
                        <div className="title">视频教程</div>
                        <div className="video-wrap">
                            <div className="big-video"></div>
                            <div className="small-video-wrap">
                                <div className="sub-title">解释视频</div>
                                <div className="small-videos">
                                    <div className="video"></div>
                                    <div className="video"></div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <RegisterFooter />
            </div>
        );
    }
}