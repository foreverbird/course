import React, { Component } from 'react';
import {Redirect} from 'react-router-dom';
import '../../css/sort.less';
import {BASE_URL} from '../../util/common';
import small from '../../img/small.svg';

export default class SortBird extends Component {

    constructor(props){
        super(props);
        this.state = {
            redirect : false,
            id : 1
        }
    }

    onClick = (e) => {
        e.preventDefault();
        this.setState({
            redirect : true
        });
    }

    render() {
        if(this.state.redirect){
            return <Redirect 
                        to = {{
                            pathname : '/info/'+this.props.info.token_id,
                            state : {token :this.props.location.state.token}
                        }}
                    />
        }else{
            return (
                <div className='fish_item'>
                    <div className="container clickable" onClick={this.onClick}>
                        <div className="td1">
                            <div className="_tb5">{this.props.info.row_no}. {this.props.info.name}</div>
                            {/* <div className="rarity _mt5 _mb5">鲸鱼</div> */}
                            <div className="rarity">{this.props.info.rarity || '-'}</div>
                            <div className="c9">等级 {this.props.info.level}</div>
                            <div className="_mt15 _tb7 clickable bird-ower"><div className="c9 _tm1">所有者:</div>{this.props.info.owner}</div>
                        </div>
                        <div className="td2">
                            <div className="fish-img-outer whale cachalot level-16">
                                <img className="clickable" src={BASE_URL+this.props.info.svg_path} />
                            </div>
                        </div>
                        <div className="td3">
                            <div className="pars">
                                <table className="pars" style={{ width: 100 }}>
                                    <tbody ><tr ><td align="left">力量:</td><td align="right" className="b">{this.props.info.power}</td></tr>
                                        {/* <tr ><td align="left">敏捷:</td><td align="right" className="b">123</td></tr> */}
                                        <tr tooltipplacement="right"><td align="left">速度:</td><td align="right" className="b">{this.props.info.speed}</td></tr>
                                        <tr tooltipplacement="right"><td align="left">经验:</td><td align="right" className="b">{this.props.info.exp}</td></tr>
                                    </tbody></table>
                            </div>
                        </div>
                        <div className="td4">
                            <div className="_tb3 lb">{this.props.info.weight} 盎司</div>
                        </div>
                        <div className="clear"></div>
                    </div>
                </div>
            )
        }
    }
}