import React,{Component} from 'react';
import Menu from './menu';
import Detail from './detail';

export default class Body extends Component{

    constructor(props){
        super(props);
    }

    render(){
        return(
            <div className="bodyCss">
                <Menu {...this.props}/>
                <Detail />
            </div>
        );
    }
}