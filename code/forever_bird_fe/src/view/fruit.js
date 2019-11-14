import React,{Component} from 'react';
import {Button} from 'antd';
 
import fruit2 from '../img/csg-02.png';
import fruit3 from '../img/csg-03.png';
import fruit4 from '../img/csg-04.png';
import '../css/fruit.less';
import 'whatwg-fetch';
import {BASE_URL} from '../util/common';
import {CONTRACT_HASH,ABI} from '../util/contract';

export default class Fruit extends Component{

    constructor(props){
        super(props);
        this.state = {
            list : [
                {id:1,price:0},
                {id:2,price:0},
                {id:3,price:0}
            ]
        }
    }

    componentDidMount(){
        let data = {
            method : 'POST',
            headers: {
                'Content-Type': 'text/plain'
            },
            credentials: 'same-origin',
        }
        fetch(BASE_URL+'/api/fruit/list',data)
        .then(response => {
            return response.json()
        }).then(json => {
            let Web3 = require('web3');
            let web3;
            if (typeof window.web3 !== 'undefined') {
                web3 = new Web3(window.web3.currentProvider)
            }
            this.setState({
                list : [
                    {id:json.data[0].id,price:web3.fromWei(json.data[0].price,'ether')},
                    {id:json.data[1].id,price:web3.fromWei(json.data[1].price,'ether')},
                    {id:json.data[2].id,price:web3.fromWei(json.data[2].price,'ether')}
                ]
            });
        })
    }

    buy = (type) => {
        let data = {
            method : 'POST',
            headers: {
                'Content-Type': 'text/plain'
            },
            body : JSON.stringify({
                type : type,
                token_id : this.props.token
            }),
            credentials: 'same-origin',
        }
        fetch(BASE_URL+'/api/buy/fruit',data)
        .then(response => {
            return response.json()
        }).then(json => {
            console.log(json);
            if(json.status === 0) {
                var price = json.data.price;
                var sign = json.data.sign;
                let Web3 = require('web3');
                let web3;
                if (typeof window.web3 !== 'undefined') {
                    console.log('no web3');
                    web3 = new Web3(window.web3.currentProvider)
                }
                var abiArray = ABI;
                var MyContract = web3.eth.contract(abiArray);
                var contractInstance = MyContract.at(CONTRACT_HASH);
                console.log(this.props.token,type,sign,{
                    from: this.props.from, value: price});
                contractInstance.feedFruit.sendTransaction(
                    this.props.token,type,sign,{
                    from: this.props.from, value: price
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
                                token_id: this.props.token,
                                type : type
                            }),
                            credentials: 'same-origin'
                        }
                        fetch(BASE_URL + '/api/buy/fruit/hash', data)
                        .then(response => {
                            return response.json()
                        }).then(json => {
                            if (json.status === 0) {
                                alert('eat fruit success');
                                this.props.handleVisibleChange(false,true);
                            } else {
                                alert('eat fruit fail.');
                            }
                        })
                    }else{
                        alert('eat fruit fail.err msg is.'+err);
                    }
                });
            }else{
                alert('err.'+json.message);
                this.props.handleVisibleChange(false,true);
            }
        })
    }

    buyWeight = () => {
        this.buy(1);
    }

    buySpeed = () => {
        this.buy(2);
    }

    buyExp = () => {
        this.buy(3);
    }

    render(){
        return(
            <div className="fruit-main">
                <div className='div fruit-wrap'>
                    <div className='imgdiv'>
                        <img className="fruit-img" src={fruit2}/>
                    </div>
                    <div className='contentdiv'>
                        <div >
                            <p>力量果实，提升5%力量，价格：{this.state.list[0].price}eth</p>
                        </div>
                        <div >
                            <Button type="primary" className="fruit-btn" onClick={this.buyWeight}>购买</Button>
                        </div>
                    </div>
                </div>
                <div className='div fruit-wrap'>
                    <div className='imgdiv'>
                        <img className="fruit-img" src={fruit3}/>
                    </div>
                    <div className='contentdiv'>
                        <div >
                            <p>速度果实，提升5%速度，价格：{this.state.list[1].price}eth</p>
                        </div>
                        <div >
                            <Button type="primary" className="fruit-btn" onClick={this.buySpeed}>购买</Button>
                        </div>
                    </div>
                </div>
                <div className='div fruit-wrap'>
                    <div className='imgdiv'>
                        <img className="fruit-img" src={fruit4}/>
                    </div>
                    <div className='contentdiv'>
                        <div >
                            <p>经验果实，提升5%经验，价格：{this.state.list[2].price}eth</p>
                        </div>
                        <div >
                            <Button type="primary" className="fruit-btn" onClick={this.buyExp}>购买</Button>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}