import React, { Component } from 'react';
import {Link} from 'react-router-dom';
import MyHeader from '../components/myheader';
import RegisterFooter from '../components/registerFooter';
import Menu from '../containers/main/menu';
import {Table, List, Pagination } from 'antd';
import {BASE_URL,TOKEN} from '../util/common';
import '../css/profit.less';
import moment from 'moment';
export default class Profit extends Component {

    constructor(props) {
        super(props);
        this.state = {
            token: this.props.location.state.token,
            activeTab: 'catch',
            statInfo: {},
            historyStatInfo: {},
            profitRank: [],
            historyList: [],
            pageInfo: {
                page_size: 10,
                start_page: 1,
                total_number: 0
            }
        };
        this.changeTab = this.changeTab.bind(this);
        this.getHistroy = this.getHistroy.bind(this);
    }

    componentDidMount() {
        // rank info
        this.getRankInfo();
        this.getHistroy();
        this.getHistoryStatInfo();
    }

    getRankInfo() {
        let data = {
            method : 'POST',
            headers: {
                'Content-Type': 'text/plain'
            },
            body : JSON.stringify({
                week : 0,
            }), 
            credentials: 'same-origin',
        }
        fetch(BASE_URL+'/api/record/profit',data)
        .then(response => {
            return response.json()
        }).then(json => {
            if(json.status === 0){
                this.setState({
                    profitRank: json.data.list,
                    statInfo: json.data.statistics
                });
            }else if(json.status === 1){
                const { history } = this.props;
                history.push('/metamask');
            }
        });
    }

    getHistoryStatInfo() {
        let data = {
            method : 'POST',
            headers: {
                'Content-Type': 'text/plain'
            },
            body : JSON.stringify({
                week : 0,
            }), 
            credentials: 'same-origin',
        }
        fetch(BASE_URL+'/api/record/tx/statistics',data)
        .then(response => {
            return response.json()
        }).then(json => {
            if(json.status === 0){
                this.setState({
                    historyStatInfo: json.data
                });
            }else if(json.status === 1){
                const { history } = this.props;
                history.push('/metamask');
            }
        });
    }

    getHistroy(page = 1, pageSize = 10) {
        let tab = this.state.activeTab;
        let url = '/api/record/catch/bird';
        if (tab === 'buy') {
            url = '/api/record/buy/bird';
        } else if (tab === 'eat') {
            url = '/api/record/buy/fruit';
        }
        let data = {
            method : 'POST',
            headers: {
                'Content-Type': 'text/plain'
            },
            body : JSON.stringify({
                week : 0,
                start_page: page,
                page_size : pageSize
            }),
            credentials: 'same-origin',
        }
        fetch(BASE_URL+url,data)
        .then(response => {
            return response.json()
        }).then(json => {
            if(json.status === 0){
                this.setState({
                    historyList: json.data.records,
                    total_number: json.data.total_number
                });
            }else if(json.status === 1){
                const { history } = this.props;
                history.push('/metamask');
            }
        });
    }

    changeTab(e) {
        if (e.target && e.target.dataset.tab) {
            this.setState({
                activeTab: e.target.dataset.tab,
                start_page: 1,
            }, () => {
                this.getHistroy();
            });
        }
    }

    render() {
        const pagination = {
            pageSize: 10,
            current: 1,
            total:1,
            onChange: 1,
        };
        let title = '';
        let content = '';
        if (this.state.activeTab !== 'buy') {
            title = <div className="total-rank-table-title">
                <p className="table-title-word">鸟名</p>
                <p className="table-title-word">钱包地址</p>
                <p className="table-title-word">交易金额</p>
                <p className="table-title-word">交易时间</p>
            </div>
            content = this.state.historyList && this.state.historyList.length && this.state.historyList.map((item, index) => {
                return (
                    <div className="total-rank-table-content" key={index}>
                        <p className="table-title-word">{item.BirdName}</p>
                        <p className="table-title-word">{item.address}</p>
                        <p className="table-title-word">{item.price} ETH</p>
                        <p className="table-title-word">{moment(item.day*1000).format('YYYY-MM-DD HH:mm:ss')}</p>
                    </div>
                )
            })
        } else {
            title = <div className="total-rank-table-title">
                <p className="table-title-word">鸟名</p>
                <p className="table-title-word">购买者地址</p>
                <p className="table-title-word">销售者地址</p>
                <p className="table-title-word">交易金额</p>
                <p className="table-title-word">交易时间</p>
            </div>
            content = this.state.historyList && this.state.historyList.length && this.state.historyList.map((item, index) => {
                return (
                    <div className="total-rank-table-content" key={index}>
                        <p className="table-title-word">{item.BirdName}</p>
                        <p className="table-title-word">{item.buyer}</p>
                        <p className="table-title-word">{item.seller}</p>
                        <p className="table-title-word">{item.price} ETH</p>
                        <p className="table-title-word">{moment(item.day*1000).format('YYYY-MM-DD HH:mm:ss')}</p>
                    </div>
                )
            })
        }
        return (
            <div className="main">
                <MyHeader />
                <Menu {...this.props}/>
                <div className="profit-wrap">
                    <div className="total-rank-div">
                        <div className="rank-table-wrap">
                            <div className="total-rank-title">上周排行榜<span>总金额{this.state.statInfo.total_profit}ETH</span></div>
                            <div className="total-rank-table">
                                <div className="total-rank-table-title">
                                    <p className="table-title-word rank-title">排名</p>
                                    <p className="table-title-word rank-title">鸟名</p>
                                    <p className="table-title-word rank-title">重量</p>
                                    <p className="table-title-word rank-title">拥有者</p>
                                    <p className="table-title-word rank-title">分红金额</p>
                                </div>
                                {
                                    this.state.profitRank && this.state.profitRank.length && this.state.profitRank.map((item, index) => {
                                        return (
                                            <div className={`total-rank-table-content ${item.rank < 4 ? 'top3' : item.rank < 11 ? 'top10' : ''}`} key={index}>
                                                <p className="table-title-word">{item.rank}</p>
                                                <p className="table-title-word">{item.bird_id}</p>
                                                <p className="table-title-word">{item.weight}盎司</p>
                                                <p className="table-title-word">{item.owner}</p>
                                                <p className="table-title-word">{item.profit} ETH</p>
                                            </div>
                                        )
                                    })
                                }
                            </div>
                        </div>
                        <div className="rank-table-wrap history-table-wrap">
                            <div className="total-rank-title">本周交易历史(<span>总交易量:{this.state.historyStatInfo.tx_count}次</span><span>总交易金额:{this.state.historyStatInfo.total_price}ETH</span>)</div>
                            <div className="tab-select-div" onClick={this.changeTab}>
                                <div className={`tab-select ${this.state.activeTab === 'catch' ? 'active' : ''}`} data-tab="catch">Catch</div>
                                <div className={`tab-select ${this.state.activeTab === 'buy' ? 'active' : ''}`} data-tab="buy">购买</div>
                                <div className={`tab-select ${this.state.activeTab === 'eat' ? 'active' : ''}`} data-tab="eat">吃水果</div>
                            </div>
                            <div className="total-rank-table">
                                {title}
                                {content}
                            </div>
                            <div className="page-div">
                                <Pagination defaultCurrent={this.state.start_page} total={this.state.total_number} onChange={this.getHistroy} />
                            </div>
                        </div>
                    </div>
                </div>

                <RegisterFooter />
            </div>
        )
    }
}