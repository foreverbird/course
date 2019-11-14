import * as React from 'react';
import './registerHeader.less';
import logo from '../img/register-header-logo.png';
export default class RegisterHeader extends React.Component {

    render() {
        return (
            <div className="register-header">
                <img className="register-header-logo" src={logo} />
                <div className="slogen-div">
                    <p className="slogen">欢迎进入Forever bird游戏</p>
                    <p className="sub-title">{this.props.subTitle}</p>
                </div>
            </div>
        );
    }
}