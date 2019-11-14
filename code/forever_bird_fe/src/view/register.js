import React, { Component } from 'react';
import '../css/register.less';
import RegisterHeader from '../components/registerHeader';
import RegisterFooter from '../components/registerFooter';
import { Form } from 'antd';
import metamask from '../util/web3';
import { BASE_URL } from '../util/common';
import { Redirect } from 'react-router-dom';
import 'whatwg-fetch';
import RegistrationForm from './regeistForm';

export default class Register extends Component {

    constructor(props) {
        super(props);
        this.state = {
            token: this.props.location.state.token,
            nickName: null
        }
    }

    componentDidMount() {
        let data = {
            method: 'POST',
            credentials: 'same-origin',
            // mode : 'cors',
            headers: {
                'Content-Type': 'text/plain'
            },
            body: JSON.stringify({
                address: this.state.token
            }),
        }
        fetch(BASE_URL + '/api/user/exist', data)
            .then(response => {
                return response.json()
            }).then(json => {
                if (json.status === 0 && json.data !== null) {
                    let nick = json.data.nick;
                    let email = json.data.email;
                    this.setState({
                        nick : nick,
                        email : email
                    });
                    let Web3 = require('web3');
                    let web3;
                    //当前存在web3实例
                    if (typeof window.web3 !== 'undefined') {
                        web3 = new Web3(window.web3.currentProvider)
                    }
                    let from = this.state.token;
                    var ethUtil = require('ethereumjs-util');
                    var msg = ethUtil.bufferToHex(new Buffer('forever bird'),'utf8');
                    var params = [msg, from];
                    var method = 'personal_sign';
                    web3.currentProvider.sendAsync({
                        method,
                        params,
                        from},(err,result) => {
                            if(result.result){
                                let data = {
                                    method: 'POST',
                                    headers: {
                                        'Content-Type': 'text/plain'
                                    },
                                    body: JSON.stringify({
                                        address: from,
                                        sign : result.result
                                    }),
                                    credentials: 'same-origin',
                                }
                                fetch(BASE_URL + '/api/login', data)
                                .then(response => {
                                    return response.json()
                                }).then(json => {
                                    if (json.status === 0) {
                                        this.setState({
                                            nickName: nick
                                        });
                                    }
                                })
                            }
                    });
                }
            })
            setInterval(() => {
                this.update();
            }, 1000);
    }

    update() {
        var web3 = metamask.getWeb3();
        if (web3 === null) {//是否安装metamask
            this.setState({
                token: null,
                netId : 1,
                isRedirect : true
            });
        } else {
            web3.version.getNetwork((err, netId) => {
                /* if(netId !== 4){
                    this.setState({
                        token : this.state.token,
                        netId : netId,
                        isRedirect : true
                    });
                }  */
            })
            var result = metamask.isWalletAvailable();
            if (result.errNum !== 0) {
                this.setState({
                    token: null,
                    isRedirect : true
                });
            }else{
                result.obj.eth.getAccounts((errNum, accounts) => {
                    if (accounts.length === 0) {
                        this.setState({
                            token: null,
                            isRedirect : true
                        });
                    } else if(accounts[0] !== this.state.token) {
                        this.setState({
                            token: accounts[0],
                            isRedirect : true
                        });
                    }
                })
            }
        }
    }

    render() {
        const { token, nickName } = this.state;
        if (nickName) {
            return (
                <Redirect to={{
                    pathname: '/mybird',
                    state: {
                        token: this.state.token,
                        nickName: this.state.nickName
                    }
                }}
                />
            )
        } else if(this.state.isRedirect){
            return (
                <Redirect to = {{
                        pathname : '/metamask',
                        state : {token :this.state.token}
                    }}
                />
            )

        } else {
            const WrappedRegistrationForm = Form.create()(RegistrationForm);
            return (
                <div className="register-main">
                    <RegisterHeader subTitle="请先注册" />
                    <div className='regeist'>
                        <WrappedRegistrationForm token={this.state.token} nick={this.state.nick} email={this.state.email} />
                    </div>
                    <RegisterFooter />
                </div>
            )
        }
    }
}