import React,{Component} from 'react';
import Menu from '../containers/main/menu';
import SortBird from '../containers/main/sortbird'
import {List,Pagination} from 'antd';
import MyHeader from '../components/myheader';
import RegisterFooter from '../components/registerFooter';
import {BASE_URL} from '../util/common';
import 'whatwg-fetch';

export default class Sort extends Component{

    constructor(props){
        super(props);
        this.state = {
            birdList : {
                count : 0,
                start : 0,
                total : 0,
                birds : []
            }
        }
    }

    componentDidMount(){
        this.getSortList(1,10);
    }

    
    getSortList(startPage,pageSize) {
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
        fetch(BASE_URL+'/api/rank',data)
        .then(response => {
            return response.json()
        }).then(json => {
            if(json.status === 0){
                this.setState({
                    birdList : {
                        birds : json.data.birds,
                        total : json.data.total_number,
                        current : startPage
                    }
                });
            }else if(json.status === 1){
                const { history } = this.props;
                history.push('/metamask');
            }
        }) 
    }

    onPageChange = (page,pageSize) => {
        this.getSortList(parseInt(page),parseInt(pageSize));
    }

    render(){
        const pagination = {
            pageSize: 10,
            current: this.state.birdList.current,
            total: this.state.birdList.total,
            onChange: this.onPageChange,
        };
        return(
             <div className="sort-main">
                <MyHeader />
                <Menu {...this.props}/>
                <div className="sort-wrap">
                        <List dataSource={this.state.birdList.birds}
                            pagination={pagination}
                            renderItem={item => (<List.Item><SortBird info={item} {...this.props}/></List.Item>)}
                        >
                        </List>
                    </div>
                <RegisterFooter />
             </div>
        )
    }
}