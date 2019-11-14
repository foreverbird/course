import React, { Component } from 'react';
import smallinfo from '../img/smallinfo.svg';
import '../css/resultbird.less';
import {BASE_URL} from '../util/common';
import info from '../img/info.png';

export default class ResultBird extends Component {

    constructor(props){
        super(props);
    }

    render(){
        console.log('result')
        console.log(this.props)
        const params = this.props.params;
        const isWin = this.props.isWin;
        var number = '+'+this.props.number;
        var winnerStyle = 'shadow_2 boxw winner result0';
        var winStyle = 'win';
        var winFont = '赢家';
        if(!isWin){
            winnerStyle = 'shadow_2 boxw looser result0';
            winStyle = 'lost';
            number = '-'+this.props.number;
            winFont = '输家';
        }
        return(
            <div  className={winnerStyle}>
                <div  className="lbs">{params.weight} 盎司</div>
                <div  className="res-bird-name">{params.name || '无名鸟'}</div>
                <div  className="rarity">{params.rarity}</div>
                <div  className="lvl">等级 {params.level}</div>
                <div  className="result1">
                    <img  border="0" className='result2' width="auto" src={BASE_URL+params.svg_path}/>
                </div>

                <div className="bird-result-owned">所有者: {params.owner}</div>
                <div className="bird-num-wrap">

                <div className="bird-row">
                    <div className="key">力量</div>
                    <div className="value">{params.power}</div>
                </div>
                <div className="bird-row">
                    <div className="key">速度</div>
                    <div className="value">{params.speed}</div>
                </div>
                <div className="bird-row">
                    <div className="key">经验</div>
                    <div className="value">{params.exp}</div>
                </div>
                </div>


                <div className={winStyle}>{winFont}!
                    <div className="cf _tm1">{number} 盎司</div>
                </div>
            </div>
        )
    }
}