import React,{Component} from 'react';
import {
    Switch,
    Route
} from 'react-router-dom';
import mainFrame from '../main';

export default class Detail extends Component{
    render(){
        return(
            <div className="detail">
                {/* <Switch location={isModal ? this.previousLocation : location}>
                    <Route exact path='/' component={mainFrame}/>
                </Switch> */}
            </div>
        );
    }
}