import React, { Component } from 'react';
import MyHeader from '../components/myheader';
import RegisterFooter from '../components/registerFooter';
import Menu from '../containers/main/menu';
import { Redirect } from 'react-router-dom';
import { Button } from 'antd';
 import '../css/catchbird.less';
import { BASE_URL } from '../util/common';
import 'whatwg-fetch';
import {CONTRACT_HASH,ABI} from '../util/contract';

export default class CatchBird extends Component {

    constructor(props) {
        super(props);
        this.state = {
            token: this.props.location.state.token,
            isRedirect : false
        };
    }

    componentDidMount() {
    }

    catch = (e) => {
        e.preventDefault();
        let data = {
            method: 'POST',
            headers: {
                'Content-Type': 'text/plain'
            },
            credentials: 'same-origin'
        }
        fetch(BASE_URL + '/api/me', data)
        .then(response => {
            return response.json()
        }).then(json => {
            if (json.status === 0) {
                let Web3 = require('web3');
                let web3;
                if (typeof window.web3 !== 'undefined') {
                    console.log('no web3');
                    web3 = new Web3(window.web3.currentProvider)
                }
                var abiArray = ABI;
                var MyContract = web3.eth.contract(abiArray);
                var contractInstance = MyContract.at(CONTRACT_HASH);
                contractInstance.catchBird.sendTransaction({
                    from: this.props.location.state.token, value: web3.toWei('0.03', 'ether')
                }, (err, result) => {
                    console.log(err)
                    console.log(result)
                    if (!err) {
                        let data = {
                            method: 'POST',
                            headers: {
                                'Content-Type': 'text/plain'
                            },
                            body: JSON.stringify({
                                hash: result,
                                address: this.props.location.state.token
                            }),
                            credentials: 'same-origin'
                        }
                        fetch(BASE_URL + '/api/catch/hash', data)
                        .then(response => {
                            return response.json()
                        }).then(json => {
                            if (json.status === 0) {
                                alert('抓鸟成功,上链中,请稍后...');
                                this.setState({
                                    isRedirect: true
                                });
                            } else {
                                alert('catch bird fail.');
                            }
                        })
                    }else{
                        alert('catch bird fail.err msg is.'+err);
                    }
                });
            } else {
                const { history } = this.props;
                history.push('/metamask');
            }
        })
    }

    render() {
        if (this.state.isRedirect) {
            return <Redirect
              to={{
                pathname: '/mybird',
                state: { token: this.props.location.state.token }
              }}
            />
        }else{
            return (
                <div className="catch-main">
                    <MyHeader />
                    <Menu {...this.props} />
                    <div className="catch-main-wrap">
                        <div className="catch" >
                            <p className="catch-logo"></p>
                            <div className="catch-btn" onClick={this.catch}>去抓鸟吧</div>
                            <div className="catch-price">0.03ETH</div>
                        </div>
                        <div className="catch-desc-title">捉鸟说明</div>
                        <div className="catch-desc-price">每次捉鸟需要花费0.03ETH</div>
                        <div className="catch-desc">每次捉鸟捕获加密鸟至少获得一条普通或更好的鸟</div>
                    </div>
                    <RegisterFooter />
                </div>
            )
        }
    }
}