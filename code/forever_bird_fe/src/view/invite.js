import React,{Component} from 'react';
import Menu from '../containers/main/menu';
import MyHeader from '../components/myheader';
import RegisterFooter from '../components/registerFooter';

export default class Invite extends Component{

    constructor(props){
        super(props);
        this.state = {
            token : this.props.location.state.token,
        };
    }


    render(){
        return(
              <div className="main">
                <MyHeader />
                <Menu {...this.props}/>
                <div className="invite-content">
                    <div className="invite-title">推荐链接</div>
                    <div className="invite-sub-title">传播你的链接来邀请朋友加入ForeverBird：</div>
                    <div className="my-super-link">https://www.baidu.com</div>
                    <div className="invite-sub-sub-title">帮我们传播ForeverBird在社交媒体分享你的链接并直接发给你的朋友：</div>
                    <div className="invite-btn-wrap">
                        <div className="invite-btn">FaceBook</div>
                        <div className="invite-btn">Twitter</div>
                        <div className="invite-btn">Reddit</div>
                        <div className="invite-btn">WhatsApp</div>
                        <div className="invite-btn">Email</div>
                    </div>
                    <div className="invite-btn-desc">有人关注你的链接并在30天内注册的话就会被认为是你的的ForeveBird里的朋友。</div>
                    <div className="invite-logo"></div>
                </div>
                <RegisterFooter />
              </div>
        );
    }
}