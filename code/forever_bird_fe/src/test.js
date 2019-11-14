import React from "react";


class Name extends React.Component{
    render(){
        return (
            <div>
                <h1>{this.props.name}</h1>
            </div>
        );
    }
}

class Link extends React.Component{
    render(){
        return (
            <a href={this.props.site}>
                {this.props.site}
            </a>
        );
    }
}

class WebSite extends React.Component{
    render(){
        return (
            <div>
                <Name name={this.props.name}/>
                <Link site={this.props.site}/>
            </div>
        );
    }
}

export default WebSite