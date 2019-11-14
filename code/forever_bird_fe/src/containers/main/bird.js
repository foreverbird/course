import React, { Component } from 'react';
import { Redirect } from 'react-router-dom';
import '../../css/bird.less';
import small from '../../img/small.svg';
import info from '../../img/info.png';
import {BASE_URL,TOKEN} from '../../util/common';
import {Popover,Button,InputNumber  } from 'antd';
import Fruit from '../../view/fruit';
import 'whatwg-fetch';
import {CONTRACT_HASH,ABI} from '../../util/contract';
import AttackModal from '../../view/attackmodal';
import NoBird from '../../view/nobird';

export default class Bird extends Component {

  constructor(props) {
    super(props);
    this.state = {
      token : this.props.location.state.token,
      // token : TOKEN,
      toInfo: false,
      toBuy: false,
      price : 0.01,
      visible: false,
      modalVisible : false,
    };
  }

  info = (e) => {
    e.preventDefault();
    this.setState({
      toInfo: true
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
          token_id : this.props.info.token_id
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
              this.props.info.token_id,price,fee,sign,{
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
                  this.props.refresh(1);
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

  sell = () => {
    var price = this.state.price;
    console.log('price.'+price);
    if(!price){
      alert('价格设置错误.');
      return;
    }
    let Web3 = require('web3');
    let web3;
    if (typeof window.web3 !== 'undefined') {
      console.log('no web3');
      web3 = new Web3(window.web3.currentProvider)
    }
    price = web3.toWei(price,'ether');
    let data = {
      method : 'POST',
      headers: {
          'Content-Type': 'text/plain'
      },
      body : JSON.stringify({
          price : price,
          token_id : this.props.info.token_id
      }), 
      credentials: 'same-origin',
    }
    fetch(BASE_URL+'/api/sell/bird',data)
    .then(response => {
        return response.json()
    }).then(json => {
        if(json.status === 0){
          this.props.refresh(1);
          // this.setState({
          //   refresh : 1
          // });
        }
    }) 
  }

  cancleSell = () => {
    let data = {
      method : 'POST',
      headers: {
          'Content-Type': 'text/plain'
      },
      body : JSON.stringify({
          token_id : this.props.info.token_id
      }), 
      credentials: 'same-origin',
    }
    fetch(BASE_URL+'/api/sell/off',data)
    .then(response => {
        return response.json()
    }).then(json => {
        if(json.status === 0){
          this.props.refresh(1);
          // this.setState({
          //   refresh : 1
          // });
        }
    }) 
  }

  onChange = (value) => {
    this.setState({
      'price' : value
    });
  }

  handleVisibleChange = (visible,refresh) => {
    console.log('handleVisibleChange')
    console.log(visible)
    console.log(refresh)
    this.setState({ visible : visible });
    if(!visible && refresh){
      this.props.refresh(1);
    }
  }

  attack = () => {
    var owener = this.props.location.state.token;
    let data = {
      method : 'POST',
      headers: {
          'Content-Type': 'text/plain'
      },
      body : JSON.stringify({
          address : owener,
          start_page : 1,
          page_size : 10000
      }), 
      credentials: 'same-origin',
    }
    fetch(BASE_URL+'/api/mybirds/pk',data)
    .then(response => {
      return response.json()
    }).then(json => {
      if(json.status === 0){
        this.props.showModal(json.data.total_number > 0 ? true : false,json.data.birds,this.props.info.token_id);
          this.setState({
            modalVisible : true,
            hasBird : json.data.total_number > 0 ? true : false,
            birdList : json.data.birds,
            attackId : this.props.info.token_id
          });
      }
    }) 
  }

  statusRender(){
    let status = this.props.info.status;
    let isSelf = this.props.isSelf;
    if(isSelf === 1){
      if(status === 0){
        return (
          <div>
            <a className="button" >初始化中...</a>
          </div>
        )
      }else if(status === 2){
        var sellContent = (
          <div>
            销售价格:<InputNumber onChange={this.onChange} defaultValue={0.01} min={0.01} step={0.01}/>以太币
            <Button type="primary" onClick={this.sell}>确定</Button>
          </div>
        );
        return (
          <div className="status-center">
            <div className='selldiv'>
              <Popover content={sellContent} placement='right' trigger='click'>
                <a className="button" >销售</a>
              </Popover>
            </div>
            <div className='selldiv'>
              <Popover content={<Fruit handleVisibleChange={this.handleVisibleChange.bind(this)} token={this.props.info.token_id} from={this.state.token}/>}  placement='right' trigger='click' visible={this.state.visible} onVisibleChange={this.handleVisibleChange} >
                <a className="button" >吃水果</a>
              </Popover>
            </div>
          </div>
        )
      }else if(status === 3){
        return (
          <div className="status-center">
            <a className="button" onClick={this.cancleSell}>取消销售</a>
          </div>
        )
      }else if(status === 4){
        return (
          <div className="status-center">
            <a className="button" >战斗中...</a>
          </div>
        )
      }else if(status === 5){
        return (
          <div className="status-center">
            <a className="button" >战斗确认中...</a>
          </div>
        )
      }else if(status === 6){
        return (
          <div className="status-center">
            <a className="button" >销售确认中...</a>
          </div>
        )
      }else if(status === 7){
        return (
          <div className="status-center">
            <a className="button" >销售确认，上链中...</a>
          </div>
        )
      }else if(status === 8){
        return (
          <div className="status-center">
            <a className="button" >吃水果中...</a>
          </div>
        )
      }
      else if(status === 9){
        return (
          <div className="status-center">
            <a className="button" >吃水果确认中...</a>
          </div>
        )
      }
    }else if(isSelf === 2){//market
      if(status === 6){
        return (
          <div className="status-center">
            <a className="button" >交易中</a>
          </div>
        )
      }else{
        return (
          <div className="status-center">
            <a className="button" onClick={this.buy}>购买</a>
          </div>
        )
      }
    }else if(isSelf === 3){//pk
      return (
        <div className="status-center">
          <a className="button" onClick={this.attack}>战斗</a>
        </div>
      )
    }else if(isSelf === 4){
      return (
        <div className="status-center">
          <a className="button" onClick={this.onPk}>战斗</a>
        </div>
      )
    }
  }

  onPk = () => {
    console.log(this.props);
    var token = this.props.token;
    console.log('token:'+token);
    console.log('changerId：'+this.props.info.token_id);
    console.log('regeistId:'+this.props.attackId);
    let data = {
      method : 'POST',
      headers: {
          'Content-Type': 'text/plain'
      },
      body : JSON.stringify({
          challengerId : this.props.info.token_id,
          resisterId : this.props.attackId,
      }), 
      credentials: 'same-origin',
    }
    fetch(BASE_URL+'/api/pk',data)
    .then(response => {
      return response.json()
    }).then(json => {
      if(json.status === 0){
        var sign = json.data.sign;
        console.log(this.props.info.token_id+'--'+this.props.attackId+'--'+sign+'--'+token);
        let Web3 = require('web3');
        let web3;
        if (typeof window.web3 !== 'undefined') {
          console.log('no web3');
          web3 = new Web3(window.web3.currentProvider)
        }
        var abiArray = ABI;
        var MyContract = web3.eth.contract(abiArray);
        var contractInstance = MyContract.at(CONTRACT_HASH);
        contractInstance.pk.sendTransaction(
          this.props.info.token_id,this.props.attackId,sign,
          {from:token},(err,result) => {
            console.log(err);
          if(!err){
            let data = {
              method : 'POST',
              headers: {
                'Content-Type': 'text/plain'
            },
            credentials: 'same-origin',
            body : JSON.stringify({
              hash : result,
              challengerId : this.props.info.token_id,
              resisterId : this.props.attackId
            }),
            }
            fetch(BASE_URL+'/api/pk/hash',data)
            .then(response => {
              return response.json()
            }).then(json => {
              if(json.status === 0){
                alert('战斗中，请稍后');
                this.props.refresh(true);
              }
            }) 
          }else{
            alert('交易失败,请重试.失败详情:'+result);
          }
        });    
      }else{
        alert(json.message);
      }
    });
  }

  onCancleM = () => {
    this.setState({
      modalVisible : false
    })
  }

  formatOwner(owener){
    if(!owener){
      return '';
    }
    if(owener.length>20){ 
        var substr = owener.substr(0,20);
        substr += "...";
    }
    return substr;
  }

  render() {
    console.log(this.props.info);
    let rarity = this.props.info.rarity;
    var owner = this.formatOwner(this.props.info.owner);
    if (this.state.toBuy) {
      return <Redirect
        to={{
          pathname: '/mybird',
          state: { token: this.state.token }
        }}
      />
    } else if (this.state.toInfo) {
      return <Redirect
        to={{
          pathname: '/info/' + this.props.info.token_id,
          state: { token: this.state.token }
        }}
      />
    }  else if(!this.props.info.token_id && this.props.info.token_id !== 0) {
      return (
        //<AttackModal visible={this.state.modalVisible} onCancleM={this.onCancleM} hasBird={this.state.hasBird} birdList={this.state.birdList} token={this.props.location.state.token} attackId={this.state.attackId} {...this.props}></AttackModal>
        <NoBird />
      );
    } else {
      return (
        <div className="bird-wrap">
          <div className="boxw">
            <div className="lbs">{this.props.info.weight} 盎司</div>
            <div className="name">{this.props.info.name || '无名鸟'}</div>
            <div className="clear"></div>
            <div className="rarity">{rarity || '-'}</div>
            <div className="lvl">等级 {this.props.info.level}</div>
            <div className="clear"></div>
            <div className="fish-img-outer rare spearfish level-1">
              <img className="clickable bird-logo" src={BASE_URL+this.props.info.svg_path} onClick={this.info} />
            </div>
            {this.props.info.auction_data ? 
              <div className="_tc">
                <div className="price bigrounded">{this.props.info.auction_data.price} ETH</div>
              </div>
              :<div className="_tc">
              <div className="price bigrounded">无</div>
            </div>
            }
            <div className="bird-num-wrap">
              <div className="bird-row">
                <div className="key">力量</div>
                <div className="value">{this.props.info.power}</div>
              </div>
              <div className="bird-row">
                <div className="key">速度</div>
                <div className="value">{this.props.info.speed}</div>
              </div>
              <div className="bird-row">
                <div className="key">经验</div>
                <div className="value">{this.props.info.exp}</div>
              </div>
            </div>
            {this.props.isSelf !== 1 ?  <div className="_tc">
              <div className="owned">所有者: <span className="owner clickable" onClick={this.info}>{owner}</span></div>
              </div>
              :<div></div>}
             {/* <a ><img className="info" src={info} onClick={this.info} /></a> */}
            <div >
              {this.statusRender()}
            </div>

          </div>
        </div>
      )
    }
  }
}