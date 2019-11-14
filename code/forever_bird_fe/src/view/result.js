import React, { Component } from 'react';
import MyHeader from '../components/myheader';
import RegisterFooter from '../components/registerFooter';
import Menu from '../containers/main/menu';
import '../css/info.less';
import ResultBird from './resultbird';
import 'whatwg-fetch';
import {BASE_URL} from '../util/common';

export default class Result extends Component {

    constructor(props){
        super(props);
        this.state = {
            hash : this.props.match.params.hash,
            attacker : {},
            victim : {},
            isWin : false,
            number : 0
        }
    }

    componentDidMount(){
        console.log("result.js")
        console.log(this.props)
        const hash = this.props.match.params.hash;
        let data = {
            method : 'POST',
            headers: {
                'Content-Type': 'text/plain'
            },
            credentials: 'same-origin',
            body : JSON.stringify({
                hash : hash
            }),
        }
        fetch(BASE_URL+'/api/record/pk/info',data)
        .then(response => {
            return response.json()
        }).then(json => {
            console.log(json);
            if(json.status === 0){
                this.setState({
                    attacker : json.data.challenger_bird,
                    victim : json.data.resister_bird,
                    isWin : json.data.pk_record.is_win,
                    number : json.data.pk_record.winner_reward_coin
                });
            }else{
                alert(json.message);
            }
        }) 
    }

    render(){
        return(
            <div className="attack-res-main">
                <div className="attack-res-header">
                    <div className="attack-res-header-log"></div>
                </div>
                <div className='attack-res-wrap'>
                    <div className="attack-res-title"></div>
                    <div className='attack-res-container'>
                        <div className='attack-bird-left'>
                            <ResultBird params={this.state.attacker} isWin={this.state.isWin} number={this.state.number}/>
                        </div>
                        <div  className="attack-bird-buffer"></div>
                        <div className='attack-bird-right'>
                            <ResultBird params={this.state.victim} isWin={!this.state.isWin} number={this.state.number}/>
                        </div>
                    </div>
                </div>
                <RegisterFooter />
            </div>
        )
    }

}