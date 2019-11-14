import React,{Component} from 'react';
import MyHeader from '../components/myheader';
import RegisterFooter from '../components/registerFooter';
import Menu from '../containers/main/menu';
import Bird from '../containers/main/bird';
import {List,Pagination,Select} from 'antd';
 
import '../css/birdmarket.less';
import {BASE_URL} from '../util/common';
import 'whatwg-fetch';

export default class BirdMarket extends Component{

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
            refresh : 0
        };
    }

    refresh(value){
        this.setState({
            refresh : value
        })
        this.getMarketBirdList(1,12,2,1);
    }

    getMarketBirdList(startPage,pageSize,sortType,orderType) {
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
        fetch(BASE_URL+'/api/market/birds',data)
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
        this.getMarketBirdList(1,12,2,1);
    }

    onPageChange = (page,pageSize) => {
        this.getMarketBirdList(parseInt(page),parseInt(pageSize),2,1);
    }

    onSelect1Change = (value) => {
        this.getMarketBirdList(1,12,parseInt(value),1);
    }

    onSelect2Change = (value) => {
        this.getMarketBirdList(1,12,2,parseInt(value));
    }

    render(){
        const tn = this.props.location.state;
        const {token} = tn;
        console.log('birdmarket user:'+token);
        const {birdList} = this.state;
        const {count,start,total,birds} = this.state.birdList;
        if(birds === 'undefined') {
            return(
                <div className="main">
                    <MyHeader />
                    <Menu {...this.props}/>
                    <div style={{paddingLeft:50}}>
                        <p>欢迎回来{token},开始捉鸟</p>
                    </div>
                    <RegisterFooter />
                </div>
            );
        } else{
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
                    <div className="bird-market">
                        <List dataSource={this.state.birdList.birds} itemLayout='vertical'
                            grid={{ gutter: 16, column: 4, xs: 1, md : 2,xl : 4, xxl : 4}}
                            pagination={pagination}
                            renderItem={item => (<List.Item><Bird info={item} refresh={this.refresh.bind(this)} isSelf={2} {...this.props}/></List.Item>)}
                        >
                        </List>
                    </div>
                    <RegisterFooter />
                </div>
            );
        }
    }
}