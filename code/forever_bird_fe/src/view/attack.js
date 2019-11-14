import React, { Component } from 'react';
import Menu from '../containers/main/menu';
import Bird from '../containers/main/bird';
import {List,Pagination,Select} from 'antd';
import MyHeader from '../components/myheader';
import RegisterFooter from '../components/registerFooter';
 
import '../css/attack.less';
import {BASE_URL} from '../util/common';
import 'whatwg-fetch';
import AttackModal from '../view/attackmodal';

export default class Attack extends Component {

    constructor(props){
        super(props);
        this.state = {
            birdList : {
                count : 0,
                start : 0,
                total : 0,
                birds : []
            },
            select1Value : '0',
            select2Value : '0',
            refresh : 0,
            modalVisible : false
        };
    }

    refresh = (value) =>{
        this.setState({
            refresh : value,
            modalVisible : false
        })
        this.getAttackBirdList(1,12,2,1);
    }

    getAttackBirdList(startPage,pageSize,sortType,orderType) {
        let data = {
            method : 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body : JSON.stringify({
                start_page : startPage,
                page_size : pageSize,
                sort_type : sortType,
                order_type : orderType,
            }), 
            credentials: 'same-origin',
        }
        console.log('date')
        console.log(data);
        fetch(BASE_URL+'/api/pk/list',data)
        .then(response => {
            return response.json()
        }).then(json => {
            if(json.status === 0){
                this.setState({
                    birdList : {
                        birds : json.data.birds,
                        total : json.data.total_number,
                        current : startPage
                    },
                    // select1Value : '0',
                    // select2Value : '0'
                });
            }else if(json.status === 1){
                const { history } = this.props;
                history.push('/metamask');
            }
        }) 
    }

    componentDidMount(){
        this.getAttackBirdList(1,12,2,1);
    }

    onPageChange = (page,pageSize) => {
        this.getAttackBirdList(parseInt(page),parseInt(pageSize),2,1);
    }

    onSelect1Change = (value) => {
        this.getAttackBirdList(1,12,parseInt(value),1);
    }

    onSelect2Change = (value) => {
        this.getAttackBirdList(1,12,2,parseInt(value));
    }

    onCancleM = () => {
        console.log('onCancleM')
        this.setState({
          modalVisible : false
        })
        // this.getAttackBirdList(1,12,2,1);
    }

    showModal = (hasBird,modalBirdList,attackId) => {
        console.log('showModal');
        console.log(this.state);
        const birdList = this.state.birdList;
        this.setState({
            modalVisible : true,
            hasBird : hasBird,
            modalBirdList : modalBirdList,
            attackId : attackId,
            birdList : birdList,
        });
    }

    render(){
        const tn = this.props.location.state;
        const {token} = tn;
        console.log('attack user:'+token);
        console.log(this.state);
        const {birdList} = this.state;
        const {count,start,total,birds} = this.state.birdList;
        const pagination = {
            pageSize: 12,
            current: this.state.birdList.current,
            total: this.state.birdList.total,
            onChange: this.onPageChange,
        };
        const Option = Select.Option;
        return(
            <div className='main'>
                <MyHeader />
                <Menu {...this.props}/>
                <AttackModal visible={this.state.modalVisible} refresh={this.refresh} onCancleM={this.onCancleM} hasBird={this.state.hasBird} birdList={this.state.modalBirdList} token={this.props.location.state.token} attackId={this.state.attackId} {...this.props}></AttackModal>
                <div className='search'>
                    <div className="search-content">
                        <div className="nickname">{/* 昵称：{this.state.nickname} */}</div>
                        <div className="token">{/* Token：{this.state.token} */}</div>
                        <div className="search-wrap">
                            <span>排序</span>
                            <Select className='select1' defaultValue={this.state.select1Value} onChange={this.onSelect1Change}>
                                <Option value='0'>稀有度</Option>
                                <Option value='1'>价格</Option>
                                <Option value='2'>时间</Option>
                            </Select>
                            <Select className='select2' defaultValue={this.state.select2Value} onChange={this.onSelect2Change}>
                                <Option value='0'>由高到低</Option>
                                <Option value='1'>由低到高</Option>
                            </Select>
                        </div>
                    </div>
                </div>
                <div className="attack-wrap">
                    <List dataSource={this.state.birdList.birds} itemLayout='vertical'
                        grid={{ gutter: 16, column: 4, xs: 1, md : 2,xl : 4, xxl : 4}}
                        pagination={pagination}
                        renderItem={item => (<List.Item><Bird info={item} showModal={this.showModal.bind(this)} refresh={this.refresh.bind(this)} isSelf={3} {...this.props}/></List.Item>)}
                    >
                    </List>
                </div>
                <RegisterFooter />
            </div>
        );
    }
}