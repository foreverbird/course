import React,{Component} from 'react';
import MyHeader from '../components/myheader';
import RegisterFooter from '../components/registerFooter';
import Menu from '../containers/main/menu';
import Bird from '../containers/main/bird';
import {BASE_URL,TOKEN} from '../util/common';
import 'whatwg-fetch';
import '../css/mybird.less';
import {Popover,Button,Tooltip,List,Pagination} from 'antd';
import NoBird from './nobird';
import {Link} from 'react-router-dom';
import {CopyToClipboard} from 'react-copy-to-clipboard';

export default class MyBird extends Component{

    constructor(props){
        super(props);
        this.state = {
            token : this.props.location.state.token,
            nickName : this.props.location.state.nickName,
            // token : TOKEN,
            // nickName : 'ls',
            birdList : {
                current : 0,
                total : 0,
                birds : []
            },
            refresh : 0
        };
    }

    refresh(value){
        this.setState({
            refresh : value
        })
        this.getBirdInfo();
    }

    componentDidMount(){
        this.getUserInfo();
        this.getBirdInfo();
    }

    getBirdInfo() {
        this.getBirdList(1,12);
    }

    getBirdList(startPage,pageSize) {
        let data = {
            method : 'POST',
            headers: {
                'Content-Type': 'text/plain'
            },
            body : JSON.stringify({
                address : this.state.token,
                start_page : startPage,
                page_size : pageSize
            }), 
            credentials: 'same-origin',
        }
        fetch(BASE_URL+'/api/mybirds',data)
        .then(response => {
            return response.json()
        }).then(json => {
            console.log('mybirds')
            var data = json.data;
            if(data !== null){
                this.setState({
                    birdList : {
                        birds : json.data.catch_birds ? json.data.catch_birds.concat(json.data.birds) : json.data.birds,
                        total : json.data.total_number,
                        current : startPage,
                        // catchbirds : json.data.catch_birds
                    }
                });
            }else if(json.status === 1){
                const { history } = this.props;
                history.push('/metamask');
            }
        }) 
    }

    onPageChange = (page,pageSize) => {
        this.getBirdList(parseInt(page),parseInt(pageSize));
    }

    getUserInfo() {
        let data = {
            method : 'POST',
            headers: {
                'Content-Type': 'text/plain'
            },
            body : JSON.stringify({
                address : this.state.token
            }), 
            credentials: 'same-origin',
            // mode : 'cors'
        }
        fetch(BASE_URL+'/api/me',data)
        .then(response => {
            return response.json()
        }).then(json => {
            console.log('me')
            console.log(json);
            if(json.status === 0){
                this.setState({
                    nickName : json.data.nick
                });
            }else if(json.status === 1){
                const { history } = this.props;
                history.push('/metamask');
            }
        }) 
    }

    render(){
        const {birdList} = this.state;
        // const hasBird = birdList.total === 0 ? 0 : 1; 
        console.log('mybird user:'+this.state.token);
        if(birdList.total > 0){
            const pagination = {
                pageSize: 12,
                current: this.state.birdList.current,
                total: this.state.birdList.total,
                onChange: this.onPageChange,
            };
            return(
                <div className='main'>
                    <MyHeader />
                    <Menu {...this.props}/>
                    <div className="username-wrap">
                        <div className='info1'>
                            {/* <img src={logo} className=''/> */}
                            昵称：{this.state.nickName}<br/>
                            <Tooltip className="copy-btn" placement='topLeft' title={this.state.token} arrowPointAtCenter>
                                <CopyToClipboard text={this.state.token}>
                                    <a className="copy-btn">复制地址</a>
                                </CopyToClipboard>
                            </Tooltip>
                        </div>
                    </div>
                    {/* {this.state.birdList.catchbirds.length > 0 ? 
                        <div className="my-bird-list-wrap">
                            <List dataSource={this.state.birdList.catchbirds} itemLayout='vertical'
                                grid={{ gutter: 16, column: 4, xs: 1, md : 2,xl : 4, xxl : 4}}
                                renderItem={item => (<List.Item><NoBird /></List.Item>)}
                            >
                            </List>
                        </div>
                        :<div></div>
                    } */}
                    <div className="my-bird-list-wrap">
                        <List dataSource={this.state.birdList.birds} itemLayout='vertical'
                            grid={{ gutter: 16, column: 4, xs: 1, md : 2,xl : 4, xxl : 4}}
                            pagination={pagination}
                            renderItem={item => (<List.Item><Bird info={item} isSelf={1} refresh={this.refresh.bind(this)} {...this.props}/></List.Item>)}
                        >
                        </List>
                    </div>
                    <RegisterFooter />
                </div>
            );
        }else{
            return(
                <div className='main'>
                    <MyHeader />
                    <Menu {...this.props}/>
                    <div className="username-wrap">
                        <div className='info1'>
                            {/* <img src={logo} className=''/> */}
                            昵称：{this.state.nickName}<br/>
                            <Tooltip className="copy-btn" placement='topLeft' title={this.state.token} arrowPointAtCenter>
                                <CopyToClipboard text={this.state.token}>
                                    <a className="copy-btn">复制地址</a>
                                </CopyToClipboard>
                            </Tooltip>
                        </div>
                    </div>
                    <div className='bird'>
                        <div className='content'>
                            <div className="no-img-tip"></div>
                            <p className="no-img-tip-word">你当前没有任何加密鸟代币</p>
                        </div>
                        <div className="no-btn-wrap">
                            <div className="no-btn">
                                <Link to={{
                                    pathname: '/catchbird',
                                    state: {token: this.state.token}
                                }}>森林捉鸟</Link>
                            </div>
                            <div className="no-btn">
                                <Link to={{
                                    pathname: '/birdmarket',
                                    state: {token: this.state.token}
                                }}>鸟市购买</Link>
                            </div>
                        </div>
                    </div>
                    <RegisterFooter />
                </div>
            )
        }
    }
}