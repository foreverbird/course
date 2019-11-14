import React, {Component} from 'react';
import { Redirect,Link } from 'react-router-dom';
import { Menu } from 'antd';
import metamask from '../../util/web3';

export default class BirdMenu extends Component{

    constructor(props){
        super(props);
        this.state = {
            token :  this.props.location ? this.props.location.state.token : null,
            // token : '0xa16ade94150e1802fa855e846b1566a85da04e71',
            netId : 1,
            isRedirect : false
        };
    }

    componentDidMount(){
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
                /* if(netId !== this.state.netId){
                    this.setState({
                        token : this.state.token,
                        netId : netId,
                        isRedirect : true
                    });
                } */
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

    render(){
        var state =  {
            token : this.state.token,
            nickName : null//this.props.location.state.nickName
        };
        let key = '1';
        if(this.props.location){
            let pathName = this.props.location.pathname;
            switch (pathName) {
                case '/birdmarket':
                    key = '2';
                    break;
                case '/mybird':
                    key = '1';
                    break;
                case '/catchbird':
                    key = '3';
                    break;
                case '/attack':
                    key = '4';
                    break;
                case '/sort':
                    key = '5';
                    break;
                case '/invite':
                    key = '6';
                    break;
                case '/profit':
                    key = '7';
                    break;
                default:
                    key = '1';
                    break;
            }
        }
        console.log(state)
        const birdmarket = {
            pathname : '/birdmarket',
            state : state
        }
        const mybird = {
            pathname : '/mybird',
            state : state,
        }
        const catchBird = {
            pathname : '/catchbird',
            state : state,
        }
        const attatck = {
            pathname : '/attack',
            state : state,
        }
        const sort = {
            pathname : '/sort',
            state : state,
        }
        const invite = {
            pathname : '/invite',
            state : state
        }
        const profit = {
            pathname : '/profit',
            state : state
        }
        if(this.state.isRedirect){
            return (
                <Redirect to = {{
                        pathname : '/metamask',
                        state : {token :this.state.token}
                    }}
                />
            )
        }else{
            return(
                <div className='menu'>
                    <Menu
                        mode="horizontal"
                        defaultSelectedKeys={[key]}
                        className="menu-wrap"
                        subMenuCloseDelay={0}
                    >
                    <Menu.Item key="1" className="menu-item"><Link to={mybird}>我的小鸟</Link></Menu.Item>
                    <Menu.Item key="2" className="menu-item"><Link to={birdmarket}>鸟市</Link></Menu.Item>
                    <Menu.Item key="3" className="menu-item"> <Link to={catchBird}>捉鸟</Link></Menu.Item>
                    <Menu.Item key="4" className="menu-item"><Link to={attatck}>战斗</Link></Menu.Item>
                    <Menu.Item key="5" className="menu-item"><Link to={sort}>排行榜</Link></Menu.Item>
                    <Menu.Item key="7" className="menu-item"><Link to={profit}>分红榜</Link></Menu.Item>
                    <Menu.Item key="6" className="menu-item"><Link to={invite}>邀请好友</Link></Menu.Item>
                    </Menu>
                </div>
            );
        }
    }
}
