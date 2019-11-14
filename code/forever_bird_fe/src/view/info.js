import React, { Component } from 'react';
import {Link} from 'react-router-dom';
import MyHeader from '../components/myheader';
import RegisterFooter from '../components/registerFooter';
import Menu from '../containers/main/menu';
import {Table, List, Pagination } from 'antd';
import smallinfo from '../img/smallinfo.svg';
import '../css/info.less';
import {BASE_URL} from '../util/common';
import 'whatwg-fetch';
import {CONTRACT_HASH,ABI} from '../util/contract';
import moment from 'moment';

export default class Info extends Component {

    constructor(props) {
        super(props);
        this.state = {
            token: this.props.location.state.token,
            info : {},
            attackList : {
                count : 0,
                start : 0,
                total : 0,
                list : []
            },
        };
    }

    onPageChange = (page,pageSize) => {
        this.getAttackList(page,pageSize);
    }

    componentDidMount() {
        let data = {
            method : 'POST',
            headers: {
                'Content-Type': 'text/plain'
            },
            body : JSON.stringify({
                token_id : parseInt(this.props.match.params.id),
            }), 
            credentials: 'same-origin',
        }
        console.log('info');console.log(data);
        fetch(BASE_URL+'/api/bird/info',data)
        .then(response => {
            return response.json()
        }).then(json => {
            if(json.status === 0){
                this.setState(
                    {info:json.data}
                );
            }
        });
        this.getAttackList(1,10);
    }

    getAttackList = (page,pageSize) => {
        let data = {
            method : 'POST',
            headers: {
                'Content-Type': 'text/plain'
            },
            body : JSON.stringify({
                token_id : parseInt(this.props.match.params.id),
                start_page : page,
                page_size : pageSize
            }), 
            credentials: 'same-origin',
        }
        console.log('pk');console.log(data);
        fetch(BASE_URL+'/api/record/pk',data)
        .then(response => {
            return response.json()
        }).then(json => {
            if(json.status === 0){
                this.setState({
                    attackList:{
                        list : json.data.birds,
                        total : json.data.total_number,
                        current : page
                    }
                });
            }
        });
    }

    buy = (e) => {
        e.preventDefault();
        let data = {
          method : 'POST',
          headers: {
              'Content-Type': 'text/plain'
          },
          body : JSON.stringify({
              token_id : this.state.info.token_id
          }), 
          credentials: 'same-origin',
        }
        fetch(BASE_URL+'/api/buy/bird',data)
        .then(response => {
            return response.json()
        }).then(json => {
            if(json.status === 0){
              var price = json.data.price;
              var fee = json.data.fee;
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
              var priceValue =  parseInt(price) + parseInt(fee);
              contractInstance.trade.sendTransaction(
                  this.state.info.token_id,price,fee,sign,{
                  from:this.props.location.state.token,value:priceValue},(err,result) => {
                  if(!err){
                    let data = {
                      method : 'POST',
                      headers: {
                        'Content-Type': 'text/plain'
                      },
                      credentials: 'same-origin',
                      body : JSON.stringify({
                        hash : result,
                        token_id : this.props.info.token_id
                      }),
                  }
                  fetch(BASE_URL+'/api/buy/bird/hash',data)
                  .then(response => {
                    return response.json()
                  }).then(json => {
                    if(json.status === 0){
                      alert('购买成功，上链中，请稍后');
                    }
                  }) 
                  }else{
                    alert('交易失败,请重试.失败详情:'+result);
                  }
              });
            }else{
              alert(json.message);
            }
        }) 
      }

    render() {
        const pagination = {
            pageSize: 10,
            current: this.state.attackList.current,
            total: this.state.attackList.total,
            onChange: this.onPageChange,
        };
        const column = [
            {
                title : '日期',
                dataIndex : 'day',
                key : 'day',
                colSpan : 1,
                align : 'center',
                render : (text,record) => {
                    return {
                        children :<span>{moment(text*1000).format('YYYY-MM-DD HH:mm:ss')}</span>
                    }
                }
            },
            {
                title : '挑战者',
                dataIndex : 'challenger_name',
                key : 'challenger_name',
                colSpan : 3,
                align : 'center',
                render : (text,record) => {
                    const result = {
                        pathname : '/result/'+record.hash,
                        state : {token: this.props.location.state.token}
                    }
                    return {
                        children : <span><Link to={result}>{text}</Link></span>,
                        props : {
                            colSpan : 3
                        }
                    }
                }
            },
            {
                title : '被挑战者',
                dataIndex : 'resister_name',
                key : 'resister_name',
                colSpan : 3,
                align : 'center',
                render : (text,record) => {
                    const result = {
                        pathname : '/result/'+record.hash,
                        state : {token: this.props.location.state.token}
                    }
                    return {
                        children : <span><Link to={result}>{text}</Link></span>,
                        props : {
                            colSpan : 3
                        }
                    }
                }
            },
            {
                title : '结果',
                dataIndex : 'isWin',
                key : 'isWin',
                colSpan : 3,
                align : 'center',
                className: 'bird-res-tr',
                render : (text,record) => {
                    var isWin = record.is_win;
                    let color = 'incolor1';
                    let result = '失败';
                    let fh = '-';
                    if(record.challenger_id !== parseInt(this.props.match.params.id)) {
                        isWin = !isWin;
                    };
                    if(isWin){
                        color = 'incolor2';
                        result = '胜利';
                        fh = '+';
                    };
                    return {
                        children : <div className={`${color} bird-res-color`}><span className="attack-res-in-table">{result}{fh}{record.winner_reward_coin}盎司</span></div>
                    }
                }
            }
        ];
        return (
            <div className="main">
                <MyHeader />
                <Menu {...this.props}/>
                <div className='inpool'>
                    <div className="netting">
                        <div className="netting bird-info-wrap">
                            <div className="bird-simple-info-wrap">
                                {/* <div className="bird-token">Crypto bird token #{this.state.info.token_id}</div> */}
                                <h2 className="bird-name" >{this.state.info.name || '无名鸟'}</h2>
                                <div className="bird-rarity" >{this.state.info.rarity || '普通'}</div>
                            </div>
                            <div  className="inimg bird-img-info">
                                <img src={BASE_URL+this.state.info.svg_path} className="fish-img common medusa level-1 inlevel" />
                            </div>
                            <div className="bird-price-info">
                                <div className="bird-level">等级 <b >{this.state.info.level || 0}</b></div>
                                <div className="bird-weight" ><b >{this.state.info.weight || 0}</b> 盎司</div>
                            </div>
                        </div>
                        <div className="bird-own">所有者:{this.state.info.owner}</div>
                    </div>
                </div>
                <div className="bird-price-wrap">
                    <div className="bird-price">售价:{this.state.info.auction_data ? this.state.info.auction_data.price + 'ETH' : '无'} </div>
                    {/* <a className="buy-bird-btn" onClick={this.buy}>购买</a> */}
                </div>
                <div className="bird-info-detail" >
                        <div className="netting">
                            <div className="bird-detail-table">
                                <h4 className="title">属性特性</h4>

                                <div className="detail-item">
                                    <div className="detail-key">力量</div>
                                    <div className="detail-value">{this.state.info.power}</div>
                                </div>
                                <div className="detail-item">
                                    <div className="detail-key">速度</div>
                                    <div className="detail-value">{this.state.info.speed}</div>
                                </div>
                                <div className="detail-item last">
                                    <div className="detail-key">经验</div>
                                    <div className="detail-value">{this.state.info.exp}</div>
                                </div>
                                <p className="birth-day">鸟鸟生日：{moment(this.state.info.birth_date*1000).format('YYYY-MM-DD HH:mm:ss')}</p>
                            </div>


                            <div className="block_a attack-history">
                                <h4 className="attack-title">战斗历史</h4>
                                <div className="_p20">
                                    <Table dataSource={this.state.attackList.list} columns={column} rowKey='hash' pagination={pagination}/>
                                </div>
                            </div>
                        </div>
                </div>
                <RegisterFooter />
            </div>
        )
    }
}