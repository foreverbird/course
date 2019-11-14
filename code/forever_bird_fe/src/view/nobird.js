import React,{Component} from 'react';
import '../css/bird.less';
import catching from  '../img/catching.png'

export default class NoBird extends Component{

    render(){
        return(
            <div className="bird-wrap">
                <div className="boxw">
                    <div className="lbs">&nbsp;</div>
                        <div className="name">&nbsp;</div>
                        <div className="clear"></div>
                        <div className="rarity">&nbsp;</div>
                        <div className="lvl">&nbsp;</div>
                        <div className="clear"></div>
                    <div className="rare spearfish level-1 no-bird-logo-wrap">
                        <img className="clickable bird-logo" src={catching}  />
                    </div>
                    
                    <p className="no-bird-slogen">上链中</p>
                
                </div>
            </div>
        );
    }
}

