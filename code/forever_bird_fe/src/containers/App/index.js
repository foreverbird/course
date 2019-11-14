import React,{Component} from 'react';
import {
    Switch,
    Route
} from 'react-router-dom';
import mainFrame from '../main';
import mybird from '../../view/mybird';
import birdmarket from '../../view/birdmarket';
import catchBird from '../../view/catchbird';
import attack from '../../view/attack';
import sort from '../../view/sort';
import invite from '../../view/invite';
import info from '../../view/info';
import result from '../../view/result';
import profit from '../../view/profit.jsx';
import metamask from '../../view/metamask';
import register from '../../view/register';
import {Redirect} from 'react-router-dom';

class App extends Component{

    previousLocation = this.props.location

    token = null

    nextLocation = null;

    componentWillUpdate(nextProps){
        // console.log('netx')
        // console.log(nextProps)
        this.nextLocation = nextProps.location.pathname;
        if(!(this.nextLocation === '/metamask' || this.nextLocation === '/')){
            this.token = nextProps.location.state.token;
        }
        /* const {location} = this.props;
        if(nextProps.history.action !== 'POP' && (!location.state || !location.state.modal)) {
           this.previousLocation = this.props.location; 
        } */
    }


    render(){
        const { location } = this.props;
        const isModal = !!(location.state && location.state.modal && this.previousLocation !== location);
        if(!this.token){
            this.token = this.props.location.state ? this.props.location.state.token : null;
        }
        const isLogin = (this.token === null || typeof this.token === 'undefined') ? false : true;
        console.log(this.token);console.log(this.nextLocation);console.log(this.props);
        if(!isLogin && this.nextLocation !== '/metamask' && location.pathname !== '/'){
            return(
                <div>
                    <Redirect to = {{
                        pathname : '/metamask'
                    }}
                />
                </div>
            )
        }else{
            return(
                <div>
                    <Switch location={isModal ? this.previousLocation : location}>
                        {/* <Route exact path='/' component={mainFrame}/> */}
                        <Route exact path='/' component={mainFrame}/>
                        <Route exact path='/metamask' component={metamask}/>
                        <Route exact path='/register' component={register}/>
                        <Route exact path='/mybird' component={mybird}/>
                        <Route exact path='/birdmarket' component={birdmarket}/>
                        <Route exact path='/attack' component={attack}/>
                        <Route exact path='/catchbird' component={catchBird}/>
                        <Route exact path='/sort' component={sort}/>
                        <Route exact path='/invite' component={invite}/>
                        <Route exact path='/profit' component={profit}/>
                        <Route path='/info/:id' component={info} />
                        <Route path='/result/:hash' component={result} />
                    </Switch>
                </div>
            )
        }
    }
}

export default App